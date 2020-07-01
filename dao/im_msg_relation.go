package dao

import (
	"time"

	"github.com/go-xorm/xorm"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/model"
)

type ImMsgRelationService struct {
}

var DBImMsgRelation *ImMsgRelationService

func init() {
	DBImMsgRelation = &ImMsgRelationService{}
}

func (t *ImMsgRelationService) table() *xorm.Session {
	return db.Table("im_msg_relation")
}

func (t *ImMsgRelationService) Insert(data *model.ImMsgRelation) (int64, error) {
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
