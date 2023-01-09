package main

import (
	"encoding/json"
	"net/http"
)

var wordsWithAlphabet = make([]string, 0)

func main() {
	var chanNewWord = make(chan string, 100)
	go poolWords(chanNewWord)

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")
		chanNewWord <- word
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(wordsWithAlphabet)
	})

	http.ListenAndServe(":8080", nil)
}

func poolWords(data <-chan string) {
	str := <-data
	for i, r := range str {
		exist := false
		for k, h := range str[i : len(str)-1] {

		}
	}
}
