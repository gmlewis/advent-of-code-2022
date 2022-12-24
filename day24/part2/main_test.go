package main

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
	"github.com/google/go-cmp/cmp"
)

func TestBlowWindNextStep(t *testing.T) {
	in := strings.TrimSpace(`
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#
`)
	want := []string{`
#.#####
#.....#
#.>...#
#.....#
#.....#
#...v.#
#####.#
`,
		`
#.#####
#...v.#
#..>..#
#.....#
#.....#
#.....#
#####.#
`,
		`
#.#####
#.....#
#...2.#
#.....#
#.....#
#.....#
#####.#
`,
		`
#.#####
#.....#
#....>#
#...v.#
#.....#
#.....#
#####.#
`,
		`
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#
`,
	}

	p := parsePuzzle(strings.Split(in, "\n"))
	if diff := cmp.Diff(in, p.String()); diff != "" {
		t.Errorf("puzzle mismatch (-want +got):\n%v\n%v", in, p)
	}

	for i, w := range want {
		p.blowWindNextStep()
		if diff := cmp.Diff(strings.TrimSpace(w), p.String()); diff != "" {
			t.Errorf("puzzle[%v] mismatch (-want +got):%v\n%v", i, w, p)
		}
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 54\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
`

/*
Initial state:
#E######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

Minute 1, move down:
#.######
#E>3.<.#
#<..<<.#
#>2.22.#
#>v..^<#
######.#

Minute 2, move down:
#.######
#.2>2..#
#E^22^<#
#.>2.^>#
#.>..<.#
######.#

Minute 3, wait:
#.######
#<^<22.#
#E2<.2.#
#><2>..#
#..><..#
######.#

Minute 4, move up:
#.######
#E<..22#
#<<.<..#
#<2.>>.#
#.^22^.#
######.#

Minute 5, move right:
#.######
#2Ev.<>#
#<.<..<#
#.^>^22#
#.2..2.#
######.#

Minute 6, move right:
#.######
#>2E<.<#
#.2v^2<#
#>..>2>#
#<....>#
######.#

Minute 7, move down:
#.######
#.22^2.#
#<vE<2.#
#>>v<>.#
#>....<#
######.#

Minute 8, move left:
#.######
#.<>2^.#
#.E<<.<#
#.22..>#
#.2v^2.#
######.#

Minute 9, move up:
#.######
#<E2>>.#
#.<<.<.#
#>2>2^.#
#.v><^.#
######.#

Minute 10, move right:
#.######
#.2E.>2#
#<2v2^.#
#<>.>2.#
#..<>..#
######.#

Minute 11, wait:
#.######
#2^E^2>#
#<v<.^<#
#..2.>2#
#.<..>.#
######.#

Minute 12, move down:
#.######
#>>.<^<#
#.<E.<<#
#>v.><>#
#<^v^^>#
######.#

Minute 13, move down:
#.######
#.>3.<.#
#<..<<.#
#>2E22.#
#>v..^<#
######.#

Minute 14, move right:
#.######
#.2>2..#
#.^22^<#
#.>2E^>#
#.>..<.#
######.#

Minute 15, move right:
#.######
#<^<22.#
#.2<.2.#
#><2>E.#
#..><..#
######.#

Minute 16, move right:
#.######
#.<..22#
#<<.<..#
#<2.>>E#
#.^22^.#
######.#

Minute 17, move down:
#.######
#2.v.<>#
#<.<..<#
#.^>^22#
#.2..2E#
######.#

Minute 18, move down:
#.######
#>2.<.<#
#.2v^2<#
#>..>2>#
#<....>#
######E#

*/
