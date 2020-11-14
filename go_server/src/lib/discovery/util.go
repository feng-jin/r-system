package discovery

import (
	"fmt"
	"net"
)

// 获取本机ip地址
func getIntranetIP() (ip string) {
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
					break
				}
			}
		}
	}
	return
}

// 自动获取本机的ip以及端口号，ip:port格式
func getListener() (listener net.Listener, host string, err error) {
	host = "0.0.0.0:0"
	listener, err = net.Listen("tcp", host)
	if err == nil {
		addr := listener.Addr().String()
		_, portString, _ := net.SplitHostPort(addr)
		host = fmt.Sprintf("%s:%s", getIntranetIP(), portString)
	}
	return
}
