package main

import (
	"fmt"
	"gopkg.in/ukautz/clif.v1"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return nil
	}
	if info.IsDir() {
		fmt.Println("Directory: ", path)
	} else {
		fmt.Println("File: ", path)
	}
	return nil
}

func cmdWalk(in clif.Input) {
	if !in.Confirm("Let's do it?") {
		return
	}
	root := "./"
	filepath.Walk(root, walkFn)
}

func cmdEcho(c *clif.Command) {
	s := c.Argument("text").String()
	if c.Option("upper").Bool() {
		s = strings.ToUpper(s)
	}
	for i := 0; i < c.Option("num").Int(); i++ {
		fmt.Println(s)
	}
}

var (
	headers = []string{"Name", "Age", "Force"}
	rows    = [][]string{
		{
			"<important>Yoda<reset>",
			"Very, very old",
			"Like the uber guy",
		},
		{
			"<important>Luke Skywalker<reset>",
			"Not that old",
			"A bit, but not that much",
		},
		{
			"<important>Anakin Skywalker<reset>",
			"Old dude",
			"He is Lukes father! Was kind of stronger in 1-3, but still failed to" +
				" kill Jar Jar Binks. Not even tried, though. What's with that?",
		},
	}
)

func cmdTable(out clif.Output) {
	table := out.Table(headers)
	table.AddRows(rows)
	fmt.Println(table.Render())
}

func cmdProgress(out clif.Output) {
	pbs := out.ProgressBars()
	pbs.Style(clif.ProgressBarStyleAscii)
	pbs.Start()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		pb, _ := pbs.Init(fmt.Sprintf("bar-%d", i+1), 200)
		go func(b clif.ProgressBar, ii int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				b.Increment()
				<-time.After(time.Millisecond * time.Duration(100*ii))
			}
		}(pb, i)
	}
	wg.Wait()
	<-pbs.Finish()
}

func main() {
	cli := clif.New("Phototool", "0.0.1", "A tool box for photographer.")
	cWalk := clif.NewCommand("walk", "Walk through folder.", cmdWalk)
	cTable := clif.NewCommand("table", "Print out a table.", cmdTable)
	cProgress := clif.NewCommand("progress", "Progress bar.", cmdProgress)

	cEcho := clif.NewCommand("echo", "Say something n times.", cmdEcho)
	cEcho.NewArgument("text", "What to be print out.", "Hello", false, false)
	cEcho.NewOption("num", "n", "Num of times.", "1", false, false)
	cEcho.NewFlag("upper", "u", "Print out with upper case.", false)

	cli.Add(cWalk).Add(cTable).Add(cProgress).Add(cEcho)
	cli.Run()
}
