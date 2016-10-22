package trygo

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/alphazero/Go-Redis"
	"github.com/deckarep/golang-set"
	"html"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"mime"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"text/tabwriter"
	"time"
)

// Hello struct
type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// DemoHTTPServ func
func DemoHTTPServ() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}

// DemoImage func
func DemoImage() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0, 0).RGBA())
}

// DemoImage2 func
func DemoImage2() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 20; x < 80; x++ {
		y := x/3 + 15
		img.Set(x, y, color.Black)
	}
	w, _ := os.Create("data/Demoimage2.png")
	defer w.Close()
	png.Encode(w, img)
}

// ShowImage using command `display` of ImageMagick
func ShowImage(filename string) {
	cmd := exec.Command("display", filename)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// DrawLine func
func DrawLine(img *image.RGBA, x1, y1, x2, y2 int, color color.Color) {
	if abs(x2-x1) >= abs(y2-y1) {
		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		for x := x1; x <= x2; x++ {
			y := (x-x1)*(y2-y1)/(x2-x1) + y1
			img.Set(x, y, color)
		}
	} else {
		if y1 > y2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			x := (y-y1)*(x2-x1)/(y2-y1) + x1
			img.Set(x, y, color)
		}
	}
}

// DemoImage3 func
func DemoImage3() {
	filename := "data/testimage3.jpg"
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))
	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{0x88, 0xff, 0x88, 0xff})
		}
	}
	DrawLine(img, 100, 100, 600, 110, color.Black)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
	ShowImage(filename)
}

// DemoImage4 func
func DemoImage4() {
	cmd := exec.Command("date")
	buf, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}

// DemoRedis func
func DemoRedis() {
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
	var v []byte
	if value == nil {
		v = []byte("Hello world!")
		client.Set(dbkey, v)
		fmt.Printf("Input>%s \n", v)
	} else {
		fmt.Printf("Receive>%s \n", v)
	}
}

