package main

import (
	"time"
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func cleanup(){
	defer waitGroup.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r) // r is the message that is returned by the recover fn i.e. passed into the panic fn
	}
}

func say(s string){
	defer cleanup() // if you panic, run the cleanup fn
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2") // panic with this message
		}
	}
	// waitGroup.Done()
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