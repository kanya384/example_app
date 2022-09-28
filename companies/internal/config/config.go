package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port int32
	Log  struct {
		Level string
	}
	PG struct {
		User    string
		Pass    string
		Port    string
		Host    string
		PoolMax int
		DbName  string
		Timeout int
	}
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
