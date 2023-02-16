import {sum} from "./utils";
import Benchmark from 'benchmark'
import {cycle, complete, options} from './utils/benchmark'

const suite = new Benchmark.Suite;

suite
.add('sum 1 + 2', function() {

  sum(1, 2) // CHANGE HERE <------

})
.on('cycle', cycle)
.on('complete', complete)
.run(options);
