"""
Simple UNIX dict based hangman game.
"""

import random

def main():
    # get a random word
    word = random.choice(open("/usr/share/dict/words").read().split("\n"))
    word = word.split("'")[0] # sanitize if necessary
    word = word.lower()

    wordlist = list(word)
    workcopy = list(word)
    for i in xrange(len(workcopy)):
        if workcopy[i] not in ['a', 'e', 'i', 'o', 'u']:
            workcopy[i] = '_'

    trials = 3
    for i in xrange(trials):
        inp = ""
        while len(inp) != 1 and inp.isalpha:
            inp = raw_input("{0}\n>> ".format("".join(workcopy)))

        for j in xrange(len(wordlist)):
            if wordlist[j] == inp:
                workcopy[j] = inp

    print "".join(workcopy)

    if '_' in workcopy:
        print "The word was: ", word
        print "Hard luck. Try again."
    else:
        print "You Win."

if __name__ == "__main__":
    main()
