package conf

import (
	"os"

	"github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
)

var PrettyIdentedStyle = &gonginx.Style{
	SortDirectives:    false,
	SpaceBeforeBlocks: true,
	StartIndent:       0,
	Indent:            4,
	Debug:             false,
}

type NginxParser struct {
	*parser.Parser
	*gonginx.Config
	file string
}

type NginxConfOption struct {
	File string `json:"file"`
	Data string `json:"data"`
}

func NewNginxParser(opt NginxConfOption) (*NginxParser, error) {
	var (
		p   *parser.Parser
		err error
	)
	if len(opt.File) > 0 {
		p, err = parser.NewParser(opt.File)
	}
	if len(opt.Data) > 0 {
		p = parser.NewStringParser(opt.Data)
	}
	if err != nil {
		return nil, err
	}
	return &NginxParser{
		file:   opt.File,
		Parser: p,
		Config: p.Parse()}, nil
}

func (p *NginxParser) Write() error {
	return os.WriteFile(p.file, []byte(gonginx.DumpBlock(p.Config.Block, PrettyIdentedStyle)), 0666)
}

func (p *NginxParser) WriteAs(file string) error {
	return os.WriteFile(file, []byte(gonginx.DumpBlock(p.Config.Block, PrettyIdentedStyle)), 0666)
}
