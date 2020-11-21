package main

import (
	"flag"
	"fmt"
	"go_server/src/api_server/config"
	"go_server/src/api_server/logic"
	"go_server/src/lib/discovery"
)

func main() {
	configFileName := flag.String("config", "./api.toml", "config file")
	flag.Parse()
	c, err := config.NewConfig(*configFileName)
	if err != nil {
		fmt.Println("config file load failed:", *configFileName)
		return
	}
	discovery.Init(c.Etcd)
	if err = logic.Init(c); err == nil {
		logic.HttpServe(c)
	} else {
		fmt.Println(err)
	}
}
