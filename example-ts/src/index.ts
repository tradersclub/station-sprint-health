import {sum} from "./utils";
import Benchmark from 'benchmark'

var suite = new Benchmark.Suite;

suite
// bench test
.add('test isDictionary', function() {
  sum(1, 2)
})
// print results
.on('cycle', function(event: Benchmark.Event) {
  console.log(String(event.target));
})
.on('complete', function(this: Benchmark.Suite) {
  console.log('Fastest is ' + this.filter('fastest').map('name'));
})
.run({ 'async': true });
