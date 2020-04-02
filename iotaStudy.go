package main

import "fmt"

func main() {
  const (
    a = iota
    b
    c
    d
  )
  const (
    e = 1 << iota
    f
    g
    h
  )

  fmt.Println(a, b, c, d)
  fmt.Println(e, f, g, h)
}
