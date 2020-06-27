package common


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


var Cmd = &CommandHandler{CmdMap: make(map[CmdType]Command)}

// 命令接口
type Command interface {
	Do(args interface{}) (interface{}, error)
}

// 上下文
type CmdContext struct {
	CmdType CmdType
	Args    interface{}
}

// 命令管理者
type CommandHandler struct {
	CmdMap map[CmdType]Command
}

// 处理命令
func (ch *CommandHandler) Handle(ctx *CmdContext) (interface{}, error) {
	if ctx == nil {
		return nil, COMMAND_NIL
	}
	cmd, ok := ch.CmdMap[ctx.CmdType]
	if ok {
		return cmd.Do(ctx.Args)
	}
	return nil, COMMAND_INVALID
}

// 注册命令
func (ch *CommandHandler) Register(cmdType CmdType, cmd Command) {
	ch.CmdMap[cmdType] = cmd
}