// DemoFile func
func DemoFile() error {
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

// DemoStack func
func DemoStack() {
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

// DemoPackage func
func DemoPackage() {
	// 包名是小写的一个单词;不应当有下划线或混合大小写
	// import bar "bytes"
	// bar.Buffer()
	// % mkdir $GOPATH/src/example/even
	// % cp even.go $GOPATH/src/example/even
	// % go build
	// % go install
	fmt.Print(Even(2), Even(3))
}

// Sorter interface
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Xi for sort
type Xi []int

// Xs for sort
type Xs []string

func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorter
func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

// DemoSort func
func DemoSort() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}

// DemoIo func
func DemoIo() {
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

// DemoIo2 func
func DemoIo2() {
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

// DemoIo3 func
func DemoIo3() {
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	s, ok := r.ReadString('\n')
	if ok == nil {
		fmt.Println(s)
	}
}

// DemoIo4 func
func DemoIo4() {
	r, _ := os.Open("/etc/passwd")
	w, _ := os.Create("/tmp/passwd")
	defer r.Close()
	defer w.Close()

	num, err := io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}

// DemoExec func
func DemoExec() {
	cmd := exec.Command("/bin/ls", "-l")
	// err := cmd.Run()
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("error exec")
	}
	fmt.Print(string(buf))
}

// DemoExec2 func
func DemoExec2() {
	var output bytes.Buffer
	cmd := exec.Command("cat")
	cmd.Stdout = &output
	stdin, _ := cmd.StdinPipe()
	cmd.Start()
	stdin.Write([]byte("widuu test"))
	stdin.Close()
	cmd.Wait()
	fmt.Printf("The output is: %s\n", output.Bytes())
}

// DemoExec3 func
func DemoExec3() {
	cmd := exec.Command("ls", "-ll")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	d, _ := ioutil.ReadAll(stdout)
	cmd.Wait()
	fmt.Println(string(d))
}

// DemoNet func
func DemoNet() {
	// conn, e := Dial("tcp", "192.0.32.10:80")
	// conn, e := Dial("udp", "192.0.32.10:80")
	// conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80")
}

// DemoHTTP func
func DemoHTTP() {
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

// Payload struct
type Payload struct {
	Blog Blogs
}

// Blogs struct
type Blogs struct {
	ID          int
	User        Users
	Title       string
	Description string
	Modified    string
	Published   bool
	Taglist     []string
}

// Users struct
type Users struct {
	ID       int
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

// DemoJSON func
func DemoJSON() {
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

// DemoHTTPJson func
func DemoHTTPJson() {
	go httpd()
	go httpc()
	time.Sleep(2e9)
}

// DemoSync func
func DemoSync() {
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

// DemoSync2 func
func DemoSync2() {
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

// DemoSync3 func
func DemoSync3() {
	runtime.GOMAXPROCS(4) // MUST only go version <= 1.4
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

// DemoSet from: github.com/deckarep/golang-set
func DemoSet() {
	rc := mapset.NewSet()
	rc.Add("Cooking")
	rc.Add("English")
	rc.Add("Math")
	rc.Add("Biology")
	ss := []interface{}{"Biology", "Chemistry"}
	sc := mapset.NewSetFromSlice(ss)
	ec := mapset.NewSet()
	ec.Add("Welding")
	ec.Add("Music")
	ec.Add("Automotive")
	bc := mapset.NewSet()
	bc.Add("Go Programming")
	bc.Add("Python Programming")
	//Show me all the available classes I can take
	allc := rc.Union(sc).Union(ec).Union(bc)
	fmt.Println(allc)
	//Is cooking considered a science class?
	fmt.Println(sc.Contains("Cooking"))
	//Show me all classes that are not science classes, since I hate science.
	fmt.Println(allc.Difference(sc))
	//Which science classes are also required classes?
	fmt.Println(sc.Intersect(rc))
	//How many bonus classes do you offer?
	fmt.Println(bc.Cardinality())
	fmt.Println(allc.IsSuperset(mapset.NewSetFromSlice(
		[]interface{}{"Welding", "Automotive", "English"})))
}

// Item struct
type Item struct {
	Key   string
	Value string
}

// DemoQueue func
func DemoQueue() {
	var v = Item{Key: "a", Value: "A"}
	queue := NewQueue()
	fmt.Println(queue.Enqueue(v).Value)
	v.Key = "b"
	v.Value = "B"
	fmt.Println(queue.Enqueue(v).Value)
	v.Key = "c"
	v.Value = "C"
	fmt.Println(queue.Enqueue(v).Value)
	v.Key = "d"
	v.Value = "D"
	fmt.Println(queue.Enqueue(v).Value)
	v.Key = "e"
	v.Value = "E"
	fmt.Println(queue.Enqueue(v).Value)

	fmt.Println(queue.Query(func(val interface{}) bool {
		if val.(Item).Key == "a" {
			return true
		}
		return false
	}).Value)
	fmt.Println(queue.Contain(v))
	fmt.Println(queue.Dequeue().Value)
	fmt.Println(queue.Dequeue().Value)
	fmt.Println(queue.Dequeue().Value)
	fmt.Println(queue.Dequeue().Value)
}

// type Interface interface {
//     Len() int
//     Less(i, j int) bool
//     Swap(i, j int)
// }
type intArray []int

func (s intArray) Len() int           { return len(s) }
func (s intArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }

// DemoSortAny func
func DemoSortAny() {
	a := []int{1, 5, 10, 4}
	sort.Sort(intArray(a))
}

// DemoSocket func
func DemoSocket() {
	var buf = make([]byte, 50)
	var n int
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("GET /\n\n"))

	for n, err = conn.Read(buf); n > 0; n, err = conn.Read(buf) {
		fmt.Print(string(buf[:n]))
	}
}

// DemoArgs func
func DemoArgs() {
	var li1 = []interface{}{"aa", 3}
	fmt.Println(li1...)

	var li2 []int
	for i := 0; i < 100; i++ {
		li2 = append(li2, i)
		fmt.Printf("i=%d len=%d cap=%d\n", i, len(li2), cap(li2))
	}
}

// DemoBufio func
func DemoBufio() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		line, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
	os.Exit(2)
}

// DemoRegexp func
func DemoRegexp() {
	var (
		before = "Test loop1 in loop1"
		after  string
	)
	if re, err := regexp.Compile("loop1"); err == nil {
		after = re.ReplaceAllString(before, "ASM001")
		fmt.Println(after)
		ret := re.FindAllString(before, -1)
		fmt.Println(ret)
	}
}

// DemoGetenv func
func DemoGetenv() {
	if val, found := syscall.Getenv("GOPATH"); found { // (1)
		fmt.Println(val)
	}
}

// DemoDeleteItem func
func DemoDeleteItem() {
	var li = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := len(li) - 1; i >= 0; i-- {
		if i%2 == 0 {
			li = append(li[:i], li[i+1:]...)
		}
	}
	fmt.Println(li[len(li):])
}

// MyStruct struct
type MyStruct struct {
	name string
}

// GetName func
func (s MyStruct) GetName() string {
	return s.name
}

// DemoReflect func
func DemoReflect() {
	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(s))
	var x = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)

	fmt.Println(typ.NumMethod())
	fmt.Println("-------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])
}

// FakeTime func
func FakeTime() {
	stop := time.After(5 * time.Second)
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			fmt.Println(time.Now())
		case <-stop:
			return
		}
	}
}

// DemoSignal func
func DemoSignal() {
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	fmt.Println("Before ctrl+c")
	<-sig
	fmt.Println("After ctrl+c")
}

// DemoSync4 func
func DemoSync4() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()

	fmt.Println("some_key:", n)
}

// DemoReflection func
func DemoReflection() {
	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("err")
	}
	r = tty
	fmt.Println(r)
	var w io.Writer
	w = r.(io.Writer)
	fmt.Println(w)

	var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)

	var xx uint8 = 'x'
	vv := reflect.ValueOf(xx)
	fmt.Println("type:", vv.Type())                            // uint8.
	fmt.Println("kind is uint8: ", vv.Kind() == reflect.Uint8) // true.
	xx = uint8(vv.Uint())                                      // v.Uint returns a uint64.

	type MyInt int
	var xxx MyInt = 7
	vvv := reflect.ValueOf(xxx)
	fmt.Println("type:", vvv.Type())                        // uint8.
	fmt.Println("kind is int: ", vvv.Kind() == reflect.Int) // true.

	yyy := vvv.Interface().(MyInt)
	fmt.Println(yyy)
	fmt.Printf("value is %d\n", vvv.Interface())
}

