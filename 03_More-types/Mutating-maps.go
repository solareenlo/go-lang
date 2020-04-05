package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["Answer"] = 42
	fmt.Println(m, m["Answer"])

	m["Answer"] = 48
	fmt.Println(m, m["Answer"])

	m["Ans"] = 60
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m)

	v, ok := m["Answer"]
	fmt.Println(v, ok)

	i, j := m["Ans"]
	fmt.Println(i, j)
}
