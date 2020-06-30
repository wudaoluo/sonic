package model

import "time"

type Config struct {
	Default ConfigDefault
	Gateway ConfigGateway
	Auth ConfigAuth
	Logic ConfigLogic
	Queue ConfigQueue
	Cache ConfigCache
	Storage ConfigStorage
}

type ConfigDefault struct {
}

type ConfigGateway struct {

}

type ConfigAuth struct {
	Addr string
	Jwt  ConfigAuthJwt
}

type ConfigAuthJwt struct {
	Algorithm string
	Timeout time.Duration
	Key string
}

type ConfigLogic struct {
	Addr string
}

type ConfigQueue struct {
	Type string `json:"type"`
	Addrs []string `json:"addrs"`
}

type ConfigCache struct {
}

type ConfigStorage struct {
	MaxIdle int
	MaxOpen int
	Debug bool
	Addr string
}