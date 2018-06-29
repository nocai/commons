package mdl

type KV struct {
	Key   interface{}
	Value interface{}
}

func (kv KV) GetKInt() int {
	if kv.Key == nil {
		return 0
	}
	return kv.Key.(int)
}

func (kv KV) GetKStr() string {
	if kv.Key == nil {
		return ""
	}
	return kv.Key.(string)
}
