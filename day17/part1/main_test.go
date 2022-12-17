package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestDropRock(t *testing.T) {
	tests := []struct {
		name    string
		puz     *puzT
		rockNum int
		want    *puzT
	}{
		{
			name:    "initial rock1 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, rock1(0, 0)),
		},
		{
			name:    "initial rock1 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, rock1(3, 0)),
		},
		{
			name:    "initial rock1 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, rock1(2, 0)),
		},
		{
			name:    "initial rock1 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 1,
			want:    genPuz(1, 4, rock1(2, 0)),
		},

		{
			name:    "initial rock2 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, rock2(0, 0)),
		},
		{
			name:    "initial rock2 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, rock2(4, 0)),
		},
		{
			name:    "initial rock2 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, rock2(2, 0)),
		},
		{
			name:    "initial rock2 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 2,
			want:    genPuz(3, 4, rock2(2, 0)),
		},

		{
			name:    "initial rock3 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, rock3(0, 0)),
		},
		{
			name:    "initial rock3 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, rock3(4, 0)),
		},
		{
			name:    "initial rock3 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, rock3(2, 0)),
		},
		{
			name:    "initial rock3 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 3,
			want:    genPuz(3, 4, rock3(2, 0)),
		},

		{
			name:    "initial rock4 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, rock4(0, 0)),
		},
		{
			name:    "initial rock4 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, rock4(6, 0)),
		},
		{
			name:    "initial rock4 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, rock4(2, 0)),
		},
		{
			name:    "initial rock4 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 4,
			want:    genPuz(4, 4, rock4(2, 0)),
		},

		{
			name:    "initial rock5 - all left",
			puz:     &puzT{gas: []rune("<<<<<<<<"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, rock5(0, 0)),
		},
		{
			name:    "initial rock5 - all right",
			puz:     &puzT{gas: []rune(">>>>>>>>"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, rock5(5, 0)),
		},
		{
			name:    "initial rock5 - even mixed gas - left right",
			puz:     &puzT{gas: []rune("><><><><"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, rock5(2, 0)),
		},
		{
			name:    "initial rock5 - even mixed gas - right left",
			puz:     &puzT{gas: []rune("<><><><>"), grid: map[keyT]rune{}},
			rockNum: 5,
			want:    genPuz(2, 4, rock5(2, 0)),
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
		gasIndex: gasIndex,
		grid:     map[keyT]rune{},
	}

	for _, addRock := range addRocks {
		addRock(puz)
	}

	return puz
}

/*
func TestExample(t *testing.T) {
	want := "Solution: 0\n"
	test.Runner(t, example1, want, process, &printf)
}
*/

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
`

/*
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##


The first rock begins falling:
|..@@@@.|
|.......|
|.......|
|.......|
+-------+

Jet of gas pushes rock right:
|...@@@@|
|.......|
|.......|
|.......|
+-------+

Rock falls 1 unit:
|...@@@@|
|.......|
|.......|
+-------+

Jet of gas pushes rock right, but nothing happens:
|...@@@@|
|.......|
|.......|
+-------+

Rock falls 1 unit:
|...@@@@|
|.......|
+-------+

Jet of gas pushes rock right, but nothing happens:
|...@@@@|
|.......|
+-------+

Rock falls 1 unit:
|...@@@@|
+-------+

Jet of gas pushes rock left:
|..@@@@.|
+-------+

Rock falls 1 unit, causing it to come to rest:
|..####.|
+-------+

A new rock begins falling:
|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|.......|
|..####.|
+-------+

Jet of gas pushes rock left:
|..@....|
|.@@@...|
|..@....|
|.......|
|.......|
|.......|
|..####.|
+-------+

Rock falls 1 unit:
|..@....|
|.@@@...|
|..@....|
|.......|
|.......|
|..####.|
+-------+

Jet of gas pushes rock right:
|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|..####.|
+-------+

Rock falls 1 unit:
|...@...|
|..@@@..|
|...@...|
|.......|
|..####.|
+-------+

Jet of gas pushes rock left:
|..@....|
|.@@@...|
|..@....|
|.......|
|..####.|
+-------+

Rock falls 1 unit:
|..@....|
|.@@@...|
|..@....|
|..####.|
+-------+

Jet of gas pushes rock right:
|...@...|
|..@@@..|
|...@...|
|..####.|
+-------+

Rock falls 1 unit, causing it to come to rest:
|...#...|
|..###..|
|...#...|
|..####.|
+-------+

A new rock begins falling:
|....@..|
|....@..|
|..@@@..|
|.......|
|.......|
|.......|
|...#...|
|..###..|
|...#...|
|..####.|
+-------+


|..@....|
|..@....|
|..@....|
|..@....|
|.......|
|.......|
|.......|
|..#....|
|..#....|
|####...|
|..###..|
|...#...|
|..####.|
+-------+

|..@@...|
|..@@...|
|.......|
|.......|
|.......|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@@@.|
|.......|
|.......|
|.......|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|.......|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|....@..|
|....@..|
|..@@@..|
|.......|
|.......|
|.......|
|..#....|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@....|
|..@....|
|..@....|
|..@....|
|.......|
|.......|
|.......|
|.....#.|
|.....#.|
|..####.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@...|
|..@@...|
|.......|
|.......|
|.......|
|....#..|
|....#..|
|....##.|
|....##.|
|..####.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@@@.|
|.......|
|.......|
|.......|
|....#..|
|....#..|
|....##.|
|##..##.|
|######.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

*/
