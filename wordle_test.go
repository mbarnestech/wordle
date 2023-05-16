package wordle

import (
	"fmt"
	"testing"
)

func TestNewWordleState(t *testing.T) {
	word := "HELLO"
	ws := newWordleState(word)
	wordleAsString := string(ws.word[:])

	t.Logf("	Created wordleState:")
	t.Logf("    word: %s", wordleAsString)
	t.Logf("    guesses: %v", ws.guesses)
	t.Logf("    currGuess: %v", ws.currGuess)

	if wordleAsString != word {
		t.Errorf("Expected word to be %s but got %s", word, wordleAsString)
	}
}

func statusToString(status letterStatus) string {
	switch status {
	case none:
		return "none"
	case correct:
		return "correct"
	case present:
		return "present"
	case absent:
		return "absent"
	default:
		return "unknown"
	}
}

func TestNewGuess(t *testing.T) {
	s := "HELLO"
	g := newGuess(s)

	if len(g) != wordSize {
		t.Errorf("Expected length of guess to be %d, but got %d", wordSize, len(g))
	}

	for i, le := range g {
		lchar, lstatus := le.char, statusToString(le.status)

		t.Logf("	Created Letter:")
		t.Logf("    letter: %c", lchar)
		t.Logf("    status: %s", lstatus)

		if lchar == 0 {
			t.Errorf("Expected length of guess to be %d but got %d", wordSize, i)
			break
		}
		if lchar != byte(s[i]) {
			t.Errorf("Expected char to be %c but got %c", byte(s[i]), lchar)
		}
		if lstatus != "none" {
			t.Errorf("Expected status to be 'none' but got %s", lstatus)
		}

	}

}

func TestUpdateLettersWithWord(t *testing.T) {
	s := "VIOLA"
	g := newGuess(s)
	secretword := "HELLO"
	word := [wordSize]byte([]byte(secretword))
	g.updateLettersWithWord(word)

	statuses := []letterStatus{
		absent,  // "V" is not in "HELLO"
		absent,  // "I" is not in "HELLO"
		present, // "O" is in "HELLO" but not in the correct position
		correct, // "L" is in "HELLO" and in the correct position
		absent,  // "A" is not in "HELLO"
	}

	for i, char := range g {
		if char.status != statuses[i] {
			t.Errorf("Expected %s but got %s", statusToString(statuses[i]), statusToString(char.status))
		}
	}
}

func TestAppendGuess(t *testing.T) {

	s := "viola"
	g := newGuess(s)
	ws := newWordleState("HELLO")

	// checks for if the max guess has been reached
	ws.currGuess = 6
	fmt.Println(ws.appendGuess(g))

	// reset currGuess
	ws.currGuess = 1

	// check for length error
	s = "hill"
	g = newGuess(s)
	fmt.Println(ws.appendGuess(g))

	// check for invalid word error
	s1 := "hgdle"
	g1 := newGuess(s1)
	fmt.Println(ws.appendGuess(g1))

	// check the guess array contains the new guess
	s = "hello"
	g = newGuess(s)
	ws.appendGuess(g)

	if ws.guesses[ws.currGuess-1] != g {
		t.Errorf("Expected word to be: %s but got: %s", g.string(), ws.guesses[ws.currGuess-1].string())
	}

}

func TestIsWordGuessed(t *testing.T) {
	s := "hello"
	g := newGuess(s)
	ws := newWordleState("HELLO")
	ws.appendGuess(g)

	if !ws.isWordGuessed() {
		t.Errorf("Expected ws.appendGuess(g) to be true")
	}

}

func TestShouldEndGame(t *testing.T) {

	s := "hello"
	g := newGuess(s)
	ws := newWordleState("HELLO")
	ws.appendGuess(g)
	if !ws.shouldEndGame() {
		t.Errorf("Expected ws.shouldEndGame() to be true")
	}

	s1 := "VIOLA"
	g1 := newGuess(s1)
	ws1 := newWordleState("HELLO")
	ws1.appendGuess(g1)
	ws1.currGuess = 6
	if !ws1.shouldEndGame() {
		t.Errorf("Expected ws.shouldEndGame() to be true")
	}
}
