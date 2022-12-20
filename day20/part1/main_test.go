package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
	"github.com/google/go-cmp/cmp"
)

func TestShift(t *testing.T) {
	tests := []struct {
		name    string
		num     int
		indices []int
		want    []int
	}{
		{
			name:    "1 moves between 2 and -3",
			num:     1,
			indices: []int{1, 2, -3, 3, -2, 0, 4},
			want:    []int{2, 1, -3, 3, -2, 0, 4},
		},
		{
			name:    "2 moves between -3 and 3",
			num:     2,
			indices: []int{2, 1, -3, 3, -2, 0, 4},
			want:    []int{1, -3, 2, 3, -2, 0, 4},
		},
		{
			name:    "-3 moves between -2 and 0",
			num:     -3,
			indices: []int{1, -3, 2, 3, -2, 0, 4},
			want:    []int{1, 2, 3, -2, -3, 0, 4},
		},
		{
			name:    "3 moves between 0 and 4",
			num:     3,
			indices: []int{1, 2, 3, -2, -3, 0, 4},
			want:    []int{1, 2, -2, -3, 0, 3, 4},
		},
		{
			name:    "-2 moves between 4 and 1",
			num:     -2,
			indices: []int{1, 2, -2, -3, 0, 3, 4},
			want:    []int{1, 2, -3, 0, 3, 4, -2},
		},
		{
			name:    "0 does not move",
			num:     0,
			indices: []int{1, 2, -3, 0, 3, 4, -2},
			want:    []int{1, 2, -3, 0, 3, 4, -2},
		},
		{
			name:    "4 moves between -3 and 0",
			num:     4,
			indices: []int{1, 2, -3, 0, 3, 4, -2},
			want:    []int{1, 2, -3, 4, 0, 3, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indices := nums2map(tt.indices)
			shift(tt.num, indices)
			_, got := map2nums(indices)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("shift mismatch (-want +got):\n%v", diff)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 3\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
1
2
-3
3
-2
0
4
`

/*
Initial arrangement:
1, 2, -3, 3, -2, 0, 4

1 moves between 2 and -3:
2, 1, -3, 3, -2, 0, 4

2 moves between -3 and 3:
1, -3, 2, 3, -2, 0, 4

-3 moves between -2 and 0:
1, 2, 3, -2, -3, 0, 4

3 moves between 0 and 4:
1, 2, -2, -3, 0, 3, 4

-2 moves between 4 and 1:
1, 2, -3, 0, 3, 4, -2

0 does not move:
1, 2, -3, 0, 3, 4, -2

4 moves between -3 and 0:
1, 2, -3, 4, 0, 3, -2
*/
