package starter

import (
	"MyTodo/conf"
	"MyTodo/engine/v2/cli"
	"MyTodo/utils"
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"
	"unicode"

	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkmidjwt "github.com/rookie-ninja/rk-entry/v2/middleware/jwt"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
)

type ServiceUnit struct {
	index   int    `json:"-"`
	Port    uint64 `json:"port"`
	Enabled bool   `json:"enabled"`
}

type GrpcBootConfig rkgrpc.BootConfig

// fix exceptional exit when meet null value
func (c *GrpcBootConfig) fixedNullPointer() {
	v := 1000000
	for i := 0; i < len(c.Grpc); i++ {
		c.Grpc[i].Middleware.Jwt.Symmetric = &rkmidjwt.SymmetricConfig{}
		c.Grpc[i].Middleware.Jwt.Asymmetric = &rkmidjwt.AsymmetricConfig{}
		c.Grpc[i].Middleware.RateLimit.ReqPerSec = &v
	}
}

func (c *GrpcBootConfig) Service() map[string]ServiceUnit {
	infos := make(map[string]ServiceUnit)
	for i := 0; i < len(c.Grpc); i++ {
		info := c.Grpc[i]
		infos[info.Name] = ServiceUnit{
			index:   i,
			Port:    info.Port,
			Enabled: info.Enabled,
		}
	}
	return infos
}

type Option struct {
	Config []byte
}

type GrpcServer struct {
	rkb        *rkboot.Boot
	GrpcOption GrpcBootConfig
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// equal to "gwOption: null"
var gwOption = []byte{103, 119, 79, 112, 116, 105, 111, 110, 58, 32, 110, 117, 108, 108}

func New(opt Option) *GrpcServer {
	tmp, err := os.CreateTemp("", "*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmp.Name())

	confxx := conf.New(conf.Option{
		Type:  "yaml",
		File:  tmp.Name(),
		Bytes: opt.Config,
	})
	cfg := GrpcBootConfig{}
	if err = confxx.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	cfg.fixedNullPointer()

	services := cfg.Service()

	cli.Execute()

	for _, f := range cli.Option.Filter {
		if isInt(f) {
			for _, unit := range services {
				port, _ := strconv.Atoi(f)
				if unit.Port == uint64(port) && unit.Enabled {
					cfg.Grpc[unit.index].Enabled = false
				}
			}
		} else {
			unit := services[f]
			if unit.Enabled {
				cfg.Grpc[unit.index].Enabled = false
			}
		}
	}

	prefix := cli.Option.Tmpl
	tmplIndex := -1

	if len(prefix) > 0 {
		if unit, ok := services[prefix]; ok {
			tmplIndex = unit.index
		} else {
			panic("invalid template")
		}
	} else {
		prefix = "unknown"
	}

	for i := 0; i < cli.Option.Range; i++ {
		port, err := utils.RandomPort()
		if err != nil {
			panic(err)
		}
		serviceName := fmt.Sprintf("%s_%d", prefix, port)
		if tmplIndex != -1 {
			info := cfg.Grpc[tmplIndex]
			info.Name = serviceName
			info.Port = uint64(port)
			info.Enabled = true
			cfg.Grpc = append(cfg.Grpc, info)
		} else {
			panic("empty template")
		}
	}

	confxx.Set("grpc", cfg.Grpc)

	if err = confxx.WriteConfig(); err != nil {
		panic(err)
	}

	opt.Config, err = os.ReadFile(tmp.Name())
	if err != nil {
		panic(err)
	}

	// fix exceptional exit when meet null value
	opt.Config = bytes.ReplaceAll(opt.Config, gwOption, []byte{})

	s := &GrpcServer{
		rkb:        rkboot.NewBoot(rkboot.WithBootConfigRaw(opt.Config)),
		GrpcOption: cfg,
	}

	return s
}

func (s *GrpcServer) Bootstrap() {
	s.rkb.Bootstrap(context.Background())
}

type LoadOption struct {
	Name         string
	GrpcHandlers []rkgrpc.GrpcRegFunc
	GWHandlers   []rkgrpc.GwRegFunc
}

type GrpcLoadFuncs []rkgrpc.GrpcRegFunc
type GWLoadFuncs []rkgrpc.GwRegFunc

func (s *GrpcServer) Load(opt LoadOption) {
	thread := rkgrpc.GetGrpcEntry(opt.Name)
	thread.AddRegFuncGrpc(opt.GrpcHandlers...)
	thread.AddRegFuncGw(opt.GWHandlers...)
}

func (s *GrpcServer) WaitForShutdownSig(shutdown ...func()) {
	s.rkb.WaitForShutdownSig(context.Background())
	if len(shutdown) > 0 {
		shutdown[0]()
	}
}

