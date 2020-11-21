package main

import (
	"flag"
	"fmt"
	"go_server/src/lib/discovery"
	proto "go_server/src/lib/proto/greet"
	"strings"
	"time"
)

var (
	EtcdAddr = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
	Env      = flag.String("Env", "test", "env")
)

func main() {
	flag.Parse()
	err := discovery.Init(strings.Split(*EtcdAddr, ";"))
	if err != nil {
		fmt.Println(err)
		return
	}
	c := discovery.GreetResolve(*Env)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Morning 调用...")
		ctx, cancel := discovery.Context1s()
		resp1, err := c.Morning(
			ctx,
			&proto.GreetRequest{Name: "Jinfeng"},
		)
		cancel()
		if err != nil {
			fmt.Println("Morning调用失败：", err)
			return
		}
		fmt.Printf("Morning 响应：%s，来自：%s\n", resp1.Message, resp1.From)

		fmt.Println("Night 调用...")
		ctx, cancel = discovery.Context1s()
		resp2, err := c.Night(
			ctx,
			&proto.GreetRequest{Name: "Jinfeng"},
		)
		cancel()
		if err != nil {
			fmt.Println("Night调用失败：", err)
			return
		}
		fmt.Printf("Night 响应：%s，来自：%s\n", resp2.Message, resp2.From)
	}
}
