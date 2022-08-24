function isAlphabetSubstituted(word1: string, word2: string) {
    let alphabet = new Map()
  
    for (let index = 0; index < word1.length; index++) {
        const letterWanted = alphabet.get(word2[index])
  
        if (letterWanted && letterWanted != word1[index]) {
            return false
        }
  
        alphabet.set(word2[index], word1[index])
    }
  
    return true
}
  
export function isIsomorphic(word1: string, word2: string) {
    return isAlphabetSubstituted(word1, word2) && isAlphabetSubstituted(word2, word1)
};
  
export function isIsomorphic2(word1: string, word2: string) {
    const alphabet = new Map()
    const mappedLetters = new Map()
  
    for (let index = 0; index < word1.length; index++) {
        const letterWanted = alphabet.get(word2[index])
  
        if (letterWanted && letterWanted != word1[index]) {
            return false
        }
  
        if (!letterWanted && mappedLetters.get(word1[index]) === true) {
          return false
        }
  
        mappedLetters.set(word1[index], true)
        alphabet.set(word2[index], word1[index])
    }
  
    return true
  };

  export function isIsomorphic3(word1: string, word2: string) {
    const alphabet:Record<string,string> = {}
    const mappedLetters:Record<string,boolean> = {}
  
    for (let index = 0; index < word1.length; index++) {
        const letterWanted = alphabet[word2[index]]
  
        if (letterWanted && letterWanted != word1[index]) {
            return false
        }
  
        if (!letterWanted && mappedLetters[word1[index]] === true) {
          return false
        }
  
        mappedLetters[word1[index]] = true
        alphabet[word2[index]] = word1[index]
    }
  
    return true
  };