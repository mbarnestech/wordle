package wordle

import (
	"errors"

	words "github.com/mbarnestech/wordle/words"
)

const (
	maxGuesses = 6
	wordSize   = 5
)

// letterStatus can be none, correct, present, or absent
type letterStatus int

const (
	// none = no status, not guessed yet
	none letterStatus = iota // 0
	// absent = not in the word
	absent // 1
	// present = in the word, but not in the correct position
	present // 2
	// correct = in the correct position
	correct // 3
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [wordSize]byte
	// guesses holds the guesses that the user has made
	guesses [maxGuesses]guess
	// currGuess is the index of the available slot in guesses
	currGuess int
}

// guess is an attempt to guess the word
type guess [wordSize]letter

type letter struct {
	// char is the letter that this struct represents
	char byte
	// status is the state of the letter (absent, present, correct)
	status letterStatus
}

// newWordleState builds a new wordleState from a string.
// Pass in the word you want the user to guess.
func newWordleState(word string) wordleState {

	w := wordleState{
		word:      [wordSize]byte([]byte(word)),
		guesses:   [maxGuesses]guess{},
		currGuess: 0,
	}

	// copy(w.word[:], word) // takes string word, copies it into array of bytes

	return w
}

// newLetter should take in a character (as a byte) and return
// a new letter with that character and a none status.

func newLetter(char byte) letter {
	return letter{
		char:   char,
		status: none,
	}
}

// newGuess should take in a string and return a new guess.
// You should loop over each letter in the string and convert them to letter structs.

func newGuess(s string) guess {
	s = ToUpper(s)
	var g guess
	for i, char := range s {
		g[i] = newLetter(byte(char))
	}
	return g
}

// updateLettersWithWord updates the status of the letters in the guess based on a word
// pointer works kind of like self. scope issue. normally when you use a receiver / argument you're using a copy. with a pointer you're using the originals and changes persist beyond the function.
func (g *guess) updateLettersWithWord(word [wordSize]byte) {
	// guess -> array of letters
	for i := range g {
		guessLetter := &g[i] // this is a reference - variable is being set to that location
		if guessLetter.char == word[i] {
			guessLetter.status = correct
		} else {
			for j := range word {
				if guessLetter.char == word[j] {
					guessLetter.status = present
					break
				}
			}
			if guessLetter.status == none {
				guessLetter.status = absent
			}
		}
	}
}

func (g *guess) string() string {
	str := ""
	for _, l := range g {
		if 'A' <= l.char && l.char <= 'Z' {
			str += string(l.char)
		}
	}
	return str
}

// FROM https://www.tutorialspoint.com/golang-program-to-convert-a-string-into-uppercase:
// function to convert characters to upper case
func ToUpper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - ('a' - 'A')
		}
	}
	return string(b)
}

// appendGuess adds a guess to the wordleState. It returns an error
// if the guess is invalid.
func (w *wordleState) appendGuess(g guess) error {

	// Error if the maximum number of guesses has been reached:
	if w.currGuess >= 6 {
		return errors.New("Sorry! You've maxed out on guesses.")
	}

	// Error if the guess isn’t long enough:
	if len(g.string()) != 5 {
		return errors.New("Sorry! The guess needs to be 5 letters long")
	}

	// Error if the guess isn’t a valid word:
	if words.IsWord(g.string()) != true {
		return errors.New("Sorry! This is not a valid word.")
	}

	w.guesses[w.currGuess] = g
	w.currGuess += 1

	return nil

}

// isWordGuessed returns true when the latest guess is the correct word
func (w *wordleState) isWordGuessed() bool {
	lg := w.guesses[w.currGuess-1].string()
	wo := string(w.word[:])

	return lg == wo
}

// The game should end when the latest guess is correct or when
// there are no more empty slots left in the guesses array.
func (w *wordleState) shouldEndGame() bool {
	return w.currGuess == 6 || w.isWordGuessed()
}
