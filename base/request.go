package base

type RequestBase struct {
	ReqNo `form:"reqNo" binding:"required"`
}
