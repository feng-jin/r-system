package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"
)

const (
	EtcdPrefix   = "/test/server/"
	ServerSerial = "1"
	Address      = "http://127.0.0.1:18081/"
)

var (
	EtcdAddress = []string{"http://127.0.0.1:2379"}
	leaseTTL    = 5
)

type HealthProvider struct {
	etcdClient *EtcdClient
}

var (
	healthProvider     *HealthProvider
	healthProviderOnce sync.Once
)

func GetHealthProvider() *HealthProvider {
	healthProviderOnce.Do(func() {
		healthProvider = &HealthProvider{
			etcdClient: NewEtcdClient(),
		}
	})
	return healthProvider
}

type EtcdClient struct {
	address  []string
	username string
	password string
	kv       clientv3.KV
	client   *clientv3.Client
	ctx      context.Context
	lease    clientv3.Lease
	leaseID  clientv3.LeaseID
	leaseTTL int64
}

func NewEtcdClient() *EtcdClient {
	var client = &EtcdClient{
		ctx:      context.Background(),
		address:  EtcdAddress,
		leaseTTL: int64(leaseTTL),
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

func (etcdClient *EtcdClient) Close() (err error) {
	return etcdClient.client.Close()
}

func (etcdClient *EtcdClient) register(address string) (*clientv3.PutResponse, error) {
	etcdClient.lease = clientv3.NewLease(etcdClient.client)
	leaseResp, err := etcdClient.lease.Grant(etcdClient.ctx, etcdClient.leaseTTL)
	if err != nil {
		return nil, err
	}
	etcdClient.leaseID = leaseResp.ID
	return etcdClient.kv.Put(etcdClient.ctx, EtcdPrefix+ServerSerial, address, clientv3.WithLease(leaseResp.ID))
}

func (etcdClient *EtcdClient) LeaseKeepAlive() error {
	if etcdClient.lease == nil {
		_, err := etcdClient.register(Address)
		if err != nil {
			return err
		}
	}
	_, err := etcdClient.lease.KeepAlive(etcdClient.ctx, etcdClient.leaseID)
	if err != nil {
		return err
	}
	return nil
}

func healthCheck(provider *HealthProvider) {
	var tick = time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			err := provider.etcdClient.LeaseKeepAlive()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func main() {

	provider := GetHealthProvider()
	go healthCheck(provider)

	defer provider.etcdClient.Close()

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "one")
	})

	engine.Run(":18081")
}
