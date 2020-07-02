package service

import (
	"encoding/json"

	"github.com/wudaoluo/golog"

	"github.com/wudaoluo/sonic/dao"

	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/model"
)

func init() {
	common.Cmd.Register(model.MsgReceive, &MsgReceive{})
	common.Cmd.Register(model.MsgSend, &MsgSend{})
}

//easygen:
type MsgReceive struct {
	OwnerUid int64 `json:"owner_uid"`
	OtherUid int64 `json:"other_uid"`
	Mid      int64 `json:"mid"`
}

func (m MsgReceive) Do(args json.RawMessage) (interface{}, error) {
	err := json.Unmarshal(args, &m)
	if err != nil {
		golog.Error("MsgReceive.Do", "args", string(args), "err", err)
		return nil, err
	}

	param := model.ImMsgRelation{
		Mid:      m.Mid,
		OwnerUid: m.OwnerUid,
		OtherUid: m.OtherUid,
	}
	data, err := dao.DBImMsgRelation.FindByMid(&param)
	if err != nil {
		golog.Error("MsgReceive.Do", "func", "dao.DBImMsgRelation.FindByMid", "err", err)
		return nil, err
	}

	check := make(map[int64]int)
	var mids []int64
	for _, r := range data {
		check[r.Mid] = r.Type
		mids = append(mids, r.Mid)
	}

	contents, err := dao.DBImMsgContent.FindByMids(mids)
	if err != nil {
		return nil, err
	}

	for i, c := range contents {
		if v, ok := check[c.Mid]; ok {
			contents[i].Type = v
		}
	}

	return contents, nil
}

//easygen: json
type MsgSend struct {
	SenderId    int64         `json:"sender_id"`
	RecipientId int64         `json:"recipient_id"`
	Content     string        `json:"content"`
	MsgType     model.MsgType `json:"msg_type"`
}

//todo 改成事务处理
func (m MsgSend) Do(args json.RawMessage) (interface{}, error) {
	err := json.Unmarshal(args, &m)
	if err != nil {
		golog.Error("MsgSend.Do", "args", string(args), "err", err)
		return nil, err
	}
	data := model.ImMsgContent{
		MsgType:     model.MsgText.Int(),
		Content:     m.Content,
		SenderId:    m.SenderId,
		RecipientId: m.RecipientId,
	}
	mid, err := dao.DBImMsgContent.Insert(&data)
	if err != nil {
		golog.Error("msg", "MsgSend.Do", "err", err)
		return nil, err
	}

	imMsgcontactData := &model.ImMsgContact{
		Mid:      mid,
		OwnerUid: m.SenderId,
		OtherUid: m.RecipientId,
		Type:     model.InBox.Int(),
	}
	_, err = dao.DBImMsgContact.Update(imMsgcontactData)
	if err != nil {
		golog.Error("msg", "MsgSend.Do", "func", "dao.DBImMsgContact.Update", "err", err)
		return nil, err
	}

	imMsgcontactData.OwnerUid = m.RecipientId
	imMsgcontactData.OtherUid = m.SenderId
	imMsgcontactData.Type = model.OutBox.Int()
	_, err = dao.DBImMsgContact.Update(imMsgcontactData)
	if err != nil {
		golog.Error("msg", "MsgSend.Do", "func", "dao.DBImMsgContact.Update", "err", err)
		return nil, err
	}

	imMsgRelationData := &model.ImMsgRelation{
		Mid:      mid,
		OwnerUid: m.SenderId,
		OtherUid: m.RecipientId,
		Type:     model.InBox.Int(),
	}

	_, err = dao.DBImMsgRelation.Insert(imMsgRelationData)
	if err != nil {
		golog.Error("msg", "MsgSend.Do", "func", "dao.DBImMsgContact.Update", "err", err)
		return nil, err
	}

	imMsgRelationData.OwnerUid = m.RecipientId
	imMsgRelationData.OtherUid = m.SenderId
	imMsgRelationData.Type = model.OutBox.Int()
	_, err = dao.DBImMsgRelation.Insert(imMsgRelationData)
	if err != nil {
		golog.Error("msg", "MsgSend.Do", "func", "dao.DBImMsgContact.Update", "err", err)
		return nil, err
	}

	return true, nil
}
