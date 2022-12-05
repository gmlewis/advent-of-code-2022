package main

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
	"github.com/google/go-cmp/cmp"
)

func TestExample(t *testing.T) {
	want := "Solution: CMZ\n"
	test.Runner(t, example1, want, process, &printf)
}

func TestParseStacks(t *testing.T) {
	parts := strings.Split(example1, "\n\n")
	got := parseStacks(parts[0])
	want := puzT{
		"1": []string{"Z", "N"},
		"2": []string{"M", "C", "D"},
		"3": []string{"P"},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseStacks mismatch (-want +got):\n%v", diff)
	}
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
`
