package assert

import "fmt"

// 断言
// if expression = true, panic the i
func Assert(expression bool, format string, i ...interface{}) {
	if expression {
		panic(fmt.Sprintf(format, i...))
	}
}
