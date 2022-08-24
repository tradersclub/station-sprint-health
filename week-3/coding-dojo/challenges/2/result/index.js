function isAlphabetSubstituted(word1, word2) {
    let alphabet = new Map()
    
    for (let index in word1) {
        
        const letterWanted = alphabet.get(word2[index])

        if (letterWanted && letterWanted != word1[index]) {
			return false
		}

        alphabet.set(word2[index], word1[index])
    }

	return true
}

function isIsomorphic(word1, word2) {
    return isAlphabetSubstituted(word1, word2) && isAlphabetSubstituted(word2, word1)
};


console.log(isIsomorphic("foo", "bar"))
console.log(isIsomorphic("add", "egg"))
console.log(isIsomorphic("paper", "title"))
