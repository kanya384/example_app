package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Kafka struct {
		Brokers []string
		Topic   string
		Group   string
	}

	Log struct {
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

	Graylog struct {
		Host string
	}

	Email struct {
		Host  string
		Port  string
		Login string
		Pass  string
	}
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
