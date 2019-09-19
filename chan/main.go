package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	userCount := 10
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(ch, i)
	}

	//这里如果不适用wait，会导致最后一个协程还没运行完就被主协程强制结束了
	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done()
	ch <- true
	fmt.Printf("go func: %d\n", i)
	time.Sleep(time.Second * 2)
	<-ch
}
