package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	EtcdAddress  = []string{"http://127.0.0.1:2379"}
	ServerPrefix = "/test/server/"
)

type EtcdClient struct {
	address  []string
	username string
	password string
	kv       clientv3.KV
	client   *clientv3.Client
	ctx      context.Context
	lease    clientv3.Lease
	leaseID  clientv3.LeaseID
}

func newEtcdClient() *EtcdClient {
	var client = &EtcdClient{
		ctx:     context.Background(),
		address: EtcdAddress,
	}
	err := client.connect()
	if err != nil {
		panic(err)
	}
	return client
}

func (etcdClient *EtcdClient) connect() (err error) {
	etcdClient.client, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdClient.address,
		DialTimeout: 5 * time.Second,
		TLS:         nil,
		Username:    etcdClient.username,
		Password:    etcdClient.password,
	})
	if err != nil {
		return
	}
	etcdClient.kv = clientv3.NewKV(etcdClient.client)
	etcdClient.ctx = context.Background()
	return
}

func (etcdClient *EtcdClient) list(prefix string) ([]string, error) {
	resp, err := etcdClient.kv.Get(etcdClient.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	servers := make([]string, 0)
	for _, value := range resp.Kvs {
		if value != nil {
			servers = append(servers, string(value.Value))
		}
	}
	return servers, nil
}

func (etcdClient *EtcdClient) close() (err error) {
	return etcdClient.client.Close()
}

func genRand(num int) int {
	return int(rand.Int31n(int32(num)))
}

func getServer(client *EtcdClient) (string, error) {
	servers, err := client.list(ServerPrefix)
	if err != nil {
		return "", err
	}
	return servers[genRand(len(servers))], nil
}

func Get(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	client := newEtcdClient()
	err := client.connect()
	if err != nil {
		panic(err)
	}
	defer client.close()

	for i := 0; i < 10; i++ {
		address, err := getServer(client)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data, err := Get(address + "ping")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
		time.Sleep(2 * time.Second)
	}
}
