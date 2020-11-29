package main

import (
	"fmt"
	"time"
)

func main() {
	tiker := time.NewTicker(3 * time.Second)
	for {
		fmt.Println(<-tiker.C)
	}
}
