package model


type GinResp struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int64       `json:"code"`
}
