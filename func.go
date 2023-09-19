package main

import (
	"fmt"
)

func GenDisplaceFn(a, v, s float32) func(float32) float32 {

	return func(t float32) float32 {
		return ((0.5 * a * (t * t)) + (v * t) + s)
	}
}

func main() {

	values := [4]float32{}

	fmt.Printf("Input Acceleration 'a': ")
	fmt.Scan(&values[0])

	fmt.Printf("Input Initial Velocity 'v\u2218': ")
	fmt.Scan(&values[1])

	fmt.Printf("Input Initial Displacement 's\u2218': ")
	fmt.Scan(&values[2])

	fmt.Printf("Input Time in Seconds: ")
	fmt.Scan(&values[3])

	fn := GenDisplaceFn(values[0], values[1], values[2])

	fmt.Println("\na:", values[0], ", v\u2218:", values[1], ", s\u2218:", values[2])
	fmt.Println("Displacement after", values[3], "seconds:", fn(values[3]))

}
