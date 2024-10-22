package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// spin up http server
// end point of your naming
// accepts a json payload of { name, age }

type Person struct {
	FullName string `json:"name"` // json tag
	Age int
}

func doStuff(w http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {		
		w.Write([]byte("Hello world"))
		return
	}
	// reading req.body & copy into a buffer
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, req.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// bytes of json --> struct representation of the json
	var p Person
	if err := json.Unmarshal(buffer.Bytes(), &p); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(p.FullName + strconv.Itoa(p.Age)))
	
}

func main() {
	http.HandleFunc("/", doStuff)
	http.HandleFunc("/something", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Something page"))
	})

	http.ListenAndServe(":3333", nil)
}