package main

/*
#include <stdio.h>

*/
import "C" // this MUST be single sentence with magic omments above !!!
import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/alphazero/Go-Redis"
	"github.com/quchunguang/testgo"
	"image"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no"

const Pi = 3.14
const (
	Big   = 1 << 60
	Small = Big >> 59
)

type Vertex struct {
	Lat, Long float64
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func divmod(x, y int) (div, mod int) {
	div = x / y
	mod = x % y
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 不能在这里使用 v，因此
	return lim
}

func testdefine() {
	m, n, p := 1, 2, 3
	const World = "世界"
	const Truth bool = true
	fmt.Println("hello world!")
	fmt.Println("Happy", math.Pi, "Day")
	fmt.Println("Hi, 3 + 4 =", add(3, 4))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(divmod(10, 3))
	fmt.Println(x, y, z, c, python, java)
	fmt.Println(m, n, p, World, Truth)
	fmt.Println(Big, Small, needInt(Small), needFloat(Small), needFloat(Big))
}

func testfor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	a := []int{4, 3, 5, 2, 1}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for _, n := range a {
		fmt.Println(n)
	}
	for pos, char := range "aΦx" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}
}

func testif() {
	fmt.Println(sqrt(-82))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func teststruct() {
	// var v Vertex = Vertex{1, 2}
	v := Vertex{1, 2}
	w := &v
	v.Lat = 4
	w.Long = 4
	fmt.Println(v)
	fmt.Println(w)
	o := new(Vertex) // type *Vertex
	o.Lat, o.Long = 11, 9
	fmt.Println(o)
}

func testmap1() {
	mvex := make(map[string]Vertex)
	mvex["Bell Labs"] = Vertex{40.68433, 74.39967}
	fmt.Println(mvex["Bell Labs"])
	mvex = nil
	// mvex["Bell Labs"] = Vertex{40.68433, 74.39967} // error
	fmt.Println(mvex["Bell Labs"])

	var nvex = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(nvex)
}

func testmap2() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func testslice() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])

	for i := 0; i < len(p); i++ {
		fmt.Println(i, p[i])
	}
}

func testslice2() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
	a = a[2:5]
	printSlice("a", a)
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func testfunc() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testclosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func testforrange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

func fromfun() int {
	fmt.Print("In fun - ")
	return 1
}

func testswitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}

	switch weekday := time.Now().Weekday(); weekday {
	case time.Monday:
		fmt.Println(time.Monday)
	case time.Sunday:
		fmt.Println(time.Sunday)
	}

	switch count := 2; count {
	case 0:
		fmt.Println("0")
	case fromfun():
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	i := 0
	switch i {
	case 0:
		fallthrough
	case 1:
		Sqrt(0)
	}
}

func Sqrt(x float64) float64 {
	var a, b, s float64 = 1, x, 0
	for i := 0; i < 50; i++ {
		s = (a + b) / 2
		// fmt.Println(i, a, b, s)
		if s*s < x {
			a = s
		} else {
			b = s
		}
	}
	return s
}

func exercise_sqrt() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}

func wordcount(s string) map[string]int {
	var wc = make(map[string]int)
	for _, v := range strings.Fields(s) {
		// fmt.Println(i, v)
		wc[v]++
	}
	return wc
}

