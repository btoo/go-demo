package loop

import "fmt"

func loop() {

	// for loops usually dont look like this
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// // while loop, but in practice the loops wont really look like this either
	// i := 0
	// for i < 10 {
	// // for { // infinite loop
	// 	fmt.Println(i)
	// 	// i++
	// 	i += 5
	// }

	// but this really is just the same thing as the first for loop
	// x := 5
	// for {
	// 	fmt.Println("Do stuff", x)
	// 	x += 3
	// 	if x > 25 {
	// 		break
	// 	}
	// }

	// in practice, this is probably what a loop would look like
	// for _, datum := range data {
	// 	fmt.Println(datum)
	// }

}