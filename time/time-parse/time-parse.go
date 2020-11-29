package main

import (
	"fmt"

	"github.com/araddon/dateparse"
)

func main() {
	timeStr := "01-Sep-2019 11:11:10.008"
	t, err := dateparse.ParseLocal(timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
