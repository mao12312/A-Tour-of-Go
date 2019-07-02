package p03

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func Main() {
	v := Vertex{1, 2}
	v.X = 4           // access useing dot(.)
	fmt.Println(v.X)
}