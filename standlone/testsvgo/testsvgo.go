package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func main() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, 曲春光", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}
