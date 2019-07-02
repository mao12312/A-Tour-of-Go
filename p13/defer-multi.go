package p13

import (
	"fmt"
)

func Main() {
	fmt.Println("counting")
	defer fmt.Println("done")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
