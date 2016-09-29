package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func main() {
	fmt.Println(govalidator.IsURL(`http://domain.com/path/page`))

	type User struct {
		FirstName string
		LastName  string
	}

	str := govalidator.ToString(&User{"John", "Juan"})
	println(str)

	data := []interface{}{1, 2, 3, 4, 5}
	var fn govalidator.Iterator = func(value interface{}, index int) {
		println(value.(int))
	}
	govalidator.Each(data, fn)
}
