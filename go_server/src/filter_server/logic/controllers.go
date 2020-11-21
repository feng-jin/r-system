package logic

import (
	"context"
	"fmt"
	"go_server/src/lib/proto/filter"
	"go_server/src/lib/proto/recall"
)

func (s *FilterServer) HistoryFilter(ctx context.Context, filterReq *filter.HistoryFilterReq) (*filter.HistoryFilterResp, error) {
	fmt.Println("HistoryFilter")
	resp := new(filter.HistoryFilterResp)
	resp.RecallResp = new(recall.RecallResp)
	return resp, nil
}

func (s *FilterServer) HistoryAdd(ctx context.Context, addReq *filter.HistoryAddReq) (*filter.HistoryAddResp, error) {
	fmt.Println("HistoryAdd")
	return new(filter.HistoryAddResp), nil
}
