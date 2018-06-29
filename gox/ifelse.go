package gox

func Ifelse(condition bool, ifVal, elseVal interface{}) interface{} {
	if condition {
		return ifVal
	} else {
		return elseVal
	}
}
