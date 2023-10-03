package conf

type ETCDOption struct {
	DialTimeout int      `yaml:"dialTimeout"`
	EndPoints   []string `yaml:"endpoints"`
}
