package dao

import (
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/model"
)

type ImUserService struct {
	table string
}

var DBImUser *ImUserService

func init() {
	DBImUser = &ImUserService{
		table: "im_user",
	}
}

func (t *ImUserService) SelectByUserName(username string) (*model.ImUser, error) {
	msg := new(model.ImUser)
	_, err := db.Table(t.table).Where("username = ?", username).Get(msg)
	if err != nil {
		golog.Error("SelectByUserName", "err", err)
		return nil, err
	}

	if msg.Uid == 0 {
		err = common.DB_NOT_FOUND_ERR
		golog.Error("SelectByUserName", "func", "msg.Uid == 0", "err", err)
		return nil, err
	}

	return msg, nil
}
