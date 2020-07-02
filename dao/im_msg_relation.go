package dao

import (
	"time"

	"github.com/wudaoluo/golog"

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

func (t *ImMsgRelationService) FindByMid(param *model.ImMsgRelation) ([]*model.ImMsgRelation, error) {
	if param.OwnerUid == 0 || param.OtherUid == 0 {
		golog.Error("FindByMid", "OwnerUid", param.OwnerUid,
			"OtherUid", param.OtherUid, "mid", param.Mid, "err", common.DB_PARAM_ERROR)
		return nil, common.DB_PARAM_ERROR
	}

	var data []*model.ImMsgRelation

	if param.Mid == 0 {
		return data, nil
	}

	err := t.table().Where("owner_uid = ?", param.OwnerUid).
		And("other_uid = ?", param.OtherUid).
		And("mid > ?", param.Mid).Cols("mid", "type").Limit(10).OrderBy("-mid").
		Find(&data)

	if err != nil {
		golog.Error("FindByMid", "err", err)
		return nil, err
	}

	return data, nil
}
