package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	var o *Object

	// assert for nil (good for errors)
	assert.Nil(t, o)

	o = &Object{"Something"}
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, o) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", o.Value)

	}

}
