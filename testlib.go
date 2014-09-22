package trygo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alphazero/Go-Redis"
	"github.com/deckarep/golang-set"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
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
func DemoHttpServ() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}

func DemoImage() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0, 0).RGBA())
}

func DemoImage2() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 20; x < 80; x++ {
		y := x/3 + 15
		img.Set(x, y, color.Black)
	}
	w, _ := os.Create("Demoimage2.png")
	defer w.Close()
	png.Encode(w, img)
}

// Show image using command `display` of ImageMagick
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
	} else {
		return x
	}
}
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
func DemoImage3() {
	filename := "testimage3.jpg"
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
	jpeg.Encode(file, img, &jpeg.Options{80})
	ShowImage(filename)
}
func DemoImage4() {
	cmd := exec.Command("date")
	buf, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
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

	if value == nil {
		value := []byte("Hello world!")
		client.Set(dbkey, value)
		fmt.Printf("Input>%s \n", value)
	} else {
		fmt.Printf("Receive>%s \n", value)
	}
}

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
func DemoSort() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}

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
func DemoIo3() {
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	s, ok := r.ReadString('\n')
	if ok == nil {
		fmt.Println(s)
	}
}

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
func DemoExec() {
	cmd := exec.Command("/bin/ls", "-l")
	// err := cmd.Run()
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("error exec")
	}
	fmt.Print(string(buf))
}
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
func DemoExec3() {
	cmd := exec.Command("ls", "-ll")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	d, _ := ioutil.ReadAll(stdout)
	cmd.Wait()
	fmt.Println(string(d))
}
func DemoNet() {
	// conn, e := Dial("tcp", "192.0.32.10:80")
	// conn, e := Dial("udp", "192.0.32.10:80")
	// conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80")
}
func DemoHttp() {
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
func DemoJson() {
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
func DemoHttpJson() {
	go httpd()
	go httpc()
	time.Sleep(2e9)
}

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
func DemoSync3() {
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

// from: github.com/deckarep/golang-set
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

//////
type Item struct {
	Key   string
	Value string
}

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
		} else {
			return false
		}
	}).Value)
	fmt.Println(queue.Contain(v))
	fmt.Println(queue.Dequeue().Value)
	fmt.Println(queue.Dequeue().Value)
	fmt.Println(queue.Dequeue().Value)
	fmt.Println(queue.Dequeue().Value)
}
