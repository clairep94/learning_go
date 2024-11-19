package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	// "strconv"
)
const secretWord = "hidden"
const maxGuesses = 8
var guesses []string
// var cache = make(map[string]int)

type Guess struct {
	Letter string `json:"letter"`
}

func homePage(w http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {		
		w.Write([]byte("Hello world"))
		return
	}
}

// TODO: add lookup against cache, how to do this with Go if it'll return 0 if not in the map instead of nil?
func indexOfElem(elem string, arr []string) int {
	// if _, ok := internationalisation[country]; !ok {
	// 	country = "en"
	// }
	for i, el := range arr {
		if elem == el {
			return i
		}
	}
	return -1
}

func splitWord() []string {
	return strings.Split(secretWord, "")
}

func DisplayConcealedWord() string {
	display := []string{"Word:"}
	for _, letter := range splitWord() {
		if indexOfElem(letter, guesses) == -1 {
			display = append(display, "_")
		} else {
			display = append(display, letter)
		}
	}
	return strings.Join(display, " ")
}

func DisplayCurrentGamestate() string {
	response := DisplayConcealedWord() + "\n"
	response += "Guesses: [" + strings.Join(guesses, ", ") + "]"
	return response
}

// func CheckGame() bool {

// }

func isGameLost() bool {
	return len(guesses) >= maxGuesses
}

func isGameWon() bool {

}

func hangmanGame(w http.ResponseWriter, req *http.Request){

	if req.Method != "POST" {
		// Display the current game state
		w.Write([]byte(DisplayCurrentGamestate()))
		return
	}

	// TODO: add win condition
	// check if already won/lost 
	if isGameLost() {
		http.Error(w, "Game already over", http.StatusBadRequest)
		return
	}

	// reading req.body & copy into a buffer
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, req.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var g Guess
	// valid json check
	if err := json.Unmarshal(buffer.Bytes(), &g); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		
	// single char check
	if len(g.Letter) != 1 {
		http.Error(w, "Invalid guess: must be a single character", http.StatusBadRequest)
		return
	}

	g.Letter = strings.ToLower(g.Letter)

	// already guessed check
	if indexOfElem(g.Letter, guesses) != -1 {
		http.Error(w, "Invalid guess: already guessed letter", http.StatusBadRequest)
		return
	} else {
		guesses = append(guesses, g.Letter)		
	}

	// lastTurn message
	lastTurnMessage := g.Letter
	if indexOfElem(g.Letter, splitWord()) == -1 {
		lastTurnMessage += " is wrong!"
	} else {
		lastTurnMessage += " is in the word!"
	}

	var checkGameMessage string
	if isGameLost() {
		checkGameMessage = "You lost!"
		// add win condition
	} else {
		checkGameMessage = "Try again!"
	}

	w.Write([]byte(DisplayCurrentGamestate() + "\n" + lastTurnMessage + "\n" + checkGameMessage))
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hangman", hangmanGame)

	port := ":3333"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}