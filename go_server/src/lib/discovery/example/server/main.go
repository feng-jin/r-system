package main

import (
	"context"
	"flag"
	"fmt"
	"go_server/src/lib/discovery"
	proto "go_server/src/lib/proto/greet"
)

var (
	Flag     = flag.String("flag", "a", "flag")
	EtcdAddr = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
	Env      = flag.String("Env", "test", "env")
)

//rpc服务接口
type GreetServer struct{}

func (gs *GreetServer) Morning(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Printf("Morning 调用: %s\n", req.Name)
	return &proto.GreetResponse{
		Message: "Good morning, " + req.Name,
		From:    *Flag,
	}, nil
}

func (gs *GreetServer) Night(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Printf("Night 调用: %s\n", req.Name)
	return &proto.GreetResponse{
		Message: "Good night, " + req.Name,
		From:    *Flag,
	}, nil
}

func main() {
	flag.Parse()
	err := discovery.Init(*EtcdAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = discovery.GreetRegister(*Env, &GreetServer{})
	if err != nil {
		fmt.Println(err)
		return
	}
	discovery.WaitForClose()
}
