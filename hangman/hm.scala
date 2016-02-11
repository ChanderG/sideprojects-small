object HangmanApp {
  def main(args: Array[String]): Unit = {
    println("Hangman game")
    val word = chooseWord
    /*
    println(word)
    println(blankifyWord(word))
    println(guessLetter(word, blankifyWord(word), 's'))
    */
   val fin = (0 until 3).foldLeft(blankifyWord(word)){ case (a, b) => round(word, a)}
   println(fin)
   if (fin == word){
     println("Good job. You win.")
   }
   else{
     println("Word was %s".format(word))
     println("Hard luck. Try again.")
   }
  }

  def chooseWord(): String = {
    val dict_file = scala.io.Source.fromFile("words", "utf-8")
    val words_raw = try dict_file.getLines.toList finally dict_file.close()
    val words = words_raw.map(_.split("'").head)
    words.apply(scala.util.Random.nextInt(words.length)).toLowerCase
  }

  def blankifyWord(inp: String): String = {
    inp.toList.map(x => if (List('a', 'e', 'i', 'o', 'u').contains(x)) x else '_' ).mkString("")
  }

  def guessLetter(full: String, part: String, guess: Char): String = {
    val pairs = full.toList.zip(part.toList)
    pairs.map{case (a, b) => if (a == guess) a else b}.mkString("")
  }

  def round(full: String, part: String): String = {
    // simulate one round of guess
    // returns the partial completed work string
    println(part)
    print(">> ")
    val inp = readChar()
    guessLetter(full, part, inp)
  }

}
