package config

type Mysql struct {
	Host     string `yaml:"host"`
	Post     int    `yaml:"post"`
	DB       string `yaml:"db"`
	user     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLever string `yaml:"log_lever"`
}
