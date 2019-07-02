package p08

import "fmt"

func Main(){
	names := [4]string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}