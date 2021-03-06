package main

import "fmt"

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan struct{})
	go func() {
		fmt.Println("sub() is finished")
		// 終了の通知
		done <- struct{}{}
	}()

	//終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
