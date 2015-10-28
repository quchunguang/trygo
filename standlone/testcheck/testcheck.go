package main

import (
	"fmt"
	"github.com/pengux/check"
	"strings"
)

type User struct {
	Username string
}

type CustomStringContainValidator struct {
	Constraint string
}

func (validator CustomStringContainValidator) Validate(v interface{}) check.Error {
	if !strings.Contains(v.(string), validator.Constraint) {
		return check.NewValidationError("customStringContainValidator", v, validator.Constraint)
	}

	return nil
}

func main() {
	u := &User{
		Username: "invalid*",
	}

	s := check.Struct{
		"Username": check.Composite{
			check.NonEmpty{},
			check.Regex{`^[a-zA-Z0-9]+$`},
			check.MinChar{10},
		},
	}

	e := s.Validate(u)

	if e.HasErrors() {
		err, ok := e.GetErrorsByKey("Username")
		if !ok {
			panic("key 'Username' does not exists")
		}
		fmt.Println(err)
	}

	// custom validator
	username := "invalid*"
	validator := CustomStringContainValidator{"admin"}
	e2 := validator.Validate(username)
	fmt.Println(check.ErrorMessages[e2.Error()])

	// custom error messages
	check.ErrorMessages["minChar"] = "the string must be minimum %v characters long"
	errMessages := e.ToMessages()
	fmt.Println(errMessages)
}
