package main

import (
	"fmt"
	"time"
)

func main() {
	begin := time.Now()
	m := make(map[string]int64)
	for i := 0; i < 5000000; i++ {
		k := fmt.Sprintf("%d.www.1.2.3.4.5.6.7.test.com", i)
		m[k] = 1
	}
	fmt.Println(time.Since(begin))
}
