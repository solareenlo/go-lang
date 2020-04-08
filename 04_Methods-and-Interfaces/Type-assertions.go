package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	s = i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
