package service

import (
	"fmt"
	"github.com/wudaoluo/sonic/common"
)



func init() {
	common.Cmd.Register(common.MsgReceive, &MsgReceive{})
}


type MsgReceive struct {

}

func (m *MsgReceive) Do(args interface{}) (interface{}, error) {
	fmt.Println("PostCommand")
	return args, nil
}


