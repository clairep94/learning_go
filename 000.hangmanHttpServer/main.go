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
var guesses []string
var cache = make(map[string]int)

func homePage(w http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {		
		w.Write([]byte("Hello world"))
		return
	}
	// reading req.body & copy into a buffer
	// buffer := new(bytes.Buffer)
	// _, err := io.Copy(buffer, req.Body)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// bytes of json --> struct representation of the json
	// if err := json.Unmarshal(buffer.Bytes(), &p); err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// w.Write([]byte(p.FullName + strconv.Itoa(p.Age)))
	
}

func indexOfElem(elem string, arr []string) int {
	// TODO: first do a lookup against the cache
	for i, el := range arr {
		if elem == el {
			return i
		}
	}
	return -1
}

func splitWord(word string) []string {
	return strings.Split(word, "")
}

func DisplayConcealedWord(word string, guesses []string) string {
	display := []string{"Word:"}
	for _, letter := range splitWord(word) {
		if indexOfElem(letter, guesses) == -1 {
			display = append(display, "_")
		} else {
			display = append(display, letter)
		}
	}
	return strings.Join(display, " ")
}

func hangmanGame(w http.ResponseWriter, req *http.Request){

	if req.Method != "POST" {
		// Display the current game state
		response := DisplayConcealedWord(secretWord, guesses) + "\n"
		response += "Guesses: [" + strings.Join(guesses, ", ") + "]"
		w.Write([]byte(response))
		return
	}

	// reading req.body & copy into a buffer
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, req.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var guess string
	if err := json.Unmarshal(buffer.Bytes(), &guess); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	
	w.Write([]byte(guess))
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