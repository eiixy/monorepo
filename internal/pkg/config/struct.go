package config

import (
	"strings"
)

type Registry struct {
	Type       string
	Addrs      string
	NacosGroup string
}

type Log struct {
	Dir          string
	Level        string
	MaxAge       int
	RotationTime int
}

type Database struct {
	Driver          string
	Dsn             string
	MaxOpenConns    string
	MaxIdleConns    string
	ConnMaxLifetime string // 秒
	ConnMaxIdleTime string // 秒
}

type Email struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

type Redis struct {
	Addr     string
	DB       string
	Password string
	Prefix   string
}

type ElasticSearch struct {
	Hosts    string
	UserName string
	Password string
	Sniff    string
}

type Kafka struct {
	Addrs    string `yaml:"addrs"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type KafkaProducer struct {
	Topic string `yaml:"topic"`
}

type KafkaConsumerGroup struct {
	GroupId string   `yaml:"groupId"`
	Topics  []string `yaml:"topics"`
}

type ClickHouse struct {
	Addr     string `json:"addr,omitempty"`
	Database string `json:"database,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (c ClickHouse) GetAddr() []string {
	return strings.Split(c.Addr, ",")
}

type Aliexpress struct {
	AppKey     string
	AppSecret  string
	GatewayUrl string
}

func (k Kafka) GetAddrs() []string {
	return strings.Split(k.Addrs, ",")
}
