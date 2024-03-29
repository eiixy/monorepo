package config

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Load[T any](path string, cfg T) (*T, error) {
	_ = godotenv.Load(".env")
	readFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(replaceEnvVariables(readFile)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func replaceEnvVariables(text []byte) []byte {
	// 替换文本中的环境变量名
	replacedText := string(text)
	for _, env := range os.Environ() {
		envPair := strings.SplitN(env, "=", 2)
		envName := envPair[0]
		envValue := envPair[1]
		replacedText = strings.ReplaceAll(replacedText, fmt.Sprintf("${%s}", envName), envValue)
	}
	return []byte(replacedText)
}
