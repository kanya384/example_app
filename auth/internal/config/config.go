package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port int

	Kafka struct {
		Brokers []string
		Topic   string
		GroupID string
	}

	Log struct {
		Level string
	}

	Graylog struct {
		Host string
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

	Token struct {
		Salt string
	}
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
