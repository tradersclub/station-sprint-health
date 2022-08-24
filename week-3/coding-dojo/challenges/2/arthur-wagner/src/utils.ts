interface Dictionary {
  dictionary: Record<string,string>;
  mappedLetters: Record<string,boolean>;
  isDictionary: boolean;
  positionsIterated: number;
}

export const isDictionary = (firstWord: string, secondWord:string): Dictionary => {
  const acc: Dictionary = {
    isDictionary: true,
    positionsIterated: 0,
    dictionary: {},
    mappedLetters: {}
  };

  const firstWordLength = firstWord.length;
  const secondWordLength = secondWord.length;
  
  if (firstWordLength !== secondWordLength) {
    acc.isDictionary = false;
    return  acc;
  }

  for (let i = 0; i < firstWordLength; i++) {
    const current = firstWord[i];
    const dictionaryValue = acc.dictionary[current];

    acc.positionsIterated = i + 1;

    if (!dictionaryValue) {
      if (acc.mappedLetters[secondWord[i]]) {
        acc.isDictionary = false
        return acc
      }
      acc.dictionary[current] = secondWord[i];
      acc.mappedLetters[secondWord[i]] = true
      continue;
    }
    if (dictionaryValue !== secondWord[i]) {
      acc.isDictionary = false;
      return acc;
    }
  }
  
  return acc;
}
