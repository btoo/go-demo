package main

import (
	"time"
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func say(s string){
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	waitGroup.Done()
}

func main(){
	
	// go say("Hey")
	// say("There")

	// // the following two lines wont finish without the third to keep the process running
	// go say("Hey")
	// go say("There")
	// time.Sleep(time.Second)

	waitGroup.Add(1)
	go say("Hey")
	waitGroup.Add(1)
	go say("There")
	waitGroup.Wait()

	waitGroup.Add(1)
	go say("this message will wait for the first go routines to finish first")
	waitGroup.Wait()

}