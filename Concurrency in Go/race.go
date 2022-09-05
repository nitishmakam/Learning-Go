package main

import (
	"fmt"
	"time"
)

/* A race condition is when multiple threads are trying to access and manipulate the same variable.

In the code below, increment and decrement are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.

*/

func increment(pointer *int) {
	(*pointer)++
	fmt.Println(*pointer)
}

func decrement(pointer *int) {
	(*pointer)--
	fmt.Println(*pointer)
}
func main() {
	i := 0

	go increment(&i)
	go decrement(&i)

	i++

	time.Sleep(1 * time.Second)
	fmt.Println()
}
