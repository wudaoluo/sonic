package model

import "encoding/json"

type CmdType int

func (c CmdType) String() string {
	var value string
	switch c {
	case MsgReceive:
		value = "msg_receive"
	case MsgSend:
		value = "msg_send"
	default:
		value = "not found"
	}
	return value
}

const (
	MsgReceive CmdType = iota + 1   //消息接受
	MsgSend                         //消息发送
	ContactPersonList              //联系人列表
	MsgTotalUnRead                 //总未读书
	MsgUnRead                      //消息未读数

)


type LogicCommand struct {
	MsgType CmdType `json:"msg_type"`
	Data json.RawMessage `json:"data"`
}
