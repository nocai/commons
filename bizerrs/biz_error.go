package bizerrs

type BizError interface {
	error
	Code() string
}

type bizError struct {
	code string
	msg  string
}

func (err bizError) Error() string {
	return err.msg
}
func (err bizError) Code() string {
	return err.code
}

func New(code, msg string) BizError {
	return bizError{code, msg}
}



