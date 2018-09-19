package gox

func IfElse(condition bool, ifVal, elseVal interface{}) interface{} {
	if condition {
		return ifVal
	} else {
		return elseVal
	}
}
