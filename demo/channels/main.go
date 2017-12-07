package main

import (
	"fmt"
	"time"
	"sync"
)

var waitGroup sync.WaitGroup

func foo(c chan int, someValue int){
	defer waitGroup.Done()
	c <- someValue * 5 // send someValue to the channel
}

func main() {
	
	fooVal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go foo(fooVal, i)
	}

	waitGroup.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

	time.Sleep(time.Second * 2)

}