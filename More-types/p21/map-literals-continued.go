package p21

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs":{40.351, -74.3474},
	"Google":{262.334, -45.436},
}

func  Main() {
	fmt.Println(m)
}