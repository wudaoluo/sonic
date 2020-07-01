package dao

import (
	"time"

	"github.com/wudaoluo/sonic/common"

	"github.com/go-xorm/xorm"

	"github.com/wudaoluo/sonic/model"
)

type ImMsgContentService struct {
}

var DBImMsgContent *ImMsgContentService

func init() {
	DBImMsgContent = &ImMsgContentService{}
}

func (t *ImMsgContentService) table() *xorm.Session {
	return db.Table("im_msg_content")
}

func (t *ImMsgContentService) Insert(data *model.ImMsgContent) (int64, error) {
	data.CreateTime = time.Now()
	_, err := t.table().InsertOne(data)
	if err != nil {
		return 0, err
	}

	if data.Mid == 0 {
		return 0, common.DB_INSERT_ERR
	}

	return data.Mid, nil
}
