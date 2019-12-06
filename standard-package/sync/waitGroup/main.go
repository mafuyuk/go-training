package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitgroup()
	fmt.Println("====================")
	waitgroupfor()
}

func waitgroup() {
	var wg sync.WaitGroup

	wg.Add(1) //ゴルーチン実行前に呼ぶ
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete")
}

func waitgroupfor() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
