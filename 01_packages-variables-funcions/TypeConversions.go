package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)

  a, b := 5, 12
  c := math.Sqrt(float64(a*a + b*b))
  d := uint(c)
  fmt.Println(a, b, d)
}
