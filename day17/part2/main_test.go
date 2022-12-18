package main

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func r1(x, y int) RenderRock { return rock1(x, y, true) }
func r2(x, y int) RenderRock { return rock2(x, y, true) }
func r3(x, y int) RenderRock { return rock3(x, y, true) }
func r4(x, y int) RenderRock { return rock4(x, y, true) }
func r5(x, y int) RenderRock { return rock5(x, y, true) }

func TestDropRock(t *testing.T) {
	tests := []struct {
		name    string
		puz     *puzT
		rockNum int
		want    *puzT
	}{
		{
			name:    "example 1 - rock 2",
			puz:     genPuz(1, 4, r1(2, 0)),
			rockNum: 2,
			want:    genPuz(4, 8, r1(2, 0), r2(2, 1)),
		},
		{
			name:    "example 1 - rock 3",
			puz:     genPuz(4, 8, r1(2, 0), r2(2, 1)),
			rockNum: 3,
			want:    genPuz(6, 13, r1(2, 0), r2(2, 1), r3(0, 3)),
		},
		{
			name:    "example 1 - rock 4",
			puz:     genPuz(6, 13, r1(2, 0), r2(2, 1), r3(0, 3)),
			rockNum: 4,
			want:    genPuz(7, 20, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3)),
		},
		{
			name:    "example 1 - rock 5",
			puz:     genPuz(7, 20, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3)),
			rockNum: 5,
			want:    genPuz(9, 24, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7)),
		},
		{
			name:    "example 1 - rock 6",
			puz:     genPuz(9, 24, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7)),
			rockNum: 6,
			want:    genPuz(10, 28, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9)),
		},
		{
			name:    "example 1 - rock 7",
			puz:     genPuz(10, 28, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9)),
			rockNum: 7,
			want:    genPuz(13, 32, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10)),
		},
		{
			name:    "example 1 - rock 8",
			puz:     genPuz(13, 32, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10)),
			rockNum: 8,
			want:    genPuz(15, 37, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10), r3(3, 12)),
		},
		{
			name:    "example 1 - rock 9",
			puz:     genPuz(15, 37, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10), r3(3, 12)),
			rockNum: 9,
			want:    genPuz(17, 3, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10), r3(3, 12), r4(4, 13)),
		},
		{
			name:    "example 1 - rock 10",
			puz:     genPuz(17, 3, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10), r3(3, 12), r4(4, 13)),
			rockNum: 10,
			want:    genPuz(17, 12, r1(2, 0), r2(2, 1), r3(0, 3), r4(4, 3), r5(4, 7), r1(1, 9), r2(1, 10), r3(3, 12), r4(4, 13), r5(0, 12)),
		},

		{
			name:    "initial rock1 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, r1(0, 0)),
		},
		{
			name:    "initial rock1 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, r1(3, 0)),
		},
		{
			name:    "initial rock1 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, r1(2, 0)),
		},
		{
			name:    "initial rock1 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, r1(2, 0)),
		},

		{
			name:    "initial rock2 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, r2(0, 0)),
		},
		{
			name:    "initial rock2 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, r2(4, 0)),
		},
		{
			name:    "initial rock2 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, r2(2, 0)),
		},
		{
			name:    "initial rock2 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, r2(2, 0)),
		},

		{
			name:    "initial rock3 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, r3(0, 0)),
		},
		{
			name:    "initial rock3 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, r3(4, 0)),
		},
		{
			name:    "initial rock3 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, r3(2, 0)),
		},
		{
			name:    "initial rock3 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, r3(2, 0)),
		},

		{
			name:    "initial rock4 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, r4(0, 0)),
		},
		{
			name:    "initial rock4 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, r4(6, 0)),
		},
		{
			name:    "initial rock4 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, r4(2, 0)),
		},
		{
			name:    "initial rock4 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, r4(2, 0)),
		},

		{
			name:    "initial rock5 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, r5(0, 0)),
		},
		{
			name:    "initial rock5 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, r5(5, 0)),
		},
		{
			name:    "initial rock5 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, r5(2, 0)),
		},
		{
			name:    "initial rock5 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, r5(2, 0)),
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.puz.dropRock(tt.rockNum - 1) // 0-index

			if got, want := tt.puz.height, tt.want.height; got != want {
				t.Errorf("dropRock(%v) height = %v, want %v", tt.rockNum, got, want)
			}

			if got, want := tt.puz.gasIndex, tt.want.gasIndex; got != want {
				t.Errorf("dropRock(%v) gasIndex = %v, want %v", tt.rockNum, got, want)
			}

			if got, want := tt.puz.String(), tt.want.String(); got != want {
				t.Errorf("dropRock(%v) puzzle got:\n%v\nwant:\n%v", tt.rockNum, got, want)
			}
		})
	}
}

func genPuz(height, gasIndex int, addRocks ...RenderRock) *puzT {
	puz := &puzT{
		height:   height,
		gas:      []rune(strings.TrimSpace(example1)),
		gasIndex: gasIndex,
		grid:     map[keyT]rune{},
	}

	for _, addRock := range addRocks {
		addRock(puz)
	}

	return puz
}

func TestExample(t *testing.T) {
	want := "Solution: 1514285714288\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
`
