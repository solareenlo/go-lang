package main

func main() {
  var sum int

  sum = 5 + 6 + 3
  avg := float32(sum / 3)

  println(avg)
  if avg > 4.5 {
    println("good")
  }
}
