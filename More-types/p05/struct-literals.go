package p05

import "fmt"

type Vertex struct {
	X, Y int
}
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)
	 
func Main() {
	fmt.Println(v1, v2, v3, p)
}
