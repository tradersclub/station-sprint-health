package main

import (
	"encoding/json"
	"net/http"
)

var wordsWithAlphabet = make([]string, 0)

func main() {
	var chanAddNewWord = make(chan string, 100)

	var chanNewWord = make(chan string, 100)
	go poolWords(chanNewWord, chanAddNewWord)

	go poolAddOnSlice(chanAddNewWord)

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")
		chanNewWord <- word
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(wordsWithAlphabet)
	})

	go http.ListenAndServe(":8080", nil)

	close(chanNewWord)
}

func poolWords(data <-chan string, addNewWord chan<- string) {
	for {
		newWord, ok := <-data
		if !ok {
			break
		}

		if !isRepeatedLetters(newWord) {
			addNewWord <- newWord
		}

	}
}

func poolAddOnSlice(data <-chan string) {
	for {
		newWord, ok := <-data
		if !ok {
			break
		}
		wordsWithAlphabet = append(wordsWithAlphabet, newWord)
	}
}

// Nˆ2
func isRepeatedLetters(word string) bool {
	var mapLetters = make(map[rune]bool)

	for _, v := range word {
		if mapLetters[v] {
			return true
		}
		mapLetters[v] = true
	}

	return false
}
