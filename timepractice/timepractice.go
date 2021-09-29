package timepractice

import (
	"fmt"
	"time"
)

func TestTime() {

	//通过时间 转化成时间
	// x := 48
	initHeight := 43
	stardDay, err := time.Parse(time.RFC3339, "2020-08-25T06:21:30+08:00")
	if err != nil {
		fmt.Printf("stardDay err is %v\n", err)
	}
	// y := x / 2
	CurHeight := 1144796 - initHeight
	stopDay := stardDay.Add(time.Duration(CurHeight) * 30 * time.Second)
	endDay, _ := time.Parse("2006-01-02", "2021-10-01")

	println("stardDay is ", stardDay.Day())
	println("endDay is ", endDay.Day())
	println("stopDay is ", stopDay.String())

	return

}
