package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_server/src/lib/discovery"
	"go_server/src/lib/proto/api"
	"go_server/src/lib/proto/rank"
	"go_server/src/lib/proto/recall"
)

func Rec(context *gin.Context) {
	recReq := new(api.RecReq)
	err := context.BindJSON(recReq)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	bucket := ""

	recallReq := recall.RecallReq{
		Req:    recReq,
		Bucket: bucket,
	}
	ctx, cancel := discovery.Context1s()
	recallResp, err := ThisServer.RecallGrpc.Recall(ctx, &recallReq)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}

	rankReq := rank.RankReq{
		Req:        recReq,
		RecallResp: recallResp,
		Bucket:     bucket,
	}
	ctx, cancel = discovery.Context1s()
	rankResp, err := ThisServer.RankGrpc.Rank(ctx, &rankReq)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}

	items := []*api.RecRespItem{
		{ItemId: "item_id1", Params: "test1"},
		{ItemId: "item_id2", Params: "test2"},
	}
	var resp = api.RecResp{
		Items: items,
	}

	context.JSON(200, resp)

	fmt.Println(*rankResp)
}
