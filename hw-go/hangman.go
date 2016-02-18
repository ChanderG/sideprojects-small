package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "math/rand"
  "time"
)

func removeApostophe(s string) string {
  return strings.Split(s, "'")[0]
}

func blankLetter(r rune) rune {
  vowels := [5]rune{'a', 'e', 'i', 'o', 'u'}
  isVowel := false
  for i:=0; i< len(vowels);i++ {
	if r == vowels[i] {
	  isVowel = true
	}
  }
  if isVowel == false {
    return '_'
  } else {
	return r
  }
}

func blankifyWord(s string) string {
  blankedWord := strings.Map(blankLetter, s)
  return blankedWord
}

func main(){
  fmt.Println("Welcome to hangman game.")

  fmt.Println("Choosing random word...")
  word := chooseWord()
  fmt.Println("word is", word)

  word = blankifyWord(word)
  fmt.Println(word)
}

func chooseWord() string {
  var filename string = "words"

  // read entire file contents
  fmt.Println("Reading file", filename)
  dat, err := ioutil.ReadFile(filename)
  if err!= nil {
	panic(err)
  }

  // convert data to string
  lines := string(dat)
  //fmt.Println(lines)

  //s := make([]string, 0)
  s := strings.Split(lines, "\n")

  // convert to lower case words
  for i:=0; i<len(s);i++ {
	s[i] = strings.ToLower(s[i])
  }

  // remove ' stuff
  for i:=0; i<len(s); i++ {
	s[i] = removeApostophe(s[i])
  }

  // init seed for actual random performance
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  // choose random value
  var word string = s[r1.Intn(len(s))]

  return word
}
