package main

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/algorithm"
	"github.com/gmlewis/advent-of-code-2021/test"
	"github.com/google/go-cmp/cmp"
)

func TestPathTo(t *testing.T) {
	p := parsePuzzle(strings.Split(strings.TrimSpace(example1), "\n"))

	tests := []struct {
		from string
		to   string
		want map[string]int
	}{
		{
			from: "DD",
			to:   "BB",
			want: map[string]int{
				"BB": 2,
				"CC": 1,
				"DD": 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v->%v", tt.from, tt.to), func(t *testing.T) {
			got := algorithm.PathTo[string, int](p, tt.from, tt.to, math.MaxInt)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PathTo mismatch (-want +got):\n%v", diff)
			}
		})
	}
}

func TestPressure_Example1(t *testing.T) {
	p := parsePuzzle(strings.Split(strings.TrimSpace(example1), "\n"))

	// Enter the provided solution:
	p.valves["DD"].openTime = 2  // rate=20
	p.valves["BB"].openTime = 5  // rate=13
	p.valves["JJ"].openTime = 9  // rate=21
	p.valves["HH"].openTime = 17 // rate=22
	p.valves["EE"].openTime = 21 // rate=3
	p.valves["CC"].openTime = 24 // rate=2

	// Move sequence:
	//   AA, DD, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC
	// Open valve sequence:
	//   DD, BB, JJ, HH, EE, CC

	if got, want := p.pressure(30), 1651; got != want {
		t.Errorf("pressure = %v, want %v", got, want)
	}

	want := map[string]*valveT{
		"AA": {flowRate: 0, neighbors: map[string]bool{"DD": true, "II": true, "BB": true}},
		"BB": {flowRate: 13, neighbors: map[string]bool{"CC": true, "AA": true}},
		"CC": {flowRate: 2, neighbors: map[string]bool{"DD": true, "BB": true}},
		"DD": {flowRate: 20, neighbors: map[string]bool{"CC": true, "AA": true, "EE": true}},
		"EE": {flowRate: 3, neighbors: map[string]bool{"FF": true, "DD": true}},
		"FF": {flowRate: 0, neighbors: map[string]bool{"EE": true, "GG": true}},
		"GG": {flowRate: 0, neighbors: map[string]bool{"FF": true, "HH": true}},
		"HH": {flowRate: 22, neighbors: map[string]bool{"GG": true}},
		"II": {flowRate: 0, neighbors: map[string]bool{"AA": true, "JJ": true}},
		"JJ": {flowRate: 21, neighbors: map[string]bool{"II": true}},
	}

	for k, v := range want {
		if v.flowRate != p.valves[k].flowRate {
			t.Errorf("valves[%q].flowRate = %v, want %v", k, p.valves[k].flowRate, v.flowRate)
		}

		if diff := cmp.Diff(v.neighbors, p.valves[k].neighbors); diff != "" {
			t.Errorf("valves mismatch (-want +got):\n%v", diff)
		}
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 1651\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`

/*
2022/12/16 18:14:53 Order: DD, BB, JJ, HH, EE, CC
2022/12/16 18:14:53 makeMoveTo: from="AA", to="DD", pathTo=map[string]int{"AA":0, "DD":2}

== Minute 1 ==
No valves are open.
Moves: AA, DD
Valves JJ, BB, CC, DD, EE, HH are still closed.
You move to valve DD.


== Minute 2 ==
No valves are open.
Moves: AA, DD
Valves CC, DD, EE, HH, JJ, BB are still closed.
You open valve DD.

2022/12/16 18:14:53 makeMoveTo: from="DD", to="BB", pathTo=map[string]int{"AA":1, "BB":3, "CC":2, "DD":0}

== Minute 3 ==
Valves DD are open, releasing 20 pressure (20 total).
Moves: AA, DD, AA
Valves CC, EE, HH, JJ, BB are still closed.
You move to valve AA.


== Minute 4 ==
Valves DD are open, releasing 20 pressure (40 total).
Moves: AA, DD, AA, CC
Valves HH, JJ, BB, CC, EE are still closed.
You move to valve CC.


== Minute 5 ==
Valves DD are open, releasing 20 pressure (60 total).
Moves: AA, DD, AA, CC, BB
Valves HH, JJ, BB, CC, EE are still closed.
You move to valve BB.


== Minute 6 ==
Valves DD are open, releasing 20 pressure (80 total).
Moves: AA, DD, AA, CC, BB
Valves HH, JJ, BB, CC, EE are still closed.
You open valve BB.

2022/12/16 18:14:53 makeMoveTo: from="BB", to="JJ", pathTo=map[string]int{"AA":1, "BB":0, "II":2, "JJ":4}

== Minute 7 ==
Valves BB, DD are open, releasing 33 pressure (113 total).
Moves: AA, DD, AA, CC, BB, AA
Valves CC, EE, HH, JJ are still closed.
You move to valve AA.


== Minute 8 ==
Valves BB, DD are open, releasing 33 pressure (146 total).
Moves: AA, DD, AA, CC, BB, AA, II
Valves JJ, CC, EE, HH are still closed.
You move to valve II.


== Minute 9 ==
Valves BB, DD are open, releasing 33 pressure (179 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ
Valves JJ, CC, EE, HH are still closed.
You move to valve JJ.


== Minute 10 ==
Valves BB, DD are open, releasing 33 pressure (212 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ
Valves JJ, CC, EE, HH are still closed.
You open valve JJ.

2022/12/16 18:14:53 makeMoveTo: from="JJ", to="HH", pathTo=map[string]int{"AA":2, "DD":3, "EE":5, "FF":6, "GG":7, "HH":9, "II":1, "JJ":0}

== Minute 11 ==
Valves BB, DD, JJ are open, releasing 54 pressure (266 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II
Valves CC, EE, HH are still closed.
You move to valve II.


== Minute 12 ==
Valves BB, DD, JJ are open, releasing 54 pressure (320 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA
Valves HH, CC, EE are still closed.
You move to valve AA.


== Minute 13 ==
Valves BB, DD, JJ are open, releasing 54 pressure (374 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD
Valves HH, CC, EE are still closed.
You move to valve DD.


== Minute 14 ==
Valves BB, DD, JJ are open, releasing 54 pressure (428 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE
Valves HH, CC, EE are still closed.
You move to valve EE.


== Minute 15 ==
Valves BB, DD, JJ are open, releasing 54 pressure (482 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF
Valves EE, HH, CC are still closed.
You move to valve FF.


== Minute 16 ==
Valves BB, DD, JJ are open, releasing 54 pressure (536 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG
Valves HH, CC, EE are still closed.
You move to valve GG.


== Minute 17 ==
Valves BB, DD, JJ are open, releasing 54 pressure (590 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH
Valves HH, CC, EE are still closed.
You move to valve HH.


== Minute 18 ==
Valves BB, DD, JJ are open, releasing 54 pressure (644 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH
Valves HH, CC, EE are still closed.
You open valve HH.

2022/12/16 18:14:53 makeMoveTo: from="HH", to="EE", pathTo=map[string]int{"EE":4, "FF":2, "GG":1, "HH":0}

== Minute 19 ==
Valves BB, DD, HH, JJ are open, releasing 76 pressure (720 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG
Valves CC, EE are still closed.
You move to valve GG.


== Minute 20 ==
Valves BB, DD, HH, JJ are open, releasing 76 pressure (796 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF
Valves CC, EE are still closed.
You move to valve FF.


== Minute 21 ==
Valves BB, DD, HH, JJ are open, releasing 76 pressure (872 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE
Valves CC, EE are still closed.
You move to valve EE.


== Minute 22 ==
Valves BB, DD, HH, JJ are open, releasing 76 pressure (948 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE
Valves EE, CC are still closed.
You open valve EE.

2022/12/16 18:14:53 makeMoveTo: from="EE", to="CC", pathTo=map[string]int{"CC":3, "DD":1, "EE":0}

== Minute 23 ==
Valves BB, DD, EE, HH, JJ are open, releasing 79 pressure (1027 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD
Valves CC are still closed.
You move to valve DD.


== Minute 24 ==
Valves BB, DD, EE, HH, JJ are open, releasing 79 pressure (1106 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC
Valves CC are still closed.
You move to valve CC.


== Minute 25 ==
Valves BB, DD, EE, HH, JJ are open, releasing 79 pressure (1185 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC
Valves CC are still closed.
You open valve CC.


== Minute 25 ==
Valves BB, CC, DD, EE, HH, JJ are open, releasing 81 pressure (1185 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC

== Minute 26 ==
Valves BB, CC, DD, EE, HH, JJ are open, releasing 81 pressure (1266 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC

== Minute 27 ==
Valves BB, CC, DD, EE, HH, JJ are open, releasing 81 pressure (1347 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC

== Minute 28 ==
Valves BB, CC, DD, EE, HH, JJ are open, releasing 81 pressure (1428 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC

== Minute 29 ==
Valves BB, CC, DD, EE, HH, JJ are open, releasing 81 pressure (1509 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC

== Minute 30 ==
Valves BB, CC, DD, EE, HH, JJ are open, releasing 81 pressure (1590 total).
Moves: AA, DD, AA, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC
2022/12/16 18:14:53 Order: DD, BB, JJ, HH, EE, CC - Final pressure 1590
*/
