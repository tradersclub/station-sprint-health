import Benchmark from 'benchmark'

export const options = { "async": true }

export function cycle(event:Benchmark.Event) {
    function round(num:number, decimal:number) {
        return Math.round(num*Math.pow(10,decimal))/Math.pow(10,decimal)
    }

    if (!event.target.hz) {
        return;
    }

    event.target.name

    const secondPerOperation = 1/event.target.hz
    const nanoSecondPerOperation = secondPerOperation*Math.pow(10,9)

    console.log(event.target.name, " => " , round(nanoSecondPerOperation, 2), "ns/op");
}

export function complete(this: Benchmark.Suite) {
    console.log('Fastest is => ', this.filter('fastest').map('name'));
}