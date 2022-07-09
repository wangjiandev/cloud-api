package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 把配置映射为config对象

func LoadConfigFromToml(fileName string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(fileName, config)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}
