package main

import (
	"flag"
	"fmt"
	"go_server/src/lib/discovery"
	"go_server/src/recall_server/config"
	"go_server/src/recall_server/logic"
)

func main() {
	configFileName := flag.String("config", "./recall.toml", "config file")
	flag.Parse()
	c, err := config.NewConfig(*configFileName)
	if err != nil {
		fmt.Println("config file load failed:", *configFileName)
		return
	}
	discovery.Init(c.Etcd)
	if err = logic.Init(c); err == nil {
		if err = discovery.RecallRegister(c.Env, &logic.RecallServer{}); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("recall_server running...")
		discovery.WaitForClose()
	} else {
		fmt.Println(err)
	}
}
