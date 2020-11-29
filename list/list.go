package main

import (
	"container/list"
	"encoding/json"
	"fmt"
)

var persions *list.List = list.New()

type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	add("zhangsan", 18)
	add("lisi", 20)
	add("wangwu", 40)
	printAll()
	remove("zhangsan")
	printAll()
	remove("zhaosi")
	printAll()
}

func printAll() {
	fmt.Printf("all persion:\n")
	for e := persions.Front(); e != nil; e = e.Next() {
		out, _ := json.Marshal(e.Value.(*Persion))
		fmt.Println(string(out))
	}
}

func add(name string, age int) {
	persions.PushBack(&Persion{name, age})
	fmt.Printf("add (%s %d) succeed\n", name, age)
}

func remove(name string) {
	for e := persions.Front(); e != nil; e = e.Next() {
		if e.Value.(*Persion).Name == name {
			persions.Remove(e)
			fmt.Printf("remove %s succeed\n", name)
			return
		}
	}
	fmt.Printf("no found %s\n", name)
}
