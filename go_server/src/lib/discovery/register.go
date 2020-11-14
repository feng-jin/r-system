package discovery

import (
	"context"
	"errors"
	"fmt"
	"go_server/src/lib/discovery/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//Service 服务端用于服务注册的对象
type Service struct {
	Name string //服务名称
	Host string //{ip}:{port}
	Env  string //所属环境

	Key string //保存在etcd中的key
}

var service *Service

func (s *Service) register() error {
	if s.Env == "" {
		return errors.New("env is null")
	}
	s.Key = fmt.Sprintf(config.DirFormat, config.Scheme, s.Env, s.Name) + s.Host
	ticker := time.NewTicker(time.Second * time.Duration(config.TickerInterval))
	go func() {
		for {
			resp, err := client.Get(context.Background(), s.Key)
			if err != nil {
				fmt.Printf("获取服务地址失败：%s", err)
			} else if resp.Count == 0 { //尚未注册
				err = s.keepAlive()
				if err != nil {
					fmt.Printf("保持连接失败：%s", err)
				}
			}
			<-ticker.C
		}
	}()
	return nil
}

// keepAlive 创建租约，绑定，并续期
func (s *Service) keepAlive() error {
	//创建租约
	leaseResp, err := client.Grant(context.Background(), config.Expires)
	if err != nil {
		fmt.Printf("创建租期失败：%s\n", err)
		return err
	}

	//将服务地址注册到etcd中
	_, err = client.Put(context.Background(), s.Key, s.Host, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		fmt.Printf("注册服务失败：%s", err)
		return err
	}

	//租约续期
	ch, err := client.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		fmt.Printf("租约续期失败：%s\n", err)
		return err
	}

	//清空keepAlive返回的channel
	go func() {
		for {
			<-ch
		}
	}()
	return nil
}

//取消注册
func (s *Service) unRegister() {
	if client != nil {
		_, _ = client.Delete(context.Background(), s.Key)
	}
}

func WaitForClose() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	sig := <-ch
	service.unRegister()
	if i, ok := sig.(syscall.Signal); ok {
		os.Exit(int(i))
	} else {
		os.Exit(0)
	}
}
