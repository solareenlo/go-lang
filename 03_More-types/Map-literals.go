package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.0, -17,
	},
	"Google": Vertex{
		2.01, 222.2,
	},
}

func main() {
	fmt.Println(m)
}
