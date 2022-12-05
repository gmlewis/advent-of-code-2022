package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 0\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2


[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 


        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3


        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3


        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3

`
