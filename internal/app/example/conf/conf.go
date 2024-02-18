package conf

import (
	"github.com/eiixy/monorepo/internal/pkg/config"
)

type Config struct {
	Server struct {
		Http config.Server
	}
	Data struct {
		Database           config.Database
		Partnership        config.Database
		Kafka              config.Kafka
		KafkaConsumerGroup config.KafkaConsumerGroup
	}
	Trace struct {
		Endpoint string
	}
	Log config.Log
}

func Load(path string) *Config {
	var cfg Config
	config.Load(path, &cfg)
	return &cfg
}
