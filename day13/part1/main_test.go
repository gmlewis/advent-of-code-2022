package main

import (
	"log"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
	"github.com/google/go-cmp/cmp"
)

func TestParsePacket(t *testing.T) {
	tests := []struct {
		s    string
		want *packetT
	}{
		{
			s:    "",
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
					{els: []*packetT{}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := parsePacket(tt.s)
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
	want := "Solution: 36\n"
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
