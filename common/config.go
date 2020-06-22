package common

import (
	"github.com/wudaoluo/sonic/model"
	"sync"
)

var once sync.Once
var config model.Config

func GetConf() *model.Config {
	once.Do(func() {
		config = model.Config{}
	})
	return &config
}