// Demon for manage connection pool of redis.
// [Reference](http://studygolang.com/articles/3029)
// Install `redis-view` by `go get github.com/dreamersdw/redis-view`.
// Show result by `redis-view`.
package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var MAX_POOL_SIZE = 20 // Pool size
var redisPoll chan redis.Conn

func putRedis(conn redis.Conn) {
	// 基于函数和接口间互不信任原则，这里再判断一次，养成这个好习惯哦
	if redisPoll == nil {
		redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
	}
	if len(redisPoll) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	redisPoll <- conn
}

func InitRedis(network, address string) redis.Conn {
	// 缓冲机制，相当于消息队列
	if len(redisPoll) == 0 {
		// 如果长度为0，就定义一个redis.Conn类型长度为MAX_POOL_SIZE的channel
		redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				c, err := redis.Dial(network, address)
				if err != nil {
					panic(err)
				}
				putRedis(c)
			}
		}()
	}
	return <-redisPoll
}

func redisServer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	c := InitRedis("tcp", "192.168.1.101:6379")
	dbkey := "netgame:info"
	if ok, err := redis.Bool(c.Do("LPUSH", dbkey, "server")); ok {
	} else {
		log.Print(err)
	}
	msg := fmt.Sprintf("用时：%s", time.Now().Sub(startTime))
	io.WriteString(w, msg)
}

func main() {
	var daemon, client bool
	flag.BoolVar(&daemon, "d", false, "Run as redis connection pool server.")
	flag.BoolVar(&client, "c", false, "Run as redis client.")
	flag.Parse()

	if daemon {
		// Server
		// go version <=1.3 利用cpu多核来处理http请求，这个没有用go默认就是单核处理http的
		// runtime.GOMAXPROCS(runtime.NumCPU())
		http.HandleFunc("/", redisServer)
		http.ListenAndServe(":9527", nil)

	} else if client {
		// Client
		resp, err := http.Get("http://192.168.1.101:9527/")
		if err != nil {
			fmt.Println("[ERROR] Can not connect to daemon, run daemon first.\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

	} else {
		flag.Usage()
	}
}
