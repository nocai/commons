package mdl

type Kvp struct {
	Key   interface{}
	Value interface{}
}

func (kvp Kvp) KeyInt() int {
	if kvp.Key == nil {
		return 0
	}
	return kvp.Key.(int)
}

func (kvp Kvp) KeyString() string {
	if kvp.Key == nil {
		return ""
	}
	return kvp.Key.(string)
}
func (kvp Kvp) KeyBool() bool {
	if kvp.Key == nil {
		return false
	}
	return kvp.Key.(bool)
}

func (kvp Kvp) ValueString() string {
	if kvp.Value == nil {
		return ""
	}
	return kvp.Value.(string)
}
