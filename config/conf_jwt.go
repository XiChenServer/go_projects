package config

type Jwy struct {
	Secret  string `json:"secret" yaml:"secret"`
	Expires int    `json:"expires" yaml:"expires"`
	Issuer  string `yaml:"issuer" json:"issuer"`
}
