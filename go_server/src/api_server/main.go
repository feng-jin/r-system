package main

import (
	"flag"
	"fmt"
	"go_server/src/api_server/config"
	"log"
)

func main() {
	configFileName := flag.String("config", "./api.toml", "config file")
	flag.Parse()
	_, err := config.NewConfig(*configFileName)
	log.Print("aaa")
	if err != nil {
		log.Print(err)
		fmt.Println("config file load failed:", *configFileName)
		return
	}
}
