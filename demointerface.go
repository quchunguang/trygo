package trygo

import (
	"fmt"
)

// 自定义类型
type Hello struct {
	What interface{}
}

func DemoInterface() {
	// 接口类型
	var Interface interface{}

	// 赋值(自定义类型)
	Interface = Hello{"World"}
	// 赋值(内置类型)
	// Interface = 2
	// 赋值(内置类型)
	// Interface = "3"

	// 输出
	switch Interface.(type) {
	case Hello:
		fmt.Println(Interface.(Hello).What.(string))
	case int:
		fmt.Println(Interface.(int))
	case string:
		fmt.Println(Interface.(string))
	}

	if value, ok := Interface.(Hello); ok {
		fmt.Println(value.What.(string))
	}
}
