// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/algorithm"
	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var logf = log.Printf
var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	puz := parsePuzzle(lines)
	// distances := algorithm.Dijkstra[string, int](puz, "AA", nil, math.MaxInt)
	// log.Printf("distances: %#v", distances)
	// maxKey, maxDistance := algorithm.Max(distances, math.MaxInt)
	// log.Printf("maxKey=%v, maxDistance=%v", maxKey, maxDistance)

	solution := puz.maxPressure()

	printf("Solution: %v\n", solution.pressure(30))
}

type puzT struct {
	elapsedTime int
	moves       []string
	visited     map[string]bool
	valves      map[string]*valveT
}

// puzT implements the algorithm.Graph interface.
var _ algorithm.Graph[string, int] = &puzT{}

type valveT struct {
	openTime  int
	flowRate  int
	neighbors map[string]bool
}

func (p *puzT) maxPressure() *puzT {
	return p
}

func (p *puzT) pressure(elapsedTime int) int {
	var result int
	for _, v := range p.valves {
		ot := elapsedTime - v.openTime
		if ot <= 0 {
			continue
		}
		result += ot * v.flowRate
	}
	return result
}

var lineRE = regexp.MustCompile(`^Valve (\S+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)

func parsePuzzle(lines []string) *puzT {
	p := &puzT{
		moves:   []string{"AA"},
		visited: map[string]bool{},
		valves:  map[string]*valveT{},
	}

	for _, line := range lines {
		m := lineRE.FindStringSubmatch(line)
		if len(m) != 4 {
			log.Fatalf("Unable to parse line: %v", line)
		}
		v := &valveT{
			openTime:  30, // not open yet
			flowRate:  must.Atoi(m[2]),
			neighbors: map[string]bool{},
		}
		for _, n := range strings.Split(m[3], ", ") {
			v.neighbors[n] = true
		}
		p.valves[m[1]] = v
	}

	return p
}

func (p *puzT) Distance(from, to string) int {
	if p.valves[from].neighbors[to] {
		return 2
	}
	return math.MaxInt
}

func (p *puzT) Each(f func(key string)) {
	for k := range p.valves {
		f(k)
	}
}

func (p *puzT) EachNeighbor(from string, f func(from, to string)) {
	for k := range p.valves[from].neighbors {
		f(from, k)
	}
}
