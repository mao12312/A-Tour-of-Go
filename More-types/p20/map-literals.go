package p20

import "fmt"

type Vertex  struct {
	Lat, Long float64
}

var m  = map[string]Vertex{
	"Bell Labs":Vertex {
		40.34134, -57.245,
	},
	"Google":Vertex{
		37.3763, -46.5345,
	},
}

func Main() {
	fmt.Println(m)
}