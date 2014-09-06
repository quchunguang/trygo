package main

// Usage:
// client:
// $ sudo nc 127.0.0.1:73
// {input username}
// {return $HOME/.plan}
import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os/user"
)

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", ":79")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	usr, _, _ := reader.ReadLine()
	if info, err := getUserInfo(string(usr)); err != nil {
		conn.Write([]byte(err.Error()))
	} else {
		conn.Write(info)
	}
}
func getUserInfo(usr string) ([]byte, error) {
	u, e := user.Lookup(usr)
	if e != nil {
		fmt.Println("[ERROR] user %s unknown!")
		return nil, e
	}
	data, err := ioutil.ReadFile(u.HomeDir + "/.plan")
	fmt.Println("[OK] user %s have a .plan file!")
	if err != nil {
		fmt.Println("[OK] user %s does not have a .plan file!")
		return data, errors.New("User doesn't have a .plan file!\n")
	}
	return data, nil
}
