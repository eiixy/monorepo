package conf

import (
	"github.com/eiixy/monorepo/internal/pkg/config"
)

type Config struct {
	Key    string
	Name   string
	Server struct {
		Http config.Server
	}
	Data struct {
		Database config.Database
	}
	Email config.Email
	Trace struct {
		Endpoint string
	}
	Log config.Log
}
