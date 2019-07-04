package p17

import (
	"fmt"
)

func Main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Println(pow)
	}
	// i change "_"  
	for i, value := range pow {
		fmt.Printf("%d %d \n", i, value)
	/*
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	*/
	}
}