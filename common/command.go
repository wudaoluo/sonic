package common

import (
	"encoding/json"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/model"
)

var Cmd = &CommandHandler{CmdMap: make(map[model.CmdType]Command)}

// 命令接口
type Command interface {
	Do(args interface{}) (interface{}, error)
}

// 命令管理者
type CommandHandler struct {
	CmdMap map[model.CmdType]Command
}

// 处理命令
func (ch *CommandHandler) Handle(buf []byte) (interface{}, error) {
	if buf == nil {
		return nil, COMMAND_NIL
	}

	var param model.LogicCommand
	err := json.Unmarshal(buf,&param)
	if err != nil {
		golog.Error("Handle","err",err)
		return nil, err
	}

	cmd, ok := ch.CmdMap[param.MsgType]
	if ok {
		return cmd.Do(param.Data)
	}
	return nil, COMMAND_INVALID
}

// 注册命令
func (ch *CommandHandler) Register(CmdType model.CmdType, cmd Command) {
	ch.CmdMap[CmdType] = cmd
}



