package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v, s, t float64
	fmt.Print("Enter value for acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Enter value for initial velocity: ")
	fmt.Scanf("%f", &v)
	fmt.Print("Enter value for initial displacement: ")
	fmt.Scanf("%f", &s)

	displacementFn := GenDisplaceFn(a, v, s)
	fmt.Print("Enter value for time: ")
	fmt.Scanf("%f", &t)

	fmt.Println("Displacement = ", displacementFn(t))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v * t) + s
	}
}
