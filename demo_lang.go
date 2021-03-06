package trygo

/*
#include <stdio.h>

*/
//import "C" // this MUST be single sentence with magic omments above !!!
import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no"

// Pi math
const Pi = 3.14
const (
	// Big int
	Big = 1 << 60
	// Small int
	Small = Big >> 59
)

// Vertex struct
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
	}
	// fmt.Printf("%g >= %g\n", v, lim) // 不能在这里使用 v
	return lim
}

// DemoDefine func
func DemoDefine() {
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
	fmt.Println( //Big,
		Small, needInt(Small), needFloat(Small), needFloat(Big))
}

// DemoFor func
func DemoFor() {
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
		fmt.Printf("character '%c' starts at byte position %d\n",
			char, pos)
	}
}

// DemoIf func
func DemoIf() {
	fmt.Println(sqrt(-82))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// DemoStruct func
func DemoStruct() {
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

// DemoMap1 func
func DemoMap1() {
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

// DemoMap2 func
func DemoMap2() {
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

// DemoSlice func
func DemoSlice() {
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

// DemoSlice2 func
func DemoSlice2() {
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

// DemoFunc func
func DemoFunc() {
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

// DemoClosure func
func DemoClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// DemoForrange func
func DemoForrange() {
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

// DemoSwitch func
func DemoSwitch() {
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

// Sqrt func
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

// ExerciseSqrt func
func ExerciseSqrt() {
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

// ExerciseWc func
func ExerciseWc() {
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

// ExerciseFibonacci func
func ExerciseFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// Abs func
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

// DemoMethod func
func DemoMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// MyFloat type
type MyFloat float64

// Abs func
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// DemoMethod2 func
func DemoMethod2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// NotChange func
func (v Vertex) NotChange() {
	v.Long = 0
}

// DemoRef func
func DemoRef() {
	v := Vertex{3, 4}
	v.NotChange()
	fmt.Println(v)

	// map is pointer type
	var m = map[int]int{5: 1, 6: 2}
	var n map[int]int // just a ref
	m[1] = 2
	fmt.Println(n == nil)
	n = m
	m[2] = 4
	fmt.Println(n)

	// slice is pointer type
	var o = []int{1, 2, 3}
	var p []int // just a ref
	fmt.Println(p == nil)
	p = o
	o[0] = 10
	fmt.Println(p)

	// array is value type
	var r = [5]int{1, 2, 3, 4, 5}
	var s [5]int
	// fmt.Println(s == nil) // err: cannot convert nil to type [5]int
	s = r
	r[0] = 10
	fmt.Println(s)
}

// Abser interface
type Abser interface {
	Abs() float64
}

// DemoInterface func
func DemoInterface() {
	var ia Abser
	f := MyFloat(-math.Sqrt2)
	ia = f
	v := &Vertex{3, 4}
	ia = v

	fmt.Println(ia.Abs())
}

// Reader interface
type Reader interface {
	Read(b []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(b []byte) (n int, err error)
}

// ReadWriter interface
type ReadWriter interface {
	Reader
	Writer
}

// DemoInterface2 func
func DemoInterface2() {
	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
}

// MyError struct
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

// DemoError func
func DemoError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// ErrNegativeSqrt sorter
type ErrNegativeSqrt float64

// Error func
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Input %f: can not be negative", e)
}

// SqrtE func
func SqrtE(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

// ExerciseError func
func ExerciseError() {
	fmt.Println(SqrtE(2))
	fmt.Println(SqrtE(-2))
}

type rot13Reader struct {
	r io.Reader
}

// Read func
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

// ExerciseIoreader func
func ExerciseIoreader() {
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

// DemoGoroutine func
func DemoGoroutine() {
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

// DemoChannel func
func DemoChannel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// DemoChannel2 func
func DemoChannel2() {
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

// DemoChannel3 func
func DemoChannel3() {
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

// DemoChannel4 func
func DemoChannel4() {
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

// DemoChannel5 func
func DemoChannel5() {
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

// Tree struct
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func walkPreorderRecursive(t *Tree, c chan int) {
	c <- t.Value
	if t.Left != nil {
		walkPreorderRecursive(t.Left, c)
	}
	if t.Right != nil {
		walkPreorderRecursive(t.Right, c)
	}
}
func walkInorderRecursive(t *Tree, c chan int) {
	if t.Left != nil {
		walkInorderRecursive(t.Left, c)
	}
	c <- t.Value
	if t.Right != nil {
		walkInorderRecursive(t.Right, c)
	}
}

func walkPostorderRecursive(t *Tree, c chan int) {
	if t.Left != nil {
		walkPostorderRecursive(t.Left, c)
	}
	if t.Right != nil {
		walkPostorderRecursive(t.Right, c)
	}
	c <- t.Value
}
func walkLevelorderRecursive(t *Tree, c chan int) {
}
func walkPreorder(t *Tree, c chan int) {
	s := NewStack()
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
func walkInorder(t *Tree, c chan int) {

}
func walkPostorder(t *Tree, c chan int) {

}
func walkLevelorder(t *Tree, c chan int) {

}
func compareChan(c1, c2 chan int) {
	var ic1, ic2 int
	var hasC1, hasC2 bool
	for {
		ic1, hasC1 = <-c1
		ic2, hasC2 = <-c2

		if hasC1 == false && hasC2 == false {
			fmt.Println("EQUARE.")
			return
		}
		if hasC1 == false || hasC2 == false || ic1 != ic2 {
			fmt.Println("NOT EQUARE.")
			return
		}
	}
}

// ExerciseChecktree func
func ExerciseChecktree() {
	// create trees and init chans
	t1 := &Tree{&Tree{&Tree{nil, 1, nil}, 1, &Tree{nil, 2, nil}},
		3, &Tree{&Tree{nil, 5, nil}, 8, &Tree{nil, 13, nil}}}
	t2 := &Tree{&Tree{&Tree{&Tree{nil, 1, nil}, 1, &Tree{nil, 2, nil}}, 3,
		&Tree{nil, 5, nil}}, 8, &Tree{nil, 13, nil}}
	c1 := make(chan int)
	c2 := make(chan int)

	// begin generate chans
	go func() {
		walkInorderRecursive(t1, c1)
		close(c1)
	}()
	go func() {
		walkInorderRecursive(t2, c2)
		close(c2)
	}()

	// begin compare chans
	compareChan(c1, c2)
}

// Fetcher interface
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

// ExerciseCraw func
func ExerciseCraw() {
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

// DemoString func
func DemoString() {
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

// DemoType func
func DemoType() {
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

// DemoGoto func
func DemoGoto() {
	i := 0
Here:
	fmt.Println(i)
	if i++; i == 5 {
		return
	}
	goto Here
}
func deferit() (ret int) {
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

// DemoDefer func
func DemoDefer() {
	fmt.Println(deferit())
}

func myfunc2(arg ...interface{}) {

}
func myfunc1(arg ...int) {

}
func callvarargs(arg ...int) {
	myfunc1(arg...)
	myfunc1(arg[:2]...)
	myfunc2(arg)
	myfunc2(4, true, "abc")
}

// DemoVarargs func
func DemoVarargs() {
	callvarargs(1, 2, 3)
}

// DemoFuncvalue func
func DemoFuncvalue() {
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

// Map func
func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

// DemoMap func
func DemoMap() {
	double := func(a int) int { return 2 * a }
	l := []int{1, 2, 3, 4, 5}
	m := Map(double, l)
	fmt.Println(m)
}
func f() {
	// panic(interface{})是一个内建函数,可以中断原有的控制流程,
	// 进入一个令人恐慌的流程中。当函数F调用panic,函数F的执行被中断,
	// 并且F中的延迟函数会正常执行,然后 F 返回到调用它的地方。
	// 在调用的地方, F的行为就像调用了 panic 。这一过程继续向上,
	// 直到程序崩溃时的所有 goroutine 返回。
	// 恐慌可以直接调用panic产生。也可以由运行时错误产生,
	// 例如访问越界的数组。
	fmt.Println("Hi, panic?")
	panic("i'm panic!")
}
func throwsPanic(f func()) (b bool) {
	// recover()是一个内建的函数,可以让进入令人恐慌的流程中的goroutine
	// 恢复过来。仅在延迟函数中有效。
	defer func() {
		if x := recover(); x != nil {
			b = true
			fmt.Println(x.(string))
		}
	}()
	f()
	return
}

// DemoPanic func
func DemoPanic() {
	var ret = throwsPanic(f)
	fmt.Println(ret)
}

// ExerciseFunctions func
func ExerciseFunctions() {
	d := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum float64
	for _, v := range d {
		sum += v
	}
	fmt.Printf("Average = %v\n", sum/float64(len(d)))
}
func printthem(them ...int) {
	for _, d := range them {
		fmt.Println(d)
	}
}

// DemoVararg2 func
func DemoVararg2() {
	printthem(1, 4, 5, 7, 4)
	printthem(1, 2, 4)
}

// DemoPkg func
func DemoPkg() {
	// fmt
	// %v 默认格式的值。当打印结构时,加号( %+v )会增加字段名;
	// %#v Go 样式的值表达;
	// %T 带有类型的 Go 样式的值表达;

	// io
	// 这个包提供了原始的 I/O 操作界面。
	// 它主要的任务是对os包这样的原始的 I/O 进行封装,增加一些其他相关,
	// 使其具有抽象功能用在公共的接口上。

	// bufio
	// 这个包实现了缓冲的 I/O 。它封装于 io.Reader 和 io.Writer 对象,
	// 创建了另一个对象(Reader和Writer)在提供缓冲的同时实现了一些文本
	// I/O的功能。

	// sort
	// sort 包提供了对数组和用户定义集合的原始的排序功能。

	// strconv
	// strconv 包提供了将字符串转换成基本数据类型,
	// 或者从基本数据类型转换为字符串的功能。

	// os
	// os 包提供了与平台无关的操作系统功能接口。其设计是 Unix 形式的。

	// sync
	// sync 包提供了基本的同步原语,例如互斥锁。

	// flag
	// flag 包实现了命令行解析。

	// encoding/json
	// encoding/json 包实现了编码与解码 RFC 4627 定义的 JSON 对象。

	// text/template
	// 数据驱动的模板,用于生成文本输出,例如 HTML 。将模板关联到某个
	// 数据结构上进行解析。模板内容指向数据结构的元素(通常结构的
	// 字段或者 map 的键)控制解析并且决定某个值会被显示。
	// 模板扫描结构以便解析,而 “ 游标 ” @ 决定了当前位置在结构中的值。

	// net/http
	// net/http 实现了 HTTP 请求、响应和 URL 的解析,并且提供了可扩展的
	// HTTP服务和基本的 HTTP 客户端。

	// unsafe
	// unsafe包包含了Go程序中数据类型上所有不安全的操作。

	// reflect
	// reflect 包实现了运行时反射,允许程序通过抽象类型操作对象。
	// 通常用于处理静态类型 interface{} 的值,并且通过 Typeof 解析出
	// 其动态类型信息,通常会返回一个有接口类型 Type 的对象。

	// os/exec
	// os/exec 包执行外部命令。
}

const (
	// Enone const
	Enone = 1
	// Einval const
	Einval = 2
)

// SyncedBuffer struct
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

// DemoPointer func
func DemoPointer() {
	var p *int
	fmt.Printf("%v\n", p)

	var i = 9
	p = &i
	*p = 8
	fmt.Printf("%v %v %v\n", p, *p, i)
	*p++ // means (*p)++
	fmt.Printf("%v\n", i)

	m := new(SyncedBuffer)
	// var n SyncedBuffer
	fmt.Println(m)

	ar := [...]string{Enone: "no error", Einval: "invalid argument"}
	sl := []string{Enone: "no error", Einval: "invalid argument"}
	ma := map[int]string{Enone: "no error", Einval: "invalid argument"}
	fmt.Println(ar, sl, ma)

}

type foo int

// NameAge struct
type NameAge struct {
	name string
	age  int
}

// DemoCustomtype func
func DemoCustomtype() {
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
	var x = bar{1}
	var y = foo(x)
	fmt.Println(x, y)

}

// E interface define the empty interface as a type
type E interface{}

func mult2(f E) E {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string)
	}
	return f
}

// Map3 func
func Map3(n []E, f func(E) E) []E {
	m := make([]E, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

// DemoMap3 func
func DemoMap3() {
	m := []E{1, 2, 3, 4}
	s := []E{"a", "b", "c", "d"}
	mf := Map3(m, mult2)
	sf := Map3(s, mult2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

// I interface
type I interface {
	Get() int
	Put(int)
}

// S struct
type S struct{ i int }

// Get func
func (p *S) Get() int { return p.i }

// Put func
func (p *S) Put(v int) { p.i = v }

// R struct
type R struct{ i int }

// Get func
func (p *R) Get() int { return p.i }

// Put func
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

func g(something interface{}) int {
	return something.(I).Get()
}

// DemoType2 func
func DemoType2() {
	var s S
	var ps *S
	ps = &s
	cheacktype(ps)
	s.Put(2)
	fmt.Println(s)

	// fmt.Println(g(s))  S are NOT I
	fmt.Println(g(ps))
}

// Interface1 interface
type Interface1 interface {
	Len(x interface{})
}

// Interface2 interface
type Interface2 interface {
	Interface1
	Push(x interface{})
	Pop() interface{}
}

// DemoInterfaceRecurve func
func DemoInterfaceRecurve() {

}

// Person struct
type Person struct {
	Name string `json:"name"` // Tag
	Age  int    `json:"age"`
}

// ShowTag func
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

// DemoIntrospection func
func DemoIntrospection() {
	p1 := new(Person)
	p1.Name = "guang"
	ShowTag(p1)
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

// DemoGoroutine2 func
func DemoGoroutine2() {
	go ready("Tea", 1)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(2 * time.Second)
	gomaxprocs := runtime.GOMAXPROCS(4) // or set env GOMAXPROCS
	gomaxprocs = runtime.GOMAXPROCS(gomaxprocs)
	fmt.Println(gomaxprocs)
}

// func DemoCgo() {
// 	C.puts(C.CString("Hello, 世界\n"))
// 	fmt.Println("hi")
// }

// DemoFunc2 func
func DemoFunc2() {
	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}
	for _, f := range fns {
		fmt.Printf("%d ", f(2, 3))
	}
	fn1 := fns[rand.Intn(len(fns))]
	x, y := 12, 5
	fmt.Printf("\n%d\n", fn1(x, y))
}

///////
type binFunc func(int, int) int
type walkFn func(*int) walkFn

func walkEqual(i *int) walkFn {
	*i += rand.Intn(7) - 3
	return walkEqual
}

// DemoTypeRecursive func
func DemoTypeRecursive() {
	fn2, progress := walkEqual, 0
	for i := 0; i < 20; i++ {
		fn2 = fn2(&progress)
		fmt.Printf("%d ", progress)
	}
}

//////
func pickFunc(fns ...func()) func() {
	return fns[rand.Intn(len(fns))]
}

func produce(c chan func(), n int, fns ...func()) {
	defer close(c)
	for i := 0; i < n; i++ {
		c <- pickFunc(fns...)
	}
}

// DemoFuncchan func
func DemoFuncchan() {
	var delay = 200 * time.Millisecond
	// time is frozen on Playground, so this is always the same.
	rand.Seed(time.Now().Unix())

	x := 10
	fns := []func(){
		func() { x++ },
		func() { x-- },
		func() { x *= 2 },
		func() { x /= 2 },
		func() { x *= x },
	}

	c := make(chan func())
	go produce(c, 10, fns...)

	for fn := range c {
		fn()
		fmt.Printf("%d ", x)
		time.Sleep(delay)
	}
}

// MyString string
type MyString string

// DemoConst func
func DemoConst() {
	const hello = "Hello, 世界"
	const typedHello string = "Hello, 世界"
	const myStringHello MyString = "Hello, 世界"

	var s string
	s = typedHello
	fmt.Println(s)

	var m MyString
	// m = typedHello // Type error
	m = MyString(typedHello)
	fmt.Println(m)

	// untyped string constant can Assigns to a variable of any type compatible with strings
	m = "Hello, 世界"
	m = hello

	// an untyped constant has a default type
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)
	fmt.Printf("%T: %v\n", myStringHello, myStringHello)

	// Default type determined by syntax
	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", 'x', 'x')
	fmt.Printf("%T %v\n", 0i, 0i)

	// boolean
	type MyBool bool
	const True = true
	const TypedTrue bool = true
	var mb MyBool
	mb = true // OK
	mb = True // OK
	// mb = TypedTrue // Bad, const TypedTrue has type bool
	fmt.Println(mb)

	// float64
	type MyFloat64 float64
	const Zero = 0.0
	const TypedZero float64 = 0.0
	var mf MyFloat64
	mf = 0.0  // OK
	mf = Zero // OK
	// mf = TypedZero // Bad
	fmt.Println(mf)

	var f32 float32
	f32 = 0.0
	f32 = Zero // OK: Zero is untyped
	// f32 = TypedZero // Bad: TypedZero is float64 not float32.
	fmt.Println(f32)

	const Huge = 1e1000
	// Bad: constant 1.00000e+1000 overflows float64
	// fmt.Println(Huge)
	// can use it in expressions with other constants and use the value of those expressions if the result can be represented in the range of a float64
	fmt.Println(Huge / 1e999)

	// complex
	type MyComplex128 complex128
	const I = (0.0 + 1.0i)
	const TypedI complex128 = (0.0 + 1.0i)
	var mc MyComplex128
	mc = (0.0 + 1.0i) // OK
	mc = I            // OK
	// mc = TypedI       // Bad
	fmt.Println(mc)

	const Two = 2.0 + 0i
	val := Two
	fmt.Printf("%T: %v\n", val, val)
	var ff float64
	var gg float64 = Two
	ff = Two
	fmt.Println(ff, "and", gg)

	// Integers
	// The same example could be built for any interger tyoes:
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// uintptr
	type MyInt int
	const Three = 3
	const TypedThree int = 3
	var mi MyInt
	mi = 3     // OK
	mi = Three // OK
	// mi = TypedThree // Bad
	fmt.Println(mi)

	// rune and byte
	var shi = '世'
	var sshi = "世"
	fmt.Println(shi)         // print rune witch is unicode based 10
	fmt.Println(len(sshi))   // the lenth of utf-8 bytes is 3
	for _, v := range sshi { // one time only. v is rune (unicode)
		fmt.Println(v)
	}
	for i := 0; i < len(sshi); i++ {
		fmt.Println(sshi[i]) // 3 times. sshi[i] is byte (utf-8)
	}
	// Error: '世' has value 0x4e16, too large.
	// type Char byte
	// var c Char = '世'

	const MaxUint = ^uint(0)
	fmt.Printf("%x\n", MaxUint)

	// one!
	var f float32 = 1
	var i int = 1.000
	var u uint32 = 1e3 - 99.0*10.0 - 9
	var c float64 = '\x01'
	var p uintptr = '\u0001'
	var r complex64 = 'b' - 'a'
	var b byte = 1.0 + 3i - 3.0i
	fmt.Println(f, i, u, c, p, r, b)

	var fff = 'a' * 1.5
	fmt.Println(fff)

}

type block []byte

func (b block) GetA() int64 { return int64(b[0])<<8 | int64(b[1]) }
func (b block) GetB() int64 { return int64(b[2]) }
func (b block) GetC() int64 { return int64(b[3]) }

func DemoBinary() {
	b := make(block, 4, 4)

	f, _ := os.Open("data")
	defer f.Close()
	io.ReadFull(f, b)
	fmt.Printf("%x\n", b.GetA())
	fmt.Printf("%x\n", b.GetB())
	fmt.Printf("%x\n", b.GetC())
}
