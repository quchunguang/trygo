package main

/*
URL: https://github.com/mccoyst/myip/blob/master/myip.go
URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
*/

import (
	"net"
	"os"
)

// http.HandleFunc("/add", authenticator.Wrap(add))

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}
