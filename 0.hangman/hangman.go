package main

import (
	"strings"
)

type Hangman struct {
	Word string
	MaxGuesses int
	Guesses []string
}

func (h *Hangman) IsInGuess(letter string) bool{
	for _, l := range h.Guesses {
		if letter == l {
			return true
		} else {
			continue
		}
	}
	return false
}

func (h *Hangman) DisplayWord() string{
	letters := strings.Split(h.Word, "")
	result := []string{}
	for _, l := range letters {
		if h.IsInGuess(l){
			result = append(result, l)
		} else {
			result = append(result, ".")
		}
	}
	return strings.Join(result, "")
}

func (h *Hangman) CheckWin() bool{
	if h.DisplayWord() == h.Word{
		return true
	} else {
		return false
	}
}

func (h *Hangman) IsEligibleForNextRound() bool{
	if len(h.Guesses) <
}

func (h *Hangman) GuessLetter(letter string) {
	h.Guesses = append(h.Guesses, letter)
}

// func (h *Hangman) SubmitGuess(guess string) {
// 	h.Guesses -= 1
	
// }