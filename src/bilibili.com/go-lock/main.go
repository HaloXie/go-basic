package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock() // 加锁
	for i := 0; i < 10; i++ {
		x = x + 1
	}
	lock.Unlock() // 解锁
	wg.Done()
}

func p() {
	for i := 0; i < 10; i++ {
		fmt.Println(" p() => ", i)
		fmt.Println(" p() x => ", x)
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go p()
	wg.Wait()
	fmt.Println(x)
}
