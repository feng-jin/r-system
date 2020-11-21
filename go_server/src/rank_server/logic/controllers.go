package logic

import (
	"context"
	"fmt"
	"go_server/src/lib/proto/rank"
)

func (s *RankServer) Rank(ctx context.Context, rankReq *rank.RankReq) (*rank.RankResp, error) {
	fmt.Println("rank")
	return new(rank.RankResp), nil
}
