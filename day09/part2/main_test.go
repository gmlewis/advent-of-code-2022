package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 1\n"
	test.Runner(t, example1, want, process, &printf)
	want = "Solution: 36\n"
	test.Runner(t, example2, want, process, &printf)
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

var example2 = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
