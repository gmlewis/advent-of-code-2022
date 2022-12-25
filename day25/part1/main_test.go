package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestSanfu2Dec(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "1", want: 1},
		{s: "2", want: 2},
		{s: "1=", want: 3},
		{s: "1-", want: 4},
		{s: "10", want: 5},
		{s: "11", want: 6},
		{s: "12", want: 7},
		{s: "2=", want: 8},
		{s: "2-", want: 9},
		{s: "20", want: 10},
		{s: "21", want: 11},
		{s: "22", want: 12},
		{s: "1==", want: 13},
		{s: "1=-", want: 14},
		{s: "1=0", want: 15},
		{s: "1=1", want: 16},
		{s: "1=2", want: 17},
		{s: "1-=", want: 18},
		{s: "1--", want: 19},
		{s: "1-0", want: 20},
		{s: "1-1", want: 21},
		{s: "1-2", want: 22},
		{s: "10=", want: 23},
		{s: "10-", want: 24},
		{s: "100", want: 25},
		{s: "101", want: 26},
		{s: "102", want: 27},
		{s: "11=", want: 28},
		{s: "11-", want: 29},
		{s: "110", want: 30},
		{s: "111", want: 31},
		{s: "112", want: 32},
		{s: "12=", want: 33},
		{s: "12-", want: 34},
		{s: "120", want: 35},
		{s: "121", want: 36},
		{s: "122", want: 37},
		{s: "2==", want: 38},
		{s: "2=-", want: 39},
		{s: "2=0", want: 40},
		{s: "2=1", want: 41},
		{s: "2=2", want: 42},
		{s: "2-=", want: 43},
		{s: "2--", want: 44},
		{s: "2-0", want: 45},
		{s: "2-1", want: 46},
		{s: "2-2", want: 47},
		{s: "20=", want: 48},
		{s: "20-", want: 49},
		{s: "200", want: 50},
		{s: "201", want: 51},
		{s: "202", want: 52},
		{s: "21=", want: 53},
		{s: "21-", want: 54},
		{s: "210", want: 55},
		{s: "211", want: 56},
		{s: "212", want: 57},
		{s: "22=", want: 58},
		{s: "22-", want: 59},
		{s: "220", want: 60},
		{s: "221", want: 61},
		{s: "222", want: 62},
		{s: "1===", want: 63},
		{s: "1=11-2", want: 2022},
		{s: "1-0---0", want: 12345},
		{s: "1121-1110-1=0", want: 314159265},
		{s: "1=-0-2", want: 1747},
		{s: "12111", want: 906},
		{s: "2=0=", want: 198},
		{s: "21", want: 11},
		{s: "2=01", want: 201},
		{s: "111", want: 31},
		{s: "20012", want: 1257},
		{s: "112", want: 32},
		{s: "1=-1=", want: 353},
		{s: "1-12", want: 107},
		{s: "12", want: 7},
		{s: "1=", want: 3},
		{s: "122", want: 37},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := snafu2dec(tt.s)
			if got != tt.want {
				t.Errorf("snafu2dec('%v') = %v, want %v", tt.s, got, tt.want)
			}

			back := dec2snafu(tt.want)
			if back != tt.s {
				t.Errorf("dec2snafu(%v) = '%v', want '%v'", tt.want, back, tt.s)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 2=-1=0\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`
