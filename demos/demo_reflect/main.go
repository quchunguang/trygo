package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)

	var w io.Writer
	w = r.(io.Writer)

	var e interface{}
	e = w

	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.ValueOf(e))

	type Myfloat float64
	var f Myfloat = 3.4
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)
	fmt.Println("type", tf)  // func TypeOf(i interface{}) Type
	fmt.Println("value", vf) // func ValueOf(i interface{}) Value
	fmt.Println("value's type", vf.Type())
	fmt.Println("valus's kind", vf.Kind())         // constant indicating what sort of item is stored
	fmt.Printf("value is %7.1e\n", vf.Interface()) // no need v.Interface().(Myfloat)

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // Error: will panic. We pass a **copy** of x to reflect.ValueOf.
	fmt.Println("settability of v:", v.CanSet())

	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())  // p build from the copy of &x
	vv := p.Elem()                                // get the pointer actual pointing to (x)
	fmt.Println("settability of vv", vv.CanSet()) // x here is settable now
	vv.SetFloat(7.1)
	fmt.Println(vv.Interface())
	fmt.Println(x)

	// structure case
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