func exercise_wc() {
	wc := wordcount("a b c a a e d e")
	for i, v := range wc {
		fmt.Println(i, v)
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	fun := func() int {
		a, b = b, a+b
		return a
	}
	return fun
}

func exercise_fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

func testmethod() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func testmethod2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (v Vertex) Notchange() {
	v.Long = 0
}

func testref() {
	v := Vertex{3, 4}
	v.Notchange()
	fmt.Println(v)
}

type Abser interface {
	Abs() float64
}

func testinterface() {
	var ia Abser
	f := MyFloat(-math.Sqrt2)
	ia = f
	v := &Vertex{3, 4}
	ia = v

	fmt.Println(ia.Abs())
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func testinterface2() {
	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it did not work!",
	}
}
func testerror() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

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

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Input %f: can not be negative", e)
}
func SqrtE(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	} else {
		return math.Sqrt(f), nil
	}
}
func exercise_error() {
	fmt.Println(SqrtE(2))
	fmt.Println(SqrtE(-2))
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := range p {
		if p[i] >= 65 && p[i] <= 90 {
			// upper case
			p[i] = (p[i]-65+13)%26 + 65
		} else if p[i] >= 97 && p[i] <= 122 {
			// lower case
			p[i] = (p[i]-97+13)%26 + 97
		}
	}
	return
}
func exercise_ioreader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func testgoroutine() {
	go say("world")
	say("hello")
}
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
func testchannel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
func testchannel2() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func testchannel3() {
	c := make(chan int, 10)
	go fibonacci2(80, c)
	for i := range c {
		fmt.Println(i)
	}
}
func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func testchannel4() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}
func testchannel5() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Print(".")
			time.Sleep(5e7)
		}
	}
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func walk_preorder_recursive(t *Tree, c chan int) {
	c <- t.Value
	if t.Left != nil {
		walk_preorder_recursive(t.Left, c)
	}
	if t.Right != nil {
		walk_preorder_recursive(t.Right, c)
	}
}
func walk_inorder_recursive(t *Tree, c chan int) {
	if t.Left != nil {
		walk_inorder_recursive(t.Left, c)
	}
	c <- t.Value
	if t.Right != nil {
		walk_inorder_recursive(t.Right, c)
	}
}

func walk_postorder_recursive(t *Tree, c chan int) {
	if t.Left != nil {
		walk_postorder_recursive(t.Left, c)
	}
	if t.Right != nil {
		walk_postorder_recursive(t.Right, c)
	}
	c <- t.Value
}
func walk_levelorder_recursive(t *Tree, c chan int) {
}
func walk_preorder(t *Tree, c chan int) {
	s := testgo.NewStack()
	s.Push(t)
	var h *Tree
	for !s.Empty() {
		h = s.Pop().(*Tree)
		c <- h.Value
		if h.Right != nil {
			s.Push(h.Right)
		}
		if h.Left != nil {
			s.Push(h.Left)
		}
	}
}
func walk_inorder(t *Tree, c chan int) {

}
func walk_postorder(t *Tree, c chan int) {

}
func walk_levelorder(t *Tree, c chan int) {

}
func compare_chan(c1, c2 chan int) {
	var ic1, ic2 int
	var has_c1, has_c2 bool
	for {
		ic1, has_c1 = <-c1
		ic2, has_c2 = <-c2

		if has_c1 == false && has_c2 == false {
			fmt.Println("EQUARE.")
			return
		}
		if has_c1 == false || has_c2 == false || ic1 != ic2 {
			fmt.Println("NOT EQUARE.")
			return
		}
	}
}
func exercise_checktree() {
	// create trees and init chans
	t1 := &Tree{&Tree{&Tree{nil, 1, nil}, 1, &Tree{nil, 2, nil}}, 3, &Tree{&Tree{nil, 5, nil}, 8, &Tree{nil, 13, nil}}}
	t2 := &Tree{&Tree{&Tree{&Tree{nil, 1, nil}, 1, &Tree{nil, 2, nil}}, 3, &Tree{nil, 5, nil}}, 8, &Tree{nil, 13, nil}}
	c1 := make(chan int)
	c2 := make(chan int)

	// begin generate chans
	go func() {
		walk_inorder_recursive(t1, c1)
		close(c1)
	}()
	go func() {
		walk_inorder_recursive(t2, c2)
		close(c2)
	}()

	// begin compare chans
	compare_chan(c1, c2)
}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func exercise_craw() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func teststring() {
	str := "LaoYing老鹰"
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
	for i, s := range str {
		fmt.Println(i, "Unicode(", s, ") string=", string(s))
	}
	r := []rune(str)
	fmt.Println("rune=", r)
	for i := 0; i < len(r); i++ {
		fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}
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
func testtype() {
	const (
		A = iota
		B
		C string = "string"
	)
	fmt.Println(A, B, C)

	var a int32 = 20
	var b int
	b = int(a)
	fmt.Println(b)

	s1 := "lkfjldsfkd" +
		"fjdsflkjds"
	s2 := `sds'"!@#$%^&*(asadsads
dsadsaddsa`
	fmt.Println(s1, s2)

	var c complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", c)

	var err error
	fmt.Println(err)
	var x, y, z int
	fmt.Println(x, y, z)
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
func testgoto() {
	i := 0
Here:
	fmt.Println(i)
	if i++; i == 5 {
		return
	}
	goto Here
}
func testdefer() (ret int) {
	defer func() {
		ret = 2
	}()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	f, err := os.Open("notexist.go")
	defer f.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 1
	}
	return 0
}
func myfunc2(arg ...interface{}) {

}
func myfunc1(arg ...int) {

}
func testvarargs(arg ...int) {
	myfunc1(arg...)
	myfunc1(arg[:2]...)
	myfunc2(arg)
	myfunc2(4, true, "abc")
}
func testfuncvalue() {
	f := func(name string) {
		fmt.Printf("Hello %s\n", name)
	}
	f("Qu")

	var xs = map[int]func() int{
		1: func() int { fmt.Println("10"); return 10 },
		2: func() int { fmt.Println("20"); return 20 },
		3: func() int { fmt.Println("30"); return 30 },
	}
	xs[2]()
}
func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}
func testmap() {
	double := func(a int) int { return 2 * a }
	l := []int{1, 2, 3, 4, 5}
	m := Map(double, l)
	fmt.Println(m)
}
func f() {
	// panic(interface{})是一个内建函数,可以中断原有的控制流程,进入一个令人恐慌的流程中。当函数 F 调用 panic ,函数 F 的执行被中断,并且 F 中的延迟函数会正常执行,然后 F 返回到调用它的地方。在调用的地方, F 的行为就像调用了 panic 。这一过程继续向上,直到程序崩溃时的所有 goroutine 返回。恐慌可以直接调用 panic 产生。也可以由 运行时错误 产生,例如访问越界的数组。
	panic("i'm panic!")
	fmt.Println("Hi, panic?")
}
func throwsPanic(f func()) (b bool) {
	// recover()是一个内建的函数,可以让进入令人恐慌的流程中的 goroutine 恢复过来。仅在延迟函数中有效。
	defer func() {
		if x := recover(); x != nil {
			b = true
			fmt.Println(x.(string))
		}
	}()
	f()
	return
}
func testpanic() {
	var ret bool = throwsPanic(f)
	fmt.Println(ret)
}

