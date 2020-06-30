package service


import (
	"fmt"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/model"
)



func init() {
	common.Cmd.Register(model.MsgReceive, &MsgReceive{})
}


type MsgReceive struct {

}

func (m *MsgReceive) Do(args interface{}) (interface{}, error) {
	fmt.Println("PostCommand")
	return args, nil
}