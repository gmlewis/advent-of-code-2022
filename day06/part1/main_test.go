package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 7\n"
	test.Runner(t, example1, want, process, &printf)

	want = "Solution: 5\n"
	test.Runner(t, example2, want, process, &printf)

	want = "Solution: 6\n"
	test.Runner(t, example3, want, process, &printf)

	want = "Solution: 10\n"
	test.Runner(t, example4, want, process, &printf)

	want = "Solution: 11\n"
	test.Runner(t, example5, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb
`

var example2 = `bvwbjplbgvbhsrlpgdmjqwftvncz
`

var example3 = `nppdvjthqldpwncqszvftbrmjlhg
`

var example4 = `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
`

var example5 = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw
`
