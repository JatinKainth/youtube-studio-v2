package config

import (
	"youtube-studio-v2/dtype"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

var appConfig dtype.AppConfig

func InitializeAppConfig() (*dtype.AppConfig, error) {
	config := koanf.New(".")
	parser := toml.Parser()
	err := config.Load(file.Provider("credentials.toml"), parser)
	if err != nil {
		return nil, err
	}

	if err := config.UnmarshalWithConf("", &appConfig, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
		return nil, err
	}

	return &appConfig, nil
}

func GetConfigValues() *dtype.AppConfig {
	return &appConfig
}
