package main

import (
	"log"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
	"github.com/google/go-cmp/cmp"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name string
		p1   string
		p2   string
		want int
	}{
		{
			name: "equal",
			p1:   "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			p2:   "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			want: 0,
		},
		{
			name: "== Pair 1 ==",
			p1:   "[1,1,3,1,1]",
			p2:   "[1,1,5,1,1]",
			// - Compare [1,1,3,1,1] vs [1,1,5,1,1]
			//   - Compare 1 vs 1
			//   - Compare 1 vs 1
			//   - Compare 3 vs 5
			//     - Left side is smaller, so inputs are in the right order
			want: -1,
		},
		{
			name: "== Pair 2 ==",
			p1:   "[[1],[2,3,4]]",
			p2:   "[[1],4]",
			// - Compare [[1],[2,3,4]] vs [[1],4]
			//   - Compare [1] vs [1]
			//     - Compare 1 vs 1
			//   - Compare [2,3,4] vs 4
			//     - Mixed types; convert right to [4] and retry comparison
			//     - Compare [2,3,4] vs [4]
			//       - Compare 2 vs 4
			//         - Left side is smaller, so inputs are in the right order
			want: -1,
		},
		{
			name: "== Pair 3 ==",
			p1:   "[9]",
			p2:   "[[8,7,6]]",
			// - Compare [9] vs [[8,7,6]]
			//   - Compare 9 vs [8,7,6]
			//     - Mixed types; convert left to [9] and retry comparison
			//     - Compare [9] vs [8,7,6]
			//       - Compare 9 vs 8
			//         - Right side is smaller, so inputs are not in the right order
			want: 1,
		},
		{
			name: "== Pair 4 ==",
			p1:   "[[4,4],4,4]",
			p2:   "[[4,4],4,4,4]",
			// - Compare [[4,4],4,4] vs [[4,4],4,4,4]
			//   - Compare [4,4] vs [4,4]
			//     - Compare 4 vs 4
			//     - Compare 4 vs 4
			//   - Compare 4 vs 4
			//   - Compare 4 vs 4
			//   - Left side ran out of items, so inputs are in the right order
			want: -1,
		},
		{
			name: "== Pair 5 ==",
			p1:   "[7,7,7,7]",
			p2:   "[7,7,7]",
			// - Compare [7,7,7,7] vs [7,7,7]
			//   - Compare 7 vs 7
			//   - Compare 7 vs 7
			//   - Compare 7 vs 7
			//   - Right side ran out of items, so inputs are not in the right order
			want: 1,
		},
		{
			name: "== Pair 6 ==",
			p1:   "[]",
			p2:   "[3]",
			// - Compare [] vs [3]
			//   - Left side ran out of items, so inputs are in the right order
			want: -1,
		},
		{
			name: "== Pair 7 ==",
			p1:   "[[[]]]",
			p2:   "[[]]",
			// - Compare [[[]]] vs [[]]
			//   - Compare [[]] vs []
			//     - Right side ran out of items, so inputs are not in the right order
			want: 1,
		},
		{
			name: "== Pair 8 ==",
			p1:   "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			p2:   "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			// - Compare [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
			//   - Compare 1 vs 1
			//   - Compare [2,[3,[4,[5,6,7]]]] vs [2,[3,[4,[5,6,0]]]]
			//     - Compare 2 vs 2
			//     - Compare [3,[4,[5,6,7]]] vs [3,[4,[5,6,0]]]
			//       - Compare 3 vs 3
			//       - Compare [4,[5,6,7]] vs [4,[5,6,0]]
			//         - Compare 4 vs 4
			//         - Compare [5,6,7] vs [5,6,0]
			//           - Compare 5 vs 5
			//           - Compare 6 vs 6
			//           - Compare 7 vs 0
			//             - Right side is smaller, so inputs are not in the right order
			//
			// 	}
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := parsePacket(tt.p1)
			p2 := parsePacket(tt.p2)
			got := compare(p1, p2)

			if got != tt.want {
				t.Errorf("compare = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePacket(t *testing.T) {
	tests := []struct {
		s    string
		want *packetT
	}{
		{
			s:    "[]",
			want: &packetT{},
		},
		{
			s: "[1,1,3,1,1]",
			want: &packetT{
				els: makeEls(1, 1, 3, 1, 1),
			},
		},
		{
			s: "[[1],[2,3,4]]",
			want: &packetT{
				els: []*packetT{
					{els: makeEls(1)},
					{els: makeEls(2, 3, 4)},
				},
			},
		},
		{
			s: "[[[]]]",
			want: &packetT{
				els: []*packetT{
					{els: []*packetT{
						{els: []*packetT{}},
					}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := parsePacket(tt.s)

			if gotStr := got.String(); gotStr != tt.s {
				t.Errorf("parsePacket.String = '%v', want '%v'", gotStr, tt.s)
			}

			if diff := cmp.Diff(tt.want, got, cmp.Comparer(packetsEqual)); diff != "" {
				t.Errorf("parsePacket('%v') mismatch (-want, +got):\n%v", tt.s, diff)
			}
		})
	}
}

func intPtr(v int) *int { return &v }
func makeEls(vs ...int) []*packetT {
	p := make([]*packetT, 0, len(vs))
	for _, v := range vs {
		p = append(p, &packetT{v: intPtr(v)})
	}
	return p
}

func packetsEqual(p1, p2 *packetT) bool {
	switch {
	case p1 == nil && p2 == nil:
		return true
	case p1 == nil || p2 == nil:
		return false
	case p1.v == nil && p2.v == nil:
		if len(p1.els) != len(p2.els) {
			return false
		}
		for i, el1 := range p1.els {
			el2 := p2.els[i]
			if !packetsEqual(el1, el2) {
				return false
			}
		}
		return true
	case p1.v == nil || p2.v == nil:
		return false
	case p1.v != nil && p2.v != nil:
		return *p1.v == *p2.v
	default:
		log.Fatalf("Unhandled case: packetsEqual:\np1=%#v\np2=%#v", p1, p2)
		return false
	}
}

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
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`