// DemoReflection2 func
func DemoReflection2() {
	var x = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

// DemoReflection3 func
func DemoReflection3() {
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
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

// DemoJSON2 func
func DemoJSON2() {
	type Message struct {
		Name string
		Body string
		Time int64
	}
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("errb, err := json.Marshal(m)")
	}
	fmt.Println(string(b))

	var mm Message
	json.Unmarshal(b, &mm)
	fmt.Println(mm)

	bb := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var mmm Message
	err = json.Unmarshal(bb, &mmm)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(mmm)
}

// DemoJSON3 func
func DemoJSON3() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("err")
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

// DemoJSON4 func
func DemoJSON4() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

// DemoBase64 func
func DemoBase64() {
	// A Buffer can turn a string or a []byte into an io.Reader.
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

// DemoBig func
func DemoBig() {
	i := new(big.Int)
	i.SetString("12345678901234567890", 10)
	fmt.Println(i)

	r := new(big.Rat)
	r.SetString("355/113")
	fmt.Println(r.FloatString(7))
}

// DemoHTTP2 func
func DemoHTTP2() {
	fmt.Println(html.EscapeString("<html>"))
}

// DemoSplit func
func DemoSplit() {
	var a []byte
	a = strconv.AppendQuote(a, "abc")
	fmt.Println(string(a))

	s := "a ,b   b, ccc"
	sa := strings.Split(s, ",")
	for i := range sa {
		sa[i] = strings.TrimSpace(sa[i])
	}
	fmt.Println(sa)
}

// DemoTime func
func DemoTime() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v\n", now)
	}
}

// DemoTabwriter func
func DemoTabwriter() {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
}

//////
type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// DemoSort2 func
func DemoSort2() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)
}

// DemoMime func
func DemoMime() {
	s := mime.TypeByExtension(".c")
	fmt.Println(s)
}

// DemoRegexp2 func
func DemoRegexp2() {
	fmt.Println(regexp.Match("H.* ", []byte("Hello World!")))

	r := bytes.NewReader([]byte("Hello World!"))
	fmt.Println(regexp.MatchReader("H.* ", r))

	fmt.Println(regexp.MatchString("H.* ", "Hello World!"))
	fmt.Println(regexp.QuoteMeta("(?P:Hello) [a-z]"))

	reg, err := regexp.Compile(`\w+`)
	fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)
	reg, err = regexp.CompilePOSIX(`[[:word:]]+`)
	fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)
	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindString("Hello World!"))
	reg = regexp.MustCompile(`\w+`)
	fmt.Printf("%q\n", reg.FindAll([]byte("Hello World!"), -1))

	reg = regexp.MustCompile(`(\w+),(\w+)`)
	src := "Golang,World!"
	dst := []byte("Say: ")
	template := "Hello $1, Hello $2"
	match := reg.FindStringSubmatchIndex(src)
	fmt.Printf("%q\n", reg.ExpandString(dst, template, src, match))

	text := `Hello World, 123 Go!`
	pattern := `(?U)H[\w\s]+o`
	reg = regexp.MustCompile(pattern)
	fmt.Printf("%q\n", reg.FindString(text))
	reg.Longest()
	fmt.Printf("%q\n", reg.FindString(text))

	reg = regexp.MustCompile(`(?U)(?:Hello)(\s+)(\w+)`)
	fmt.Println(reg.NumSubexp())

	b := []byte("Hello World, 123 Go!")
	reg = regexp.MustCompile(`(Hell|G)o`)
	rep := []byte("${1}ooo")
	fmt.Printf("%q\n", reg.ReplaceAll(b, rep))
	// "Hellooo World, 123 Gooo!"

	s := "Hello World!"
	reg = regexp.MustCompile("(H)ello")
	rep2 := "$0$1"
	fmt.Printf("%s\n", reg.ReplaceAllString(s, rep2))
	// HelloH World!
	fmt.Printf("%s\n", reg.ReplaceAllStringFunc(s,
		func(b string) string {
			return b + "$1"
		}))
}
