package logic

import (
	"go_server/src/lib/discovery"
	"go_server/src/lib/proto/filter"
	"go_server/src/lib/util"
	"go_server/src/recall_server/config"
)

type RecallServer struct {
	RedisDB    *util.RedisClient
	FilterGrpc filter.FilterClient
}

var ThisServer *RecallServer

func Init(c *config.Config) error {
	rdb, err := util.NewRedisClient(&c.RedisConfig)
	if err != nil {
		return err
	}
	filterGrpc := discovery.FilterResolve(c.Env)
	ThisServer = &RecallServer{
		RedisDB:    rdb,
		FilterGrpc: filterGrpc,
	}
	return nil
}
