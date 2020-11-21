package discovery

import (
	"fmt"
	"go_server/src/lib/discovery/config"
	"go_server/src/lib/proto/filter"
	"go_server/src/lib/proto/greet"
	"go_server/src/lib/proto/rank"
	"go_server/src/lib/proto/recall"

	"google.golang.org/grpc"
)

func GreetRegister(env string, server greet.GreetServer) error {
	listener, host, err := getListener()
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return err
	}
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

func RecallRegister(env string, server recall.RecallServer) error {
	listener, host, err := getListener()
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return err
	}
	srv := grpc.NewServer()
	go srv.Serve(listener)
	recall.RegisterRecallServer(srv, server)
	service = &Service{Name: config.RecallServer, Host: host, Env: env}
	err = service.register()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RankRegister(env string, server rank.RankServer) error {
	listener, host, err := getListener()
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return err
	}
	srv := grpc.NewServer()
	go srv.Serve(listener)
	rank.RegisterRankServer(srv, server)
	service = &Service{Name: config.RankServer, Host: host, Env: env}
	err = service.register()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func FilterRegister(env string, server filter.FilterServer) error {
	listener, host, err := getListener()
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return err
	}
	srv := grpc.NewServer()
	go srv.Serve(listener)
	filter.RegisterFilterServer(srv, server)
	service = &Service{Name: config.FilterServer, Host: host, Env: env}
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

func RecallResolve(env string) recall.RecallClient {
	return recall.NewRecallClient(Resolver(env, config.RecallServer))
}

func RankResolve(env string) rank.RankClient {
	return rank.NewRankClient(Resolver(env, config.RankServer))
}

func FilterResolve(env string) filter.FilterClient {
	return filter.NewFilterClient(Resolver(env, config.FilterServer))
}
