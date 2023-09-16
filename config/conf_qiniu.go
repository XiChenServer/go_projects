package config

type QiNiu struct {
	Enable    bool   `json:"enable" yaml:"enable"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	Bucket    string `json:"bucket" yaml:"bucket"`
	CDN       string `yaml:"cdn" json:"cdn"`
	Zone      string `json:"zone" yaml:"zone"`
	Size      int    `yaml:"size" json:"size"`
}
