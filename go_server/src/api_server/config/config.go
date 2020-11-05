package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Host string   `toml:"server_host"`
	Etcd []string `toml:"etcd"`
	Env  string   `toml:"env"`
}

func NewConfig(fileName string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fileName, &c)
	return &c, err
}
