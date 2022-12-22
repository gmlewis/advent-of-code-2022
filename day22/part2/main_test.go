package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestWrappedPos(t *testing.T) {
	tests := []struct {
		name string
		pos  posT
		want posT
	}{
		{
			name: "face 2 to face 5 - right (top) to up",
			pos:  posT{dx: 1, dy: 0, key: keyT{50, 151}},
			want: posT{dx: 0, dy: -1, key: keyT{51, 150}},
		},
		{
			name: "face 2 to face 5 - right (bottom) to up",
			pos:  posT{dx: 1, dy: 0, key: keyT{50, 200}},
			want: posT{dx: 0, dy: -1, key: keyT{100, 150}},
		},
		{
			name: "face 2 to face 1 - left (top) to down",
			pos:  posT{dx: -1, dy: 0, key: keyT{1, 151}},
			want: posT{dx: 0, dy: 1, key: keyT{51, 1}},
		},
		{
			name: "face 2 to face 1 - left (bottom) to down",
			pos:  posT{dx: -1, dy: 0, key: keyT{1, 200}},
			want: posT{dx: 0, dy: 1, key: keyT{100, 1}},
		},
		{
			name: "face 2 to face 6 - down (left) to down",
			pos:  posT{dx: 0, dy: 1, key: keyT{1, 200}},
			want: posT{dx: 0, dy: 1, key: keyT{101, 1}},
		},
		{
			name: "face 2 to face 6 - down (right) to down",
			pos:  posT{dx: 0, dy: 1, key: keyT{50, 200}},
			want: posT{dx: 0, dy: 1, key: keyT{150, 1}},
		},
		{
			name: "face 5 to face 6 - right (top) to left",
			pos:  posT{dx: 1, dy: 0, key: keyT{100, 101}},
			want: posT{dx: -1, dy: 0, key: keyT{150, 50}},
		},
		{
			name: "face 5 to face 6 - right (bottom) to left",
			pos:  posT{dx: 1, dy: 0, key: keyT{100, 150}},
			want: posT{dx: -1, dy: 0, key: keyT{150, 1}},
		},
		{
			name: "face 3 to face 1 - left (top) to right",
			pos:  posT{dx: -1, dy: 0, key: keyT{1, 101}},
			want: posT{dx: 1, dy: 0, key: keyT{51, 50}},
		},
		{
			name: "face 3 to face 1 - left (bottom) to right",
			pos:  posT{dx: -1, dy: 0, key: keyT{1, 150}},
			want: posT{dx: 1, dy: 0, key: keyT{51, 1}},
		},
		{
			name: "face 5 to face 2 - down (left) to left",
			pos:  posT{dx: 0, dy: 1, key: keyT{51, 150}},
			want: posT{dx: -1, dy: 0, key: keyT{50, 151}},
		},
		{
			name: "face 5 to face 2 - down (right) to left",
			pos:  posT{dx: 0, dy: 1, key: keyT{100, 150}},
			want: posT{dx: -1, dy: 0, key: keyT{50, 200}},
		},
		{
			name: "face 3 to face 4 - up (left) to right",
			pos:  posT{dx: 0, dy: -1, key: keyT{1, 101}},
			want: posT{dx: 1, dy: 0, key: keyT{51, 51}},
		},
		{
			name: "face 3 to face 4 - up (right) to right",
			pos:  posT{dx: 0, dy: -1, key: keyT{50, 101}},
			want: posT{dx: 1, dy: 0, key: keyT{51, 100}},
		},
		{
			name: "face 4 to face 6 - right (top) to up",
			pos:  posT{dx: 1, dy: 0, key: keyT{100, 51}},
			want: posT{dx: 0, dy: -1, key: keyT{101, 50}},
		},
		{
			name: "face 4 to face 6 - right (bottom) to up",
			pos:  posT{dx: 1, dy: 0, key: keyT{100, 100}},
			want: posT{dx: 0, dy: -1, key: keyT{150, 50}},
		},
		{
			name: "face 4 to face 3 - left (top) to down",
			pos:  posT{dx: -1, dy: 0, key: keyT{51, 51}},
			want: posT{dx: 0, dy: 1, key: keyT{1, 101}},
		},
		{
			name: "face 4 to face 3 - left (bottom) to down",
			pos:  posT{dx: -1, dy: 0, key: keyT{51, 100}},
			want: posT{dx: 0, dy: 1, key: keyT{50, 101}},
		},
		{
			name: "face 6 to face 5 - right (top) to left",
			pos:  posT{dx: 1, dy: 0, key: keyT{150, 1}},
			want: posT{dx: -1, dy: 0, key: keyT{100, 150}},
		},
		{
			name: "face 6 to face 5 - right (bottom) to left",
			pos:  posT{dx: 1, dy: 0, key: keyT{150, 50}},
			want: posT{dx: -1, dy: 0, key: keyT{100, 101}},
		},
		{
			name: "face 1 to face 3 - left (top) to right",
			pos:  posT{dx: -1, dy: 0, key: keyT{51, 1}},
			want: posT{dx: 1, dy: 0, key: keyT{1, 150}},
		},
		{
			name: "face 1 to face 3 - left (bottom) to right",
			pos:  posT{dx: -1, dy: 0, key: keyT{51, 50}},
			want: posT{dx: 1, dy: 0, key: keyT{1, 101}},
		},
		{
			name: "face 6 to face 4 - down (left) to left",
			pos:  posT{dx: 0, dy: 1, key: keyT{101, 50}},
			want: posT{dx: -1, dy: 0, key: keyT{100, 51}},
		},
		{
			name: "face 6 to face 4 - down (right) to left",
			pos:  posT{dx: 0, dy: 1, key: keyT{150, 50}},
			want: posT{dx: -1, dy: 0, key: keyT{100, 100}},
		},
		{
			name: "face 6 to face 2 - up (left) to up",
			pos:  posT{dx: 0, dy: -1, key: keyT{101, 1}},
			want: posT{dx: 0, dy: -1, key: keyT{1, 200}},
		},
		{
			name: "face 6 to face 2 - up (right) to up",
			pos:  posT{dx: 0, dy: -1, key: keyT{150, 1}},
			want: posT{dx: 0, dy: -1, key: keyT{50, 200}},
		},
		{
			name: "face 1 to face 3 - left (top) to right",
			pos:  posT{dx: -1, dy: 0, key: keyT{51, 1}},
			want: posT{dx: 1, dy: 0, key: keyT{1, 150}},
		},
		{
			name: "face 1 to face 3 - left (bottom) to right",
			pos:  posT{dx: -1, dy: 0, key: keyT{51, 50}},
			want: posT{dx: 1, dy: 0, key: keyT{1, 101}},
		},
		{
			name: "face 1 to face 2 - up (left) to right",
			pos:  posT{dx: 0, dy: -1, key: keyT{51, 1}},
			want: posT{dx: 1, dy: 0, key: keyT{1, 151}},
		},
		{
			name: "face 1 to face 2 - up (right) to right",
			pos:  posT{dx: 0, dy: -1, key: keyT{100, 1}},
			want: posT{dx: 1, dy: 0, key: keyT{1, 200}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &puzT{maxCol: 150}
			got := p.wrappedPos(tt.pos)

			if got.key != tt.want.key || got.dx != tt.want.dx || got.dy != tt.want.dy {
				t.Errorf("wrappedPos = %v, want %v", got, tt.want)
			}

			// now go back the opposite direction
			got.dx = -got.dx
			got.dy = -got.dy
			got = p.wrappedPos(got)
			if got.key != tt.pos.key {
				t.Errorf("opposite direction: wrappedPos = %v, want %v", got.key, tt.pos.key)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 5031\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

/*
        ...#
        .#..
        #...
        ....
...#.D.....#
........#...
B.#....#...A
.....C....#.
        ...#....
        .....#..
        .#......
        ......#.


        >>v#
        .#v.
        #.v.
        ..v.
...#...v..v#
>>>v...>#.>>
..#v...#....
...>>>>v..#.
        ...#....
        .....#..
        .#......
        ......#.
*/
