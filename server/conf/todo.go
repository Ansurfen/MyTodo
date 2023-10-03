package conf

import (
	"bytes"

	"github.com/spf13/viper"
)

type TodoConf struct {
	file string
	*viper.Viper
}

type Option struct {
	Delimiter string `yaml:"delimiter" json:"delimiter"`
	File      string `yaml:"file" json:"file"`
	Bytes     []byte `yaml:"bytes" json:"bytes"`
	Type      string `yaml:"type" json:"type"`
}

func New(opt Option) *TodoConf {
	opts := []viper.Option{}
	if len(opt.Delimiter) > 0 {
		opts = append(opts, viper.KeyDelimiter(opt.Delimiter))
	}
	c := viper.NewWithOptions(opts...)
	if len(opt.Type) > 0 {
		c.SetConfigType(opt.Type)
	}
	if len(opt.File) > 0 {
		c.SetConfigFile(opt.File)
	}
	if len(opt.Bytes) > 0 {
		err := c.ReadConfig(bytes.NewReader(opt.Bytes))
		if err != nil {
			panic(err)
		}
	}
	return &TodoConf{Viper: c, file: opt.File}
}

func (c *TodoConf) SetFile(file string) *TodoConf {
	c.Viper.SetConfigFile(file)
	c.file = file
	return c
}

func (c *TodoConf) MustRead() *TodoConf {
	err := c.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return c
}

func (c *TodoConf) MustBind(v any) *TodoConf {
	err := c.Viper.Unmarshal(v)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *TodoConf) GetString(key string, defaultValue ...string) string {
	v := c.Viper.GetString(key)
	if len(defaultValue) > 0 && len(v) == 0 {
		return defaultValue[0]
	}
	return v
}
