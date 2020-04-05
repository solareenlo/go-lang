package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.1, 20.2},
	"Google":    {23.2, -11.0},
}

func main() {
	fmt.Println(m)
}
