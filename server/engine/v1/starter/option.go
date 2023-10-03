package starter

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
)

type Middleware []gin.HandlerFunc

type RouteRegistry map[string]func(g *MyTodoServerGroup)

type Option struct {
	Profile    string            `yaml:"profile" json:"profile"`
	Mode       string            `yaml:"mode" json:"mode"`
	SW         OptionSwagger     `yaml:"sw" json:"sw"`
	Middleware []gin.HandlerFunc `yaml:"-" json:"-"`
	Registry   RouteRegistry     `yaml:"-" json:"-"`
	Filter     []string          `yaml:"filter" json:"filter"`
}

type OptionSwagger struct {
	Enabled          bool       `yaml:"enabled" json:"enabled"`
	Spec             *swag.Spec `yaml:"-" json:"-"`
	Version          string     `yaml:"version" json:"version"`
	BasePath         string     `yaml:"basePath" json:"basePath"`
	Host             string     `yaml:"host" json:"host"`
	Schemes          []string   `yaml:"schemes" json:"schemes"`
	Title            string     `yaml:"title" json:"title"`
	Description      string     `yaml:"desc" json:"desc"`
	InfoInstanceName string     `yaml:"infoName" json:"infoName"`
	SwaggerTemplate  string     `yaml:"template" json:"template"`
	LeftDelim        string     `yaml:"leftDelim" json:"leftDelim"`
	RightDelim       string     `yaml:"rightDelim" json:"rightDelim"`
}
