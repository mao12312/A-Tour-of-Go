package p09

import (
	"fmt"
)

func Main() {
	q := []int{2,3,4,6,8,22}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s :=[]struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{6, true},
		{8, false},
		{22, true},
	}
	fmt.Println(s)
}