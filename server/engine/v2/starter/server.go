package starter

import (
	"MyTodo/conf"
	"MyTodo/engine/v2/cli"
	"MyTodo/utils"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"

	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkmidjwt "github.com/rookie-ninja/rk-entry/v2/middleware/jwt"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	rkgrpcjwt "github.com/rookie-ninja/rk-grpc/v2/middleware/jwt"
)

type ServiceUnit struct {
	index   int    `json:"-"`
	Port    uint64 `json:"port"`
	Enabled bool   `json:"enabled"`
}

type GrpcBootConfig rkgrpc.BootConfig

type IssuedConfig struct {
	Jwt rkmidjwt.BootConfig `yaml:"jwt"`
}

// fix exceptional exit when meet null value
func (c *GrpcBootConfig) fixedNullPointer() {
	v := 1000000
	for i := 0; i < len(c.Grpc); i++ {
		c.Grpc[i].Middleware.Jwt.Symmetric = &rkmidjwt.SymmetricConfig{}
		c.Grpc[i].Middleware.Jwt.Asymmetric = &rkmidjwt.AsymmetricConfig{}
		c.Grpc[i].Middleware.RateLimit.ReqPerSec = &v
	}
}

// returns avaliable services
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

var (
	ErrInvalidTemplate = errors.New("invalid template")
	ErrEmptyTemplate   = errors.New("empty template")
)

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
			panic(ErrInvalidTemplate)
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
			panic(ErrEmptyTemplate)
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

func (s *GrpcServer) ApplyIssuedMiddleware() {
	var conf IssuedConfig
	if err := s.DefaultUserConf().Unmarshal(&conf); err != nil {
		panic(err)
	}
	if conf.Jwt.Enabled {
		for name := range s.GrpcOption.Service() {
			rkgrpc.GetGrpcEntry(name).AddUnaryInterceptors(rkgrpcjwt.UnaryServerInterceptor(rkmidjwt.ToOptions(
				&rkmidjwt.BootConfig{
					Enabled: true,
					Symmetric: &rkmidjwt.SymmetricConfig{
						Algorithm: conf.Jwt.Symmetric.Algorithm,
						Token:     conf.Jwt.Symmetric.Token,
					},
					Ignore:      conf.Jwt.Ignore,
					TokenLookup: conf.Jwt.TokenLookup,
					AuthScheme:  conf.Jwt.AuthScheme,
				}, name, rkgrpc.GrpcEntryType,
			)...))
		}
	}
}

func (s *GrpcServer) DefaultUserConf() *rkentry.ConfigEntry {
	return rkentry.GlobalAppCtx.GetConfigEntry("default")
}

func (s *GrpcServer) Bootstrap() {
	s.rkb.Bootstrap(context.Background())
}

type LoadOption struct {
	Name         string
	GrpcHandlers []rkgrpc.GrpcRegFunc
	GWHandlers   []rkgrpc.GwRegFunc
}

type (
	GrpcLoadFuncs []rkgrpc.GrpcRegFunc
	GWLoadFuncs   []rkgrpc.GwRegFunc
)

func (s *GrpcServer) NewThread(opt LoadOption) *rkgrpc.GrpcEntry {
	thread := rkgrpc.GetGrpcEntry(opt.Name)
	thread.AddRegFuncGrpc(opt.GrpcHandlers...)
	thread.AddRegFuncGw(opt.GWHandlers...)
	return thread
}

func (s *GrpcServer) WaitForShutdownSig(shutdown ...func()) {
	s.rkb.WaitForShutdownSig(context.Background())
	if len(shutdown) > 0 {
		shutdown[0]()
	}
}
