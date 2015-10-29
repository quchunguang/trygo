package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func runCmd(client *ssh.Client, cmd string) string {
	session, err := client.NewSession()
	checkErr(err)
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(cmd)
	checkErr(err)
	return b.String()
}

func main() {
	config := &ssh.ClientConfig{
		User: "pi",
		Auth: []ssh.AuthMethod{
			ssh.Password("qu-cg123"),
		},
	}
	client, err := ssh.Dial("tcp", "192.168.2.109:22", config)
	checkErr(err)

	var output string
	output = runCmd(client, "/usr/bin/whoami")
	fmt.Print(output)
	output = runCmd(client, "ls")
	fmt.Print(output)
}
