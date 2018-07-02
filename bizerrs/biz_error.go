package bizerrs

type BizError interface {
	error
	Code() string
}

type defaultErr struct {
	code string
	msg  string
}

func (err defaultErr) Error() string {
	return err.msg
}
func (err defaultErr) Code() string {
	return err.code
}

func New(code, msg string) BizError {
	return defaultErr{code, msg}
}
