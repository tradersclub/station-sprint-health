export const sum = (a:number, b:number) :number => {
  return a+b
}

export const shuffleList = (items: string[]) => items
  .map(value => ({ value, sort: Math.random() }))
  .sort((a, b) => a.sort - b.sort)
  .map(({ value }) => value)