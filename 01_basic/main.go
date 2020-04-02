package main

import "fmt"

func main() {
	var price int
    var msg string = "Hello world"
    msg2 := "Hi"
    const msg3 = "OK"
    const (
      a = 1 + 3
      b
      c
    )

	fmt.Print("Hello")
	fmt.Print("World\n")
	fmt.Print("A", 100, true, 1.5, "\n")

    fmt.Print("Please Input Price...")
	fmt.Scan(&price)
    fmt.Printf("Price: %d\n", price)

    println(msg)
    println(msg2)
    println(msg3)

    fmt.Print(a, b, c)
}
