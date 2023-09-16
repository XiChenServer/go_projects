package config

type Email struct {
	Host             string `yaml:"host" json:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `yaml:"user" json:"user"`
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"`
	UerSSL           bool   `json:"uer_ssl" yaml:"uer_ssl"`
	UserTls          bool   `yaml:"user_tls" json:"user_tls"`
}
