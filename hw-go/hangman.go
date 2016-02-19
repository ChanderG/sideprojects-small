package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "math/rand"
  "time"
  "bufio"
  "os"
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

/* 
 * Prompt for an input char, run on word and return new word.
 */
func guessLetter(actual string, curr string) string {
  // scan for an input
  fmt.Print("> ")
  // read one line input
  b := bufio.NewReader(os.Stdin)
  input, err := b.ReadString('\n')
  if err != nil {
	panic(err)
  }
  char_guess := []rune(input)[0]

  curr_list := []rune(curr)
  actual_list := []rune(actual)

  for i:=0; i< len(curr_list);i++ {
	if actual_list[i] == char_guess {
	  // illuminate
	  curr_list[i] = char_guess
	}
  }
  return string(curr_list)
}

func playGame(actual string, word string, n int) {
  curr := word
  for i:=0; i< n; i++ {
	fmt.Println(curr)
	curr = guessLetter(actual, curr)
  }
  if curr == actual{
	fmt.Println("Good job. You won.")
  } else {
	fmt.Println("Word was", actual)
	fmt.Println("Hard Luck. Try again.")
  }
}

func main(){
  fmt.Println("Welcome to hangman game.")

  fmt.Println("Choosing random word...")
  word := chooseWord()
  //fmt.Println("word is", word)
  /*
  Partially correct way to find length of word.
  Still need to figure out a way to loop over chars
  fmt.Println("len of word: ", len([]rune(word)))
  */

  //fmt.Println(word)
  newword := blankifyWord(word)

  playGame(word, newword, 3)
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
