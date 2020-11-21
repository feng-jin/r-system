package logic

import (
	"go_server/src/lib/util"
	"go_server/src/rank_server/config"
)

type RankServer struct {
	RedisDB *util.RedisClient
}

var ThisServer *RankServer

func Init(c *config.Config) error {
	rdb, err := util.NewRedisClient(&c.RedisConfig)
	if err != nil {
		return err
	}
	ThisServer = &RankServer{
		RedisDB: rdb,
	}
	return nil
}
