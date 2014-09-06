package likegui

import (
	"fmt"
	"os"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// A simple event callback framework by golang.
//
// $ go build eventframework.go
// $ ./eventframework
// --> OnloadForm1(): Form1 was loaded.
//     form1 width=640, height=480 Title=This is main form of application
// --> OnoverCheckbox1(): Checkbox1 was overed.
//     Interface function Tostring()=Checkbox
//     checkbox1 at (20, 0) Checked=true
// --> OnclickButton1(): Button1 was overed.
// --> OnclickButton1(): Button1 was clicked.
//     Interface function Tostring()=Button
//     btn1 at (0, 0) Text=OK

///////////////////////////////////////////////////////////////////////////////
// User side code

// user programming entry OnloadApplication()
func OnloadApplication(c Control) {
	// using app access functions privid by framework
	var app Application = c.(Application)
	var obj Control

	// create object
	obj = Form{640, 480, "This is main form of application"}
	// binding (objid, eventtype) with target function
	app.Binding("frmMain", LOAD, OnloadForm1)
	// add objects to gui
	app.Add("frmMain", obj)

	obj = Button{0, 0, "OK"}
	app.Binding("btnOK", CLICK, OnclickButton1)
	app.Binding("btnOK", OVER, OnoverButton1) // if commet binding, quiet
	app.Add("btnOK", obj)

	obj = Checkbox{20, 0, true}
	app.Binding("ckbUpcase", OVER, OnoverCheckbox1)
	app.Add("ckbUpcase", obj)
}

// user customized event callback functions
func OnclickButton1(c Control) {
	// upcast Control to concrete type
	var btn Button = c.(Button)

	fmt.Println("--> OnclickButton1(): Button1 was clicked.")
	// call method implemented the interface Control
	fmt.Printf("    Interface function Tostring()=%s\n", btn.Tostring())
	// access data/call method of concrete control
	fmt.Printf("    btn1 at (%d, %d) Text=%s\n", btn.X, btn.Y, btn.Text)

	os.Exit(0)
}

func OnoverButton1(c Control) {
	fmt.Println("--> OnclickButton1(): Button1 was overed.")
}

func OnloadForm1(c Control) {
	var frm Form = c.(Form)
	fmt.Println("--> OnloadForm1(): Form1 was loaded.")
	fmt.Printf("    form1 width=%d, height=%d Title=%s\n", frm.Width, frm.Height, frm.Text)
}

func OnoverCheckbox1(c Control) {
	var ckb Checkbox = c.(Checkbox)

	fmt.Println("--> OnoverCheckbox1(): Checkbox1 was overed.")
	fmt.Printf("    Interface function Tostring()=%s\n", ckb.Tostring())
	fmt.Printf("    checkbox1 at (%d, %d) Checked=%v\n", ckb.X, ckb.Y, ckb.Checked)
}

///////////////////////////////////////////////////////////////////////////////
// Framework side code

// Control is the base interface of all control types
type Control interface {
	Tostring() string
}

// control Form
type Form struct {
	Width  int
	Height int
	Text   string
}

func (f Form) Tostring() string {
	return "Form"
}

// control Button
type Button struct {
	X    int
	Y    int
	Text string
}

func (b Button) Tostring() string {
	return "Button"
}

// control Checkbox
type Checkbox struct {
	X       int
	Y       int
	Checked bool
}

func (c Checkbox) Tostring() string {
	return "Checkbox"
}

// control Application, Yeah!!! Application is control too.
const (
	CLICK = iota
	DCLICK
	EXIT
	LOAD
	OVER
)

type eventkey struct {
	objid   string
	eventid int
}
type Application struct {
	Text     string
	Objects  map[string]Control
	Bindings map[eventkey]func(Control)
}

func (a Application) Tostring() string {
	return "Application"
}

// add target control to application, LOAD event generated either
func (a Application) Add(objid string, c Control) {
	a.Objects[objid] = c

	switch c.(type) {
	case Application:
		// create a application
	case Form:
		// create a form
	case Button:
		// create a button
	case Checkbox:
		// create a checkbox
	}
	// common events LOAD
	a.Event(objid, LOAD)
}

// binding (object, event) with process function
func (a Application) Binding(objid string, eventid int, op func(Control)) {
	key := eventkey{objid, eventid}
	a.Bindings[key] = op
}

// process event will cause callback function executed
func (a Application) Event(objid string, eventid int) {
	key := eventkey{objid, eventid}
	// do nothing if no function binding on this (object, event)
	if a.Bindings[key] != nil {
		a.Bindings[key](a.Objects[objid])
	}
}

// emulate_events
type event struct {
	objid   string
	eventid int
}

var events = make(chan event)

// emulate event generator will raise one event per-second
func generate_events() {

	time.Sleep(1000000000) // 1 second
	events <- event{"ckbUpcase", OVER}

	time.Sleep(1000000000)
	events <- event{"btnOK", OVER}

	time.Sleep(1000000000)
	events <- event{"btnOK", CLICK}
}

// framework's main event loop
func main() {
	// create application and add to dictionary
	// this will cause binding LOAD event raised
	app := Application{"Application", make(map[string]Control), make(map[eventkey]func(Control))}
	app.Binding("Application", LOAD, OnloadApplication)
	app.Add("Application", app)

	// generate events from go routine
	go generate_events()

	// event process loop, receive events from goroutine
	var e event
	for {
		e = <-events
		app.Event(e.objid, e.eventid)
	}
}
