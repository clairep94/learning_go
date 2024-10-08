package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIsInGuess(t *testing.T){
	t.Run(
		"When the letter is in Guesses",
		func(t *testing.T){
			hangman := Hangman{Word:"tester", MaxGuesses: 7, Guesses:[]string{"t", "a", "b"}}
			got := hangman.IsInGuess("t")
			want := true
			if got != want {
				t.Error("got: false, want: true")
			}
		})
	t.Run(
		"When the letter is not in Guesses",
		func(t *testing.T){
			hangman := Hangman{Word:"tester", MaxGuesses: 7, Guesses:[]string{"t", "a", "b"}}
			got := hangman.IsInGuess("c")
			want := false
			if got != want {
				t.Error("got: true, want: false")
			}
		})
}

func TestDisplayWord(t *testing.T){
	t.Run(
		"When there are no guesses",
		func(t *testing.T) {
			hangman := Hangman{"tester", 7, nil}
			got := hangman.DisplayWord()
			want := "......"
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"When there is one matching letter",
		func(t *testing.T) {
			hangman := Hangman{"tester", 7, []string{"t"}}
			got := hangman.DisplayWord()
			want := "t..t.."
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"When there are more matching letters",
		func(t *testing.T) {
			hangman := Hangman{"tester", 7, []string{"t", "e", "a", "s"}}
			got := hangman.DisplayWord()
			want := "teste."
			assertCorrectMessage(t, got, want)
		})
}

func TestCheckWin(t *testing.T) {
	t.Run(
		"When there is not a win",
		func(t *testing.T) {
			hangman := Hangman{"tester", 7, []string{"t", "e", "a", "s"}}
			got := hangman.CheckWin()
			want := false
			if got != want {
				t.Error("got: true, want: false")
			}
		})
	t.Run(
		"When there is a win",
		func(t *testing.T) {
			hangman := Hangman{"tester", 7, []string{"t", "e", "a", "s", "r"}}
			got := hangman.CheckWin()
			want := true
			if got != want {
				t.Error("got: false, want: true")
			}
		})
}

// func TestCheckAnswer(t *testing.T){
// 	hangman := Hangman{"test", 7, []}

// 	hangman.SubmitGuess("n")

// 	got := hangman.CheckAnswer()
// 	want := `Wrong!
// guesses: n | word: .... | errors: 1/7`
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }