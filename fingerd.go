package trygo

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

// Fingerd func
func Fingerd() {
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
		fmt.Printf("[ERROR] user %s unknown!\n", usr)
		return nil, e
	}
	data, err := ioutil.ReadFile(u.HomeDir + "/.plan")
	if err != nil {
		fmt.Printf("[OK] user %s does not have a .plan file!\n", usr)
		return data, errors.New("User doesn't have a .plan file!\n")
	}
	fmt.Printf("[OK] user %s have a .plan file!\n", usr)
	return data, nil
}
