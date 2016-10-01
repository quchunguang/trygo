package trygo

import (
	"fmt"
)

// 自定义类型
type UserDef struct {
	What interface{}
}

func DemoInterface3() {
	// 接口类型
	var Interface interface{}

	// 赋值(自定义类型)
	Interface = UserDef{"World"}
	// 赋值(内置类型)
	// Interface = 2
	// 赋值(内置类型)
	// Interface = "3"

	// 输出
	switch Interface.(type) {
	case UserDef:
		fmt.Println(Interface.(UserDef).What.(string))
	case int:
		fmt.Println(Interface.(int))
	case string:
		fmt.Println(Interface.(string))
	}

	if value, ok := Interface.(UserDef); ok {
		fmt.Println(value.What.(string))
	}
}
