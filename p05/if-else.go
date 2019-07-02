package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v<lim {
		return v
	} else {
		//%e(scientific notation, e.g. -1.234456e+78) for large exponents, %f otherwise. 
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(20, 2, 3),
		pow(20, 10, 10),
	)
}