package main

import "fmt"
import "time"
import "math/rand"

func main() {
  t := time.Now().UnixNano()
  rand.Seed(t)
  s := rand.Intn(6)
  fmt.Println(s+1)

  if s+1 == 6 {
    fmt.Println("大吉")
  } else if s+1 == 5 || s+1 == 4 {
    fmt.Println("中吉")
  } else if s+1 == 3 || s+1 == 2 {
    fmt.Println("吉")
  } else  {
    fmt.Println("凶")
  }

}
