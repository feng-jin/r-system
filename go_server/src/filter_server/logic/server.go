package logic

import (
	"go_server/src/filter_server/config"
	"go_server/src/lib/proto/filter"
	"go_server/src/lib/proto/rank"
	"go_server/src/lib/proto/recall"
	"go_server/src/lib/util"
)

type FilterServer struct {
	RedisDB    *util.RedisClient
	RecallGrpc recall.RecallClient
	RankGrpc   rank.RankClient
	FilterGrpc filter.FilterClient
}

var ThisServer *FilterServer

func Init(c *config.Config) error {
	rdb, err := util.NewRedisClient(&c.RedisConfig)
	if err != nil {
		return err
	}
	ThisServer = &FilterServer{
		RedisDB: rdb,
	}
	return nil
}
