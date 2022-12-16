package main

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestPressure_Example1(t *testing.T) {
	p := parsePuzzle(strings.Split(strings.TrimSpace(example1), "\n"))

	// Enter the provided solution:
	p.valves["DD"].openTime = 2
	p.valves["BB"].openTime = 5
	p.valves["JJ"].openTime = 9
	p.valves["HH"].openTime = 17
	p.valves["EE"].openTime = 21
	p.valves["CC"].openTime = 24

	// Move sequence:
	// AA, DD, CC, BB, AA, II, JJ, II, AA, DD, EE, FF, GG, HH, GG, FF, EE, DD, CC

	if got, want := p.pressure(30), 1651; got != want {
		t.Errorf("pressure = %v, want %v", got, want)
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
