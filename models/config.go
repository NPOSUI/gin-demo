package models

var Config config

type config struct {
	SaName    string `mapstructure:"sa_name" yaml:"sa_name" json:"sa_name"`
	SaKey     string `mapstructure:"sa_key" yaml:"sa_key" json:"sa_key"`
	FileShare string `mapstructure:"file_share" yaml:"file_share" json:"file_share"`
}
