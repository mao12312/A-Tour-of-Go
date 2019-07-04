package p16

import (
	 "fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64,128 }

func Main() {
	/*
Range: range returns two variables per iteration.
The first variable is the index
The second is a copy of the elements of the location of the index.
	*/
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
