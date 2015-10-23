package main

import (
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

const (
	dx       = 100          // Width of image
	dy       = 40           // Height of image
	fontFile = "IMPACT.TTF" // Font file
	fontSize = 20           // Font size
	fontDPI  = 72           // DPI of image
)

func main() {
	// Output image
	imgfile, _ := os.Create("out.png")
	defer imgfile.Close()

	// Create a RGBA bitmap
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	// Draw background
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// Set RGBA of each point
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}

	// Get font from file
	fontBytes, err := ioutil.ReadFile(`C:\Windows\Fonts\` + fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Create context for draw
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)

	// Position to draw
	pt := freetype.Pt(20, 20+int(c.PointToFixed(fontSize)>>8))

	// Draw text with font to bitmap
	_, err = c.DrawString("ABCDE", pt)
	if err != nil {
		log.Println(err)
		return
	}

	// Save as PNG image
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
