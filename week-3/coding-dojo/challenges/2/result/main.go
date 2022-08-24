package main

import (
	"fmt"
)

func main() {
	result := isIsomorphic("foo", "bar")
	fmt.Println(result)

	result2 := isIsomorphic("add", "egg")
	fmt.Println(result2)

	result3 := isIsomorphic("paper", "title")
	fmt.Println(result3)

}

func isIsomorphic(word1 string, word2 string) bool {
	return isAlphabetSubstituted(word1, word2) && isAlphabetSubstituted(word2, word1)
}

func isAlphabetSubstituted(word1 string, word2 string) bool {
	alphabet := make(map[byte]byte)

	for index := range word1 {

		if word, ok := alphabet[word2[index]]; ok && word != word1[index] {
			return false
		}
		alphabet[word2[index]] = word1[index]
	}

	return true
}
