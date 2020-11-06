package main

import (
	"flag"
	"fmt"
	"go_server/src/api_server/config"
	"go_server/src/api_server/logic"
	"log"
)

func main() {
	configFileName := flag.String("config", "./api.toml", "config file")
	flag.Parse()
	c, err := config.NewConfig(*configFileName)

	if err != nil {
		log.Print(err)
		fmt.Println("config file load failed:", *configFileName)
		return
	}
	logic.InitServer(c)
}
