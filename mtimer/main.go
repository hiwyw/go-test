package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	for i := 1; i < 40000; i++ {
		go func(id int) {
			c := 0
			randInt := rand.Intn(10)
			if randInt == 0 {
				randInt = 5
			}
			if randInt < 0 {
				randInt = -randInt
			}
			t := time.NewTicker(time.Second * time.Duration(randInt))
			for {
				<-t.C
				c++
				fmt.Printf("thead %v count %v\n", id, c)
			}
		}(i)
	}

	fmt.Printf("thead umber: %v\n", runtime.NumGoroutine())

	for {
		c := 0
		t := time.NewTicker(time.Second * time.Duration(3))
		for {
			<-t.C
			c++
			fmt.Printf("main thead count %v\n", c)
			fmt.Printf("thead umber: %v\n", runtime.NumGoroutine())
		}
	}
}
