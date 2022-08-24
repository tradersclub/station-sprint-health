export const sum = (a:number, b:number) :number => {
  return a+b
}


interface Return {
  dictionary: Record<string,string>;
  isDictionary: boolean;
  positionsIterated: number;
}

export const isDictionary = (a: string, b:string): Return => {

  if (a?.length !== b?.length) {
    return {
      isDictionary: false,
      positionsIterated: 0,
      dictionary: {},
    };
  }
  
  const parsedA: Array<string> = a.split('');
  return parsedA.reduce<Return>((acc: Return, current: string, index: number) => {
    const dictionaryValue = acc.dictionary?.[current];
    acc.positionsIterated = index + 1;
    
    if (!dictionaryValue) {
      acc.dictionary[current] = b[index];
      return acc;
    }
    if (dictionaryValue !== b[index]) {
      acc.isDictionary = false;
      return acc;
    }

  }, {
    isDictionary: true,
    positionsIterated: 0,
    dictionary: {},
  });

}
