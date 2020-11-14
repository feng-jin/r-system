package discovery

import (
	"fmt"
	"go_server/src/lib/discovery/config"
	"go_server/src/lib/proto/greet"
	"google.golang.org/grpc"
)

func GreetRegister(env string, server greet.GreetServer) error {
	listener, host, err := getListener()
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return err
	}
	fmt.Println("host:", host)
	srv := grpc.NewServer()
	go srv.Serve(listener)
	greet.RegisterGreetServer(srv, server)
	service = &Service{Name: config.GreetServer, Host: host, Env: env}
	err = service.register()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GreetResolve(env string) greet.GreetClient {
	return greet.NewGreetClient(Resolver(env, config.GreetServer))
}