func exercise_functions() {
	d := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum float64
	for _, v := range d {
		sum += v
	}
	fmt.Printf("Average = %v\n", sum/float64(len(d)))
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
func printthem(them ...int) {
	for _, d := range them {
		fmt.Println(d)
	}
}
func testvararg2() {
	printthem(1, 4, 5, 7, 4)
	printthem(1, 2, 4)
}
func testpackage() {
	// 包名是小写的一个单词;不应当有下划线或混合大小写
	// import bar "bytes"
	// bar.Buffer()
	// % mkdir $GOPATH/src/example/even
	// % cp even.go $GOPATH/src/example/even
	// % go build
	// % go install
	fmt.Print(testgo.Even(2), testgo.Even(3))
}
func testtest() {
	// see example/even
	// cd $GOPATH/src/example/even
	// go test
}
func commonpkg() {
	// fmt
	// %v 默认格式的值。当打印结构时,加号( %+v )会增加字段名;
	// %#v Go 样式的值表达;
	// %T 带有类型的 Go 样式的值表达;

	// io
	// 这个包提供了原始的 I/O 操作界面。它主要的任务是对os包这样的原始的 I/O 进行封装,增加一些其他相关,使其具有抽象功能用在公共的接口上。

	//bufio
	//这个包实现了缓冲的 I/O 。它封装于 io.Reader 和 io.Writer 对象,创建了另一个对象( Reader 和 Writer )在提供缓冲的同时实现了一些文本 I/O 的功能。

	// sort
	// sort 包提供了对数组和用户定义集合的原始的排序功能。

	// strconv
	// strconv 包提供了将字符串转换成基本数据类型,或者从基本数据类型转换为字符串的功能。

	// os
	// os 包提供了与平台无关的操作系统功能接口。其设计是 Unix 形式的。

	// sync
	// sync 包提供了基本的同步原语,例如互斥锁。

	// flag
	// flag 包实现了命令行解析。

	// encoding/json
	// encoding/json 包实现了编码与解码 RFC 4627 定义的 JSON 对象。

	// text/template
	// 数据驱动的模板,用于生成文本输出,例如 HTML 。将模板关联到某个数据结构上进行解析。模板内容指向数据结构的元素(通常结构的字段或者 map 的键)控制解析并且决定某个值会被显示。模板扫描结构以便解析,而 “ 游标 ” @ 决定了当前位置在结构中的值。

	// net/http
	// net/http 实现了 HTTP 请求、响应和 URL 的解析,并且提供了可扩展的 HTTP服务和基本的 HTTP 客户端。

	// unsafe
	// unsafe 包包含了 Go 程序中数据类型上所有不安全的操作。 通常无须使用这个。

	// reflect
	// reflect 包实现了运行时反射,允许程序通过抽象类型操作对象。通常用于处理静态类型 interface{} 的值,并且通过 Typeof 解析出其动态类型信息,通常会返回一个有接口类型 Type 的对象。

	// os/exec
	// os/exec 包执行外部命令。
}

const (
	Enone  = 1
	Einval = 2
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func testpointer() {
	var p *int
	fmt.Printf("%v\n", p)

	var i int = 9
	p = &i
	*p = 8
	fmt.Printf("%v %v %v\n", p, *p, i)
	*p++ // means (*p)++
	fmt.Printf("%v\n", i)

	m := new(SyncedBuffer)
	var n SyncedBuffer
	fmt.Println(m, n)

	ar := [...]string{Enone: "no error", Einval: "invalid argument"}
	sl := []string{Enone: "no error", Einval: "invalid argument"}
	ma := map[int]string{Enone: "no error", Einval: "invalid argument"}
	fmt.Println(ar, sl, ma)

}

type foo int
type NameAge struct {
	name string
	age  int
}

func testcustomtype() {
	a := new(NameAge)
	a.name = "Pete"
	a.age = 42
	fmt.Printf("%v\n", a)

	// var n NameAge
	// n.dosomething    // x.m() == (&x).m()

	type Mutex int
	// NewMutex is Mutex, but no methods of Mutex
	type NewMutex Mutex
	// methods of Mutex is binding to PrintableMutex's no-name field Mutex
	type PrintableMutex struct{ Mutex }
	mystring := "UTF-8 编码"
	byteslice := []byte(mystring)
	runeslice := []rune(mystring)
	fmt.Println(byteslice, runeslice)

	b := []byte{'h', 'e', 'l', 'l', 'o'}
	s := string(b)
	i := []rune{257, 1024, 65}
	r := string(i)
	fmt.Println(b, s, i, r)

	type foo struct{ int }
	type bar foo
	var x bar = bar{1}
	var y foo = foo(x)
	fmt.Println(x, y)

}

///////////
//* define the empty interface as a type
type e interface{}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}
func Map3(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}
func testmap3() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	mf := Map3(m, mult2)
	sf := Map3(s, mult2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

///////////
type I interface {
	Get() int
	Put(int)
}
type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

func cheacktype(p I) {
	switch p.(type) {
	case *S:
		fmt.Println("*S")
	case *R:
		fmt.Println("*R")
		// S, R are not I, while *S, *R does !!!
		// case S:
		// 	fmt.Println("S")
		// case R:
		// 	fmt.Println("R")
	}
}
func testtype2() {
	var s S
	var ps *S
	ps = &s
	cheacktype(ps)
	s.Put(2)
	fmt.Println(s)

	// fmt.Println(g(s))  S are NOT I
	fmt.Println(g(ps))
}
func g(something interface{}) int {
	return something.(I).Get()
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

////////
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

///////////////
type Interface1 interface {
	Len(x interface{})
}
type Interface2 interface {
	Interface1
	Push(x interface{})
	Pop() interface{}
}

func testinterface_recurve() {

}

///////////////
type Person struct {
	name string "namestr" // Tag
	age  int
}

func ShowTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
	}
	switch i.(type) {
	case *Person:
		tt := reflect.TypeOf(i)
		vv := reflect.ValueOf(i)
		tag := tt.Elem().Field(0).Tag
		name := vv.Elem().Field(0).String()

		fmt.Println("tag: ", tag, "\nname: ", name)
	}
}
func testintrospection() {
	p1 := new(Person)
	p1.name = "guang"
	ShowTag(p1)
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}
func testgoroutine2() {
	go ready("Tea", 1)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(2 * time.Second)
	gomaxprocs := runtime.GOMAXPROCS(4) // or set env GOMAXPROCS
	gomaxprocs = runtime.GOMAXPROCS(gomaxprocs)
	fmt.Println(gomaxprocs)
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
func dnslookup() {
	dnssec := flag.Bool("dnssec", false, "Request DNSSEC records")
	port := flag.String("port", "53", "Set the query port")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *dnssec {
	}
	if *port == "53" {

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

func psgrp() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, _ := ps.Output()
	child := make(map[int][]int)
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 || len(s) == 0 {
			continue
		}
		f := strings.Fields(s)
		fp, _ := strconv.Atoi(f[0])
		fpp, _ := strconv.Atoi(f[1])
		child[fpp] = append(child[fpp], fp)
	}
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
		if len(child[ppid]) == 1 {
			fmt.Printf(": %v\n", child[ppid])
			continue
		}
		fmt.Printf("ren: %v\n", child[ppid])
	}
}

func wc() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin)
	for {
		switch s, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("%d %d %d\n", chars, words, lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
func uniq() {
	list := []string{"a", "b", "a", "a", "c", "d", "e", "f"}
	first := list[0]
	fmt.Printf("%s ", first)
	for _, v := range list[1:] {
		if first != v {
			fmt.Printf("%s ", v)
			first = v
		}
	}
}

// Usage
//     $ ./test		    # server side
//     $ nc 127.0.0.1 8053  # client side
//     abc
//     abc
func echoserver() {
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
		return
	}
	for {
		if c, err := l.Accept(); err == nil {
			go Echo(c)
		}
	}
}
func Echo(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}
	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return
	}
}

