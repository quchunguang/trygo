package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later := now.Add(time.Duration(3) * time.Minute)
	spt1 := time.Date(2015, 9, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(later.Format(time.RFC3339))
	fmt.Println(spt1.Format(time.Kitchen))
	const layout = "Jan 2, 2006 at 3:04pm"
	fmt.Println(spt1.Format(layout))

	const shortForm = "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(shortForm, "2015-09-29", loc)
	fmt.Println(t)
}
