package p01

import (
	"fmt"
)

func Main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)
}