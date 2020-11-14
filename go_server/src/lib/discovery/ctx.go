package discovery

import (
	"context"
	"time"
)

// 1s超时
func Context1s() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.TODO(), time.Second)
}
