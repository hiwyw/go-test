package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("www.baidu.com", "1.1.1.1")
	m.Store("www.qq.com", "2.2.2.2")
	v, ok := m.Load("www.baidu.com")
	if ok == true {
		fmt.Printf("the value of key %s is %s\n", "www.baidu.com", v)
	}
	fmt.Printf("key %s doesn't exist\n", "www.baidu.com")
}
