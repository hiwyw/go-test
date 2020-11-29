package main

import (
	"fmt"
	"log"
	"time"
)

const (
	testZone       = "test.com"
	testDomainIP   = "1.1.1.1"
	testServerRoom = "300100000001"
)

type Worker struct {
	statMap map[string]int64
}

func NewWorker() *Worker {
	w := &Worker{
		statMap: map[string]int64{},
	}
	return w
}

func (w *Worker) Start(inch <-chan string) {
	for {
		i, ok := <-inch
		if ok == false {
			log.Printf("input channel closed, will return\n")
			return
		}
		w.add(i)
	}
}

func (w *Worker) add(key string) {
	v, ok := w.statMap[key]
	if ok == false {
		w.statMap[key] = 1
		return
	}
	v++
	w.statMap[key] = v
}

func main() {
	now := time.Now()
	ch := make(chan string, 10)
	w1 := NewWorker()
	go w1.Start(ch)
	w2 := NewWorker()
	go w2.Start(ch)
	w3 := NewWorker()
	go w3.Start(ch)

	for i := 0; i < 10; i++ {
		for j := 0; j < 5000000; j++ {
			ch <- genDomainKey(j)
		}
		log.Printf("%d ==>%s", i, time.Since(now))
	}
	time.Sleep(30 * time.Second)
}

func genDomainKey(index int) string {
	return fmt.Sprintf("%d.%s|%s|%s", index, testZone, testServerRoom, testDomainIP)
}
