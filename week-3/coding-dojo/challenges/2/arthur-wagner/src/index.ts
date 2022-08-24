import {isDictionary} from "./utils";
import {isIsomorphic, isIsomorphic2, isIsomorphic3} from "./utils_leo";
import Benchmark from 'benchmark'

var suite = new Benchmark.Suite;

suite.add('test isDictionary', function() {
  isDictionary('paper', 'title')
})
.add('test isIsomorphic', function() {
  isIsomorphic('paper', 'title')
})
.add('test isIsomorphic2', function() {
  isIsomorphic2('paper', 'title')
})
.add('test isIsomorphic3', function() {
  isIsomorphic3('paper', 'title')
})
.on('cycle', function(event: Benchmark.Event) {
  console.log(String(event.target));
})
.on('complete', function(this: Benchmark.Suite) {
  console.log('Fastest is ' + this.filter('fastest').map('name'));
})
.run({ 'async': true });
