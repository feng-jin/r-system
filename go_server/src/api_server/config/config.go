package config

import (
	"go_server/src/lib/logger"
	"go_server/src/lib/util"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host        string           `toml:"host"`
	Env         string           `toml:"env"`
	RedisConfig util.RedisConfig `toml:"Redis"`
	LogConfig   logger.LogConfig `toml:"Log"`
	Etcd        []string         `toml:"etcd"`
}

func NewConfig(fileName string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fileName, &c)
	return &c, err
}
