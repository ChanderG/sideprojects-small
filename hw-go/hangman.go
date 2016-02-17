package main

import "fmt"
import "os"

func main(){
  fmt.Println("Welcome to hangman game.")
  var filename string = "sample"

  fd, err := os.Open(filename)
  if err!= nil {
	panic(err)
  }

  fmt.Println("Opened file", filename)

  err = fd.Close()
  if err != nil {
	panic(err)
  }
}
