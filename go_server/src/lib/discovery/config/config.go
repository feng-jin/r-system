package config

import "time"

// etcd
const (
	Timeout        = 15 * time.Second
	Expires        = 10
	TickerInterval = 5
	// scheme
	Scheme = "etcd"
	// etcd中存储key的格式前缀：/scheme/authority/endpoint
	DirFormat = "/%s/%s/%s/"
	// grpc resolver中自定义解析需要提供的格式：scheme://authority/endpoint
	// 其中scheme可以理解为解析策略，authority可以理解为权限管理，endpoint为地址
	TargetFormat = "%s://%s/%s"
)

// server name
const (
	GreetServer = "greet_server"
)
