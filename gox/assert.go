package gox

// 断言
// if expression = false, panic the i
func Assert(expression bool, i interface{}) {
	if !expression {
		panic(i)
	}
}
