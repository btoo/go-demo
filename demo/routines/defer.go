package main

import "fmt"

func foo(){
	defer fmt.Println("Done!") // only executes once the rest of the function foo finishes
	defer fmt.Println("Are we done?") // defer statements are last in first out
	fmt.Println("doing some stuff, who knows what")
}

func main(){
	foo()
}