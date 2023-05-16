package wordle

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
	var g guess
	for i, char := range s {
		g[i] = newLetter(byte(char))
	}
	return g
}
