package main

import "fmt"

func main() {
  var a int = 7
  fmt.Println("Running a forloop", a, "times")

  for ; a >= 0; a-- {
	fmt.Println("Hello there #", a)
  }
}
