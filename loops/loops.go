package main

import (
	"fmt"
	"math"
)

//Sqrt Run below
// first implement
/*
func Sqrt(x float64) float64 {
	z := 1.
	for i := 0; i < 11; i++  {
		z -= (z*z - x) / (2*x)
	}
	return z

}
*/

// second implement
// Loop end when the difference between z and n falls below the criteria described in the if statement
func Sqrt(x float64) float64 {
	z := 1.
	n := 0.
	for {
		z, n = z-(z*z-x)/(2*z), z
		if math.Abs(n-z) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	x := 2.
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
