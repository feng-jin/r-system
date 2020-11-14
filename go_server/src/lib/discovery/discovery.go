package discovery

import (
	"fmt"
	"go_server/src/lib/discovery/config"
	"go_server/src/lib/logger"
	"strings"

	"go.etcd.io/etcd/clientv3"
)

var (
	client *clientv3.Client
)

// Init 初始化etcd
func Init(etcdAddr string) error {
	var err error
	if client == nil {
		//构建etcd client
		client, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(etcdAddr, ";"),
			DialTimeout: config.Timeout,
		})
		if err != nil {
			logger.Error("连接etcd失败：%s\n", err)
			fmt.Printf("连接etcd失败：%s\n", err)
			return err
		}
	}
	return nil
}
