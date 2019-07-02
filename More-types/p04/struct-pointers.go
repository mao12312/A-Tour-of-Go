package p04

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func Main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	p.Y = 1e6
	fmt.Println(v)
}