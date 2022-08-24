interface Dictionary {
  dictionary: Record<string,string>;
  mappedLetters: Record<string,boolean>;
  isDictionary: boolean;
  positionsIterated: number;
}

export const isDictionary = (firstWord: string, secondWOrd:string): Dictionary => {
  if (firstWord.length !== secondWOrd.length) {
    return {
      isDictionary: false,
      positionsIterated: 0,
      dictionary: {},
      mappedLetters: {}
    };
  }
  
  const parsedA: Array<string> = firstWord.split('');

  return parsedA.reduce<Dictionary>((acc: Dictionary, current: string, index: number) => {
    const dictionaryValue = acc.dictionary[current];
    acc.positionsIterated = index + 1;

    if (!dictionaryValue) {
      if (acc.mappedLetters[secondWOrd[index]]) {
        acc.isDictionary = false
        return acc
      }
      acc.dictionary[current] = secondWOrd[index];
      acc.mappedLetters[secondWOrd[index]] = true
      return acc;
    }
    if (dictionaryValue !== secondWOrd[index]) {
      acc.isDictionary = false;
      return acc;
    }

    return acc;
  }, {
    isDictionary: true,
    positionsIterated: 0,
    dictionary: {},
    mappedLetters: {},
  });
}
