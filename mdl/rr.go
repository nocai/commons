package mdl

import "time"

type Request struct {
	ReqNo string `form:"reqNo" binding:"required"`
}

type Response struct {
	ReqNo string      `json:"reqNo"`
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Now   time.Time   `json:"now"`
	Data  interface{} `json:"data"`
}
