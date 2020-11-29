package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-09-08 10:00:01"
	out := time.Now().Local().Format(layout)
	fmt.Println(out)
}
