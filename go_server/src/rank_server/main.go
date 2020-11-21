package main

import (
	"flag"
	"fmt"
	"go_server/src/lib/discovery"
	"go_server/src/rank_server/config"
	"go_server/src/rank_server/logic"
)

func main() {
	configFileName := flag.String("config", "./rank.toml", "config file")
	flag.Parse()
	c, err := config.NewConfig(*configFileName)
	if err != nil {
		fmt.Println("config file load failed:", *configFileName)
		return
	}
	discovery.Init(c.Etcd)
	if err = logic.Init(c); err == nil {
		if err = discovery.RankRegister(c.Env, &logic.RankServer{}); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("rank_server running...")
		discovery.WaitForClose()
	} else {
		fmt.Println(err)
	}
}
