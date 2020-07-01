package dao

import (
	"time"

	"github.com/wudaoluo/sonic/common"

	"github.com/go-xorm/xorm"
	"github.com/wudaoluo/sonic/model"
)

type ImMsgContactService struct {
}

var DBImMsgContact *ImMsgContactService

func init() {
	DBImMsgContact = &ImMsgContactService{}
}

func (t *ImMsgContactService) table() *xorm.Session {
	return db.Table("im_msg_contact")
}

func (t *ImMsgContactService) Insert(data *model.ImMsgContact) (int64, error) {
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

func (t *ImMsgContactService) Update(data *model.ImMsgContact) (int64, error) {
	return t.table().Where("owner_uid = ?", data.OwnerUid).
		And("other_uid = ?", data.OtherUid).Cols("mid", "type").Update(data)
}
