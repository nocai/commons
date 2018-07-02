package mdl

type Kvp struct {
	Key   interface{}
	Value interface{}
}

func (kvp Kvp) GetKInt() int {
	if kvp.Key == nil {
		return 0
	}
	return kvp.Key.(int)
}

func (kvp Kvp) GetKStr() string {
	if kvp.Key == nil {
		return ""
	}
	return kvp.Key.(string)
}
