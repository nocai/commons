package bizerrs

type BizError interface {
	error
	Code() string

	Cause() error
}

type bizError struct {
	code string
	msg  string

	cause error
}

func (err bizError) Error() string {
	result := err.msg
	if err.cause != nil {
		result = result + ":" + err.cause.Error()
	}
	return result
}
func (err bizError) Code() string {
	return err.code
}
func (err bizError) Cause() error {
	return err.cause
}

func New(code, msg string) BizError {
	return bizError{code: code, msg: msg}
}

func Wrap(err BizError, cause error) BizError {
	return bizError{err.Code(), err.Error(), cause}
}
