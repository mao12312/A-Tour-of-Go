package p06

import "fmt"

func Main() {
	var  a [2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // output string
	fmt.Println(a)		    // output array

	params := [6]int{2, 4, 5, 6, 3, 2}
	fmt.Println(params)     // output array
}