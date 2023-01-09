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
	for {
		word := <-data
		if !checkDuplicatedLetter(word) {
			wordsWithAlphabet = append(wordsWithAlphabet, word)
		}
	}
}

func checkDuplicatedLetter(word string) bool {
	for i, letter := range word {
		for k, next := range word[i : len(word)-1] {
			if i != k && letter == next {
				return true
			}
		}
	}
	return false
}
