package conf

type EmailOption struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
	Auth struct {
		Host     string `yaml:"host"`
		ID       string `yaml:"id"`
		Username string `yaml:"username"`
		Password string `yaml:"passowrd"`
	} `yaml:"auth"`
}
