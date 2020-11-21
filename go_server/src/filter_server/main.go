package main

import (
	"flag"
	"fmt"
	"go_server/src/filter_server/config"
	"go_server/src/filter_server/logic"
	"go_server/src/lib/discovery"
)

func main() {
	configFileName := flag.String("config", "./filter.toml", "config file")
	flag.Parse()
	c, err := config.NewConfig(*configFileName)
	if err != nil {
		fmt.Println("config file load failed:", *configFileName)
		return
	}
	discovery.Init(c.Etcd)
	if err = logic.Init(c); err == nil {
		if err = discovery.FilterRegister(c.Env, &logic.FilterServer{}); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("filter_server running...")
		discovery.WaitForClose()
	} else {
		fmt.Println(err)
	}
}
