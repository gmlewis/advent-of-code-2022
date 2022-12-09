package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 13\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
