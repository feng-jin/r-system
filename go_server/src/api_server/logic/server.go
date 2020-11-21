package logic

import (
	"fmt"
	"go_server/src/api_server/config"
	"go_server/src/lib/discovery"
	"go_server/src/lib/proto/filter"
	"go_server/src/lib/proto/rank"
	"go_server/src/lib/proto/recall"
	"go_server/src/lib/util"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	RedisDB    *util.RedisClient
	RecallGrpc recall.RecallClient
	RankGrpc   rank.RankClient
	FilterGrpc filter.FilterClient
}

var ThisServer *ApiServer

func Init(c *config.Config) error {
	rdb, err := util.NewRedisClient(&c.RedisConfig)
	if err != nil {
		return err
	}
	recallGrpc := discovery.RecallResolve(c.Env)
	rankGrpc := discovery.RankResolve(c.Env)
	filterGrpc := discovery.FilterResolve(c.Env)
	ThisServer = &ApiServer{
		RedisDB:    rdb,
		RecallGrpc: recallGrpc,
		RankGrpc:   rankGrpc,
		FilterGrpc: filterGrpc,
	}
	return nil
}

func HttpServe(c *config.Config) {
	handler := getHttpHandler()
	fmt.Println("running ...")
	gracehttp.Serve(&http.Server{Addr: c.Host, Handler: handler})
}

func getHttpHandler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Any("/rec", Rec)
	return e
}
