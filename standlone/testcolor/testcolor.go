package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
)

func main() {
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Yellow("yellow"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Magenta("magenta"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.White("white"))
	fmt.Println(color.Grey("grey"))

	fmt.Println(color.BlackBg("black background", color.Wht))
	fmt.Println(color.RedBg("red background"))
	fmt.Println(color.GreenBg("green background"))
	fmt.Println(color.YellowBg("yellow background"))
	fmt.Println(color.BlueBg("blue background"))
	fmt.Println(color.MagentaBg("magenta background"))
	fmt.Println(color.CyanBg("cyan background"))
	fmt.Println(color.WhiteBg("white background"))

	fmt.Println(color.Bold("bold"))
	fmt.Println(color.Dim("dim"))
	fmt.Println(color.Italic("italic"))
	fmt.Println(color.Underline("underline"))
	fmt.Println(color.Inverse("inverse"))
	fmt.Println(color.Hidden("hidden"))
	fmt.Println(color.Strikeout("strikeout"))

	fmt.Println(color.Green("bold green with white background", color.B, color.WhtBg))
	fmt.Println(color.Red("underline red", color.U))
	fmt.Println(color.Yellow("dim yellow", color.D))
	fmt.Println(color.Cyan("inverse cyan", color.In))
	fmt.Println(color.Blue("bold underline dim blue", color.B, color.U, color.D))

	color.Disable()
	color.Enable()

	c := color.New()
	c.Green("green")
}
