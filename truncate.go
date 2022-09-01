package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Enter a floating point number")
	fmt.Scan(&x)

	fmt.Println(int(x))
}
