import {isDictionary} from "./utils";
import Benchmark from 'benchmark'

var suite = new Benchmark.Suite;

function isAlphabetSubstituted(word1: string, word2: string) {
  let alphabet = new Map()
  
  for (let index in word1 as unknown as string[]) {
      const letterWanted = alphabet.get(word2[index])

      if (letterWanted && letterWanted != word1[index]) {
          return false
      }

      alphabet.set(word2[index], word1[index])
  }

  return true
}

function isIsomorphic(word1: string, word2: string) {
  return isAlphabetSubstituted(word1, word2) && isAlphabetSubstituted(word2, word1)
};

suite.add('test isDictionary', function() {
  isDictionary('paper', 'title')
})
.add('test isIsomorphic', function() {
  isIsomorphic('paper', 'title')
})
.on('cycle', function(event: Benchmark.Event) {
console.log(String(event.target));
})
.on('complete', function(this: Benchmark.Suite) {
console.log('Fastest is ' + this.filter('fastest').map('name'));
})
.run({ 'async': true });
