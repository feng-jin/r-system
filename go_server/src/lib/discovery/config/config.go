package config

import "time"

// etcd
const (
	Timeout        = 15 * time.Second
	Expires        = 10
	TickerInterval = 5

	Scheme       = "etcd"       // scheme
	DirFormat    = "/%s/%s/%s/" // /scheme/authority/endpoint
	TargetFormat = "%s://%s/%s" // scheme://authority/endpoint

)

const (
	GreetServer = "greet_server"
)
