package model

import "time"

//easygen: json   联系人列表：
type ImMsgContact struct {
	CreateTime time.Time `json:"create_time" xorm:"not null TIMESTAMP"`
	Mid        int64     `json:"mid" xorm:"not null INT"`
	OtherUid   int64     `json:"other_uid" xorm:"not null pk INT"`
	OwnerUid   int64     `json:"owner_uid" xorm:"not null pk INT"`
	Type       int       `json:"type" xorm:"not null INT"`
}

type MailBox int

const (
	InBox MailBox = iota
	OutBox
)

func (m MailBox) String() string {
	var value string
	switch m {
	case InBox:
		value = "收件箱"
	case OutBox:
		value = "发件箱"
	}
	return value
}

func (m MailBox) Int() int {
	return int(m)
}

//easygen: json
type ImMsgContent struct {
	Content     string    `json:"content" xorm:"not null VARCHAR(1000)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null TIMESTAMP"`
	Mid         int64     `json:"mid" xorm:"not null pk autoincr INT"`
	MsgType     int       `json:"msg_type" xorm:"not null INT"`
	RecipientId int64     `json:"recipient_id" xorm:"not null INT"`
	SenderId    int64     `json:"sender_id" xorm:"not null INT"`
}

type MsgType int

const (
	MsgText MsgType = iota + 1
)

func (m MsgType) String() string {
	var value string
	switch m {
	case MsgText:
		value = "文本消息"
	}
	return value
}

func (m MsgType) Int() int {
	return int(m)
}

//easygen: json 消息索引表：
type ImMsgRelation struct {
	CreateTime time.Time `json:"create_time" xorm:"not null TIMESTAMP"`
	Mid        int64     `json:"mid" xorm:"not null pk index(idx_owneruid_otheruid_msgid) INT"`
	OtherUid   int64     `json:"other_uid" xorm:"not null index(idx_owneruid_otheruid_msgid) INT"`
	OwnerUid   int64     `json:"owner_uid" xorm:"not null pk index(idx_owneruid_otheruid_msgid) INT"`
	Type       int       `json:"type" xorm:"not null INT"`
}

//easygen: json
type ImUser struct {
	Avatar   string `json:"avatar" xorm:"not null VARCHAR(500)"`
	Email    string `json:"email" xorm:"VARCHAR(250)"`
	Password string `json:"password" xorm:"not null VARCHAR(500)"`
	Uid      int64  `json:"uid" xorm:"not null pk INT"`
	Username string `json:"username" xorm:"not null VARCHAR(500)"`
}

type ImMsg struct {
	Content    string    `json:"content" xorm:"not null VARCHAR(1000)"`
	CreateTime time.Time `json:"create_time" xorm:"not null TIMESTAMP"`
	Mid        int64     `json:"mid" xorm:"not null pk autoincr INT"`
	MsgType    int       `json:"msg_type" xorm:"not null INT"`
	Type       int       `json:"type" xorm:"not null INT"`
}
