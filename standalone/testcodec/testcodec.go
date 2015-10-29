package main

import (
	"bytes"
	"fmt"
	"github.com/ugorji/go/codec"
)

func main() {
	mh := &codec.MsgpackHandle{RawToString: true}
	data := []interface{}{"abc", 12345, 1.2345}
	buf := &bytes.Buffer{}
	enc := codec.NewEncoder(buf, mh)
	enc.Encode(data)
	fmt.Printf("%x", buf)
}
