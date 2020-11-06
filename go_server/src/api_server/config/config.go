package config

import (
	"fmt"
	"go_server/src/lib/util"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host        string           `toml:"server_host"`
	RedisConfig util.RedisConfig `toml:"Redis"`
	Etcd        []string         `toml:"etcd"`
	Env         string           `toml:"env"`
}

func NewConfig(fileName string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fileName, &c)
	fmt.Println(c.Host)
	return &c, err
}
