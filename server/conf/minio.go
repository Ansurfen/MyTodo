package conf

type MinioOption struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	ID     string `yaml:"id"`
	Secret string `yaml:"secret"`
	Secure bool   `yaml:"secure"`
}
