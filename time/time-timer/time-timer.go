package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	fmt.Println(<-timer.C)
}
