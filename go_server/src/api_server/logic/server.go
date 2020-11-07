package logic

import (
	"go_server/src/api_server/config"
	"go_server/src/lib/logger"
	"go_server/src/lib/util"
)

type Server struct {
	RedisDB *util.RedisClient
}

var ThisServer *Server

func Init(c *config.Config) error {
	logger.InitLogger(&c.LogConfig)
	rdb, err := util.NewRedisClient(&c.RedisConfig)
	if err != nil {
		return err
	}
	ThisServer = &Server{
		RedisDB: rdb,
	}
	return nil
}
