package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/alphazero/Go-Redis"
	"github.com/quchunguang/trygo"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
func testhttpserv() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}

func testimage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func testredis() {
	spec := redis.DefaultSpec().Db(0).Password("")
	client, err := redis.NewSynchClientWithSpec(spec)
	if err != nil {
		fmt.Println("connect error", err)
		return
	}

	dbkey := "info"
	value, err := client.Get(dbkey)
	if err != nil {
		fmt.Println("Get error", err)
		return
	}

	if value == nil {
		value := []byte("Hello world!")
		client.Set(dbkey, value)
		fmt.Printf("插入数据>%s \n", value)
	} else {
		fmt.Printf("接收到数据>%s \n", value)
	}
}

func testfile() error {
	name := "test2.go"
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("Error")
		return err
	}
	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error")
		return err
	}
	fmt.Println(fileinfo.Size())
	return nil
}

type simplestack struct {
	i    int
	data [10]int
}

func (s *simplestack) push(k int) {
	if s.i >= 10 {
		return
	}
	s.data[s.i] = k
	s.i++
}
func (s *simplestack) pop() int {
	if s.i <= 0 {
		return 0
	}
	s.i--
	return s.data[s.i]
}

// impletement Stringer interface for print out
func (s simplestack) String() (ret string) {
	ret = "["
	for _, v := range s.data[:s.i] {
		ret += strconv.Itoa(v)
		ret += " "
	}
	ret += "]"
	return
}
func teststack() {
	var s simplestack
	fmt.Println(s)
	s.push(9)
	s.push(10)
	s.push(11)
	s.push(9)
	s.push(10)
	s.push(11)
	s.push(9)
	s.push(10)
	s.push(11)
	s.push(12)
	s.push(13)
	s.pop()
	fmt.Println(s)
}

func testpackage() {
	// 包名是小写的一个单词;不应当有下划线或混合大小写
	// import bar "bytes"
	// bar.Buffer()
	// % mkdir $GOPATH/src/example/even
	// % cp even.go $GOPATH/src/example/even
	// % go build
	// % go install
	fmt.Print(trygo.Even(2), trygo.Even(3))
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Xi []int
type Xs []string

func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}
func testsort() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}

func testio() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
func testio2() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()

	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
	}
}
func testio3() {
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	s, ok := r.ReadString('\n')
	if ok == nil {
		fmt.Println(s)
	}
}

func testexec() {
	cmd := exec.Command("/bin/ls", "-l")
	// err := cmd.Run()
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("error exec")
	}
	fmt.Print(string(buf))
}

func testnet() {
	// conn, e := Dial("tcp", "192.0.32.10:80")
	// conn, e := Dial("udp", "192.0.32.10:80")
	// conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80")
}
func testhttp() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err == nil {
		fmt.Printf("%s", string(b))
	}
}

////////
type Payload struct {
	Blog Blogs
}
type Blogs struct {
	Id          int
	User        Users
	Title       string
	Description string
	Modified    string
	Published   bool
	Taglist     []string
}
type Users struct {
	Id       int
	Name     string
	Username string
}

func getjson() ([]byte, error) {
	u := Users{1001, "qu", "Kevin Qu."}
	b := Blogs{10010001, u,
		"Being me", "How to be a president",
		"2009-03-17T03:53:36Z", true,
		[]string{"president", "usa", "john", "quincy", "adams"},
	}
	p := Payload{b}
	return json.MarshalIndent(p, "", "    ")
}
func testjson() {
	res, _ := getjson()
	fmt.Print(string(res))

	var pay Payload
	json.Unmarshal(res, &pay)
	fmt.Printf("\n\nUsername: %s    Title: %s\n",
		pay.Blog.User.Username, pay.Blog.Title)
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getjson()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(response))
}
func httpc() {
	time.Sleep(1e9)
	url := "http://localhost:1337"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var pay Payload
	err = json.Unmarshal(body, &pay)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Username: %s    Title: %s\n",
		pay.Blog.User.Username, pay.Blog.Title)
}
func httpd() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)
}
func testhttpjson() {
	go httpd()
	go httpc()
	time.Sleep(2e9)
}

func testsync() {
	done := make(chan bool)
	go func() {
		httpd()
		done <- true
	}()
	go func() {
		httpc()
		done <- true
	}()
	<-done
	<-done
}

func initConifg(once *sync.Once, handler func()) {
	once.Do(func() {
		time.Sleep(2e9)
		fmt.Println("2 second finished init.")
	})

	handler()

}
func testsync2() {
	var once sync.Once
	completeChan := []chan bool{make(chan bool, 1), make(chan bool, 1)}

	// MUST be pointer
	go initConifg(&once, func() {
		fmt.Println("init 1 channel!")
		completeChan[0] <- true
	})

	go initConifg(&once, func() {
		fmt.Println("init 2 channel!")
		completeChan[1] <- true
	})

	for _, ch := range completeChan {
		<-ch
		close(ch)
	}
}
func payload() {
	runtime.Gosched()
	fmt.Println("in payload()")
}
func testsync3() {
	runtime.GOMAXPROCS(4) // MUST
	waitchan := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			payload()
			waitchan <- i
		}()
	}
	// wait for 3 (at least) goroutine finished
	for i := 0; i < 3; i++ {
		<-waitchan
	}
}

func main() {
	// testhttpserv() // true loop by default!
	// testimage()
	// testredis() // need redis open
	// testfile()
	// teststack()
	// testpackage()
	// testsort()
	// testio()
	// testio2()
	// testio3()
	// testexec()
	// testnet()
	// testhttp() // need network connection
	// testjson()
	// testhttpjson() // true loop by default!
	// testsync()
	// testsync2()
	// testsync3()
}