func testcgo() {
	testgo.Seed(1)
	fmt.Println(testgo.Random())

	C.puts(C.CString("Hello, 世界\n"))
	fmt.Println("hi")
}

///////////////
func main() {
	// testdefine()
	// testfor()
	// testif()
	// teststruct()
	// testmap1()
	// testmap2()
	// testslice()
	// testslice2()
	// testfunc()
	// testclosure()
	// testforrange()
	// testswitch()
	// exercise_sqrt()
	// exercise_wc()
	// exercise_fibonacci()
	// testmethod()
	// testmethod2()
	// testref()
	// testinterface()
	// testinterface2()
	// testerror()
	// testhttpserv()
	// testimage()
	// exercise_error()
	// exercise_ioreader()
	// testgoroutine()
	// testchannel()
	// testchannel2()
	// testchannel3()
	// testchannel4()
	// testchannel5()
	// exercise_checktree()
	// exercise_craw()
	// testredis()
	// teststring()
	// testtype()
	// testfile()
	// testgoto()
	// fmt.Println(testdefer())
	// testvarargs(1, 2, 3)
	// testfuncvalue()
	// testmap()
	// testpanic()
	// exercise_functions()
	// teststack()
	// testvararg2()
	// testpackage()
	// testtest()
	// commonpkg()
	// testpointer()
	// testcustomtype()
	// testmap3()
	// testtype2()
	// testsort()
	// testinterface_recurve()
	// testintrospection()
	// testgoroutine2()
	// testio()
	// testio2()
	// testio3()
	// testexec()
	// testnet()
	// testhttp()
	// psgrp()
	// wc()
	// uniq()
	// echoserver()
	// testcgo()
}
