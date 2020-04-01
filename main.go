package main

import "fmt"

func main() {
	var price int

	fmt.Print("Hello")
	fmt.Print("World\n")
	fmt.Print("A", 100, true, 1.5, "\n")

	fmt.Scan(&price)
	fmt.Printf("%d\n", price)
}
