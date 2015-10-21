package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"time"
)

func main() {
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982))
	fmt.Printf("This was touched %s\n",
		humanize.Time(time.Now().Add(time.Hour)))
	fmt.Printf("You're my %s best friend.\n", humanize.Ordinal(193))
	fmt.Printf("You owe $%s.\n", humanize.Comma(6582491))

	fmt.Printf("%f\n", 2.24)                // 2.240000
	fmt.Printf("%s\n", humanize.Ftoa(2.24)) // 2.24
	fmt.Printf("%f\n", 2.0)                 // 2.000000
	fmt.Printf("%s\n", humanize.Ftoa(2.0))  // 2

	fmt.Printf("%s\n", humanize.SI(0.00000000223, "M")) // 2.23nM
}
