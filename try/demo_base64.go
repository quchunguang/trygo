package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Hello, 世界的就扫大街啊送idasdsdsdjsal收到啦时代科技实力肯定难受啦坑爹呢d家啊是滴哦撒"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	fmt.Println()
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	base64.RawStdEncoding
	fmt.Println(string(decoded))
}
