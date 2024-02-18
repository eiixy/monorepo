package config

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/joho/godotenv"
)

func Load(path string, cfg any, envs ...string) {
	var fs []string
	for _, s := range envs {
		if s != "" {
			fs = append(fs, ".env."+s)
		}
	}
	LoadEnv(fs...)
	c := config.New(config.WithSource(file.NewSource(path), env.NewSource()))
	if err := c.Load(); err != nil {
		fmt.Printf("load config: %s\r\n", err.Error())
	}

	if err := c.Scan(cfg); err != nil {
		fmt.Printf("scan config: %s\r\n", err.Error())
	}
}

func LoadEnv(filenames ...string) {
	if len(filenames) == 0 {
		filenames = append(filenames, ".env")
	}
	if err := godotenv.Load(filenames...); err != nil {
		fmt.Printf("loading env file error: %s\r\n", err.Error())
	}
}
