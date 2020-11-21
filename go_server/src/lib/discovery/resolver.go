package discovery

import (
	"context"
	"fmt"
	"go_server/src/lib/discovery/config"
	"strings"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

//EtcdResolver解析器
type EtcdResolver struct {
	dir        string
	clientConn resolver.ClientConn
}

func Resolver(env string, name string) *grpc.ClientConn {
	//注册etcd解析器
	r := &EtcdResolver{}
	resolver.Register(r)
	target := fmt.Sprintf(config.TargetFormat, r.Scheme(), env, name)
	//客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
	dailOpts := []grpc.DialOption{
		grpc.WithBalancerName("round_robin"), // grpc内部提供的轮询负载均衡
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024 * 1024 * 16),
		),
	}
	conn, err := grpc.Dial(target, dailOpts...)
	if err != nil {
		fmt.Println("连接服务器失败：", err)
	}
	return conn
}

func (r *EtcdResolver) Scheme() string {
	return config.Scheme
}

//构建解析器 grpc.Dial()同步调用
func (r *EtcdResolver) Build(target resolver.Target, clientConn resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	r.clientConn = clientConn
	r.dir = fmt.Sprintf(config.DirFormat, target.Scheme, target.Authority, target.Endpoint)
	go r.watch()
	return r, nil
}

//监听etcd中某个key前缀的服务地址列表的变化
func (r *EtcdResolver) watch() {
	//初始化服务地址列表
	var addrList []resolver.Address

	resp, err := client.Get(context.Background(), r.dir, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("获取服务地址列表失败：", err)
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), r.dir)})
		}
	}

	r.clientConn.NewAddress(addrList)

	//监听服务地址列表的变化
	rch := client.Watch(context.Background(), r.dir, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), r.dir)
			switch ev.Type {
			case clientv3.EventTypePut:
				if !exists(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.clientConn.NewAddress(addrList)
				}
			case clientv3.EventTypeDelete:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					r.clientConn.NewAddress(addrList)
				}
			}
		}
	}
}

func exists(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

//Close ...
func (r *EtcdResolver) Close() {}

//ResolveNow ...
func (r *EtcdResolver) ResolveNow(_ resolver.ResolveNowOption) {}
