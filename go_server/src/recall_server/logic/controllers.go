package logic

import (
	"context"
	"fmt"

	"go_server/src/lib/discovery"
	"go_server/src/lib/proto/filter"
	"go_server/src/lib/proto/recall"
)

func (s *RecallServer) Recall(ctx context.Context, recallReq *recall.RecallReq) (*recall.RecallResp, error) {
	fmt.Println("Recall")
	filterReq := filter.HistoryFilterReq{
		UserId:     recallReq.Req.UserId,
		RecallResp: new(recall.RecallResp),
	}
	ctx, cancel := discovery.Context1s()
	filterResp, err := ThisServer.FilterGrpc.HistoryFilter(ctx, &filterReq)
	cancel()
	if err != nil {
		fmt.Println(err)
	}
	return filterResp.RecallResp, nil
}
