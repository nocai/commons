package mdl

import "time"

type Request struct {
	ReqNo string `form:"reqNo" binding:"required"`
}

type Response struct {
	RespNo string    `json:"respNo"`
	Code   string    `json:"code"`
	Msg    string    `json:"msg"`
	Now    time.Time `json:"now"`
	Data   string    `json:"data"`
}
