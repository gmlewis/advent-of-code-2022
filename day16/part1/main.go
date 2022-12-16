// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/algorithm"
	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
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

	for puz.elapsedTime < maxTime {
		puz = puz.bestMove()
		log.Printf("puz=%#v", puz)
	}

	// maxKey, maxDistance := algorithm.Max(distances, math.MaxInt)
	// log.Printf("maxKey=%v, maxDistance=%v", maxKey, maxDistance)

	// solution := puz.maxPressure()

	printf("Solution: %v\n", puz.pressure(maxTime))
}

func (p *puzT) bestMove() *puzT {
	if p.elapsedTime >= maxTime {
		return p
	}

	name := p.moves[len(p.moves)-1]
	// valve := p.valves[name]

	distances := algorithm.Dijkstra[string, int](p, name, nil, math.MaxInt)
	log.Printf("distances from %q: %#v", name, distances)

	needToOpen := Filter(maps.Keys(p.valves), func(n string) bool { return p.valves[n].flowRate > 0 && p.valves[n].openTime >= maxTime })
	log.Printf("valves that still need to be opened: %+v", needToOpen)

	if len(needToOpen) == 0 {
		p.elapsedTime++
		return p
	}

	var bestValve string
	var bestPressure int
	for _, n := range needToOpen {
		pathTo := algorithm.PathTo[string, int](p, name, n, math.MaxInt)
		log.Printf("Path from %q to %q: %+v", name, n, pathTo)

		pressure := p.potentialPressureIfGoTo(name, n, pathTo)
		if pressure == 0 {
			continue
		}

		log.Printf("potential added pressure if %v were next: %v", n, pressure)
		if bestValve == "" || bestPressure < pressure {
			bestValve = n
			bestPressure = pressure
		}
	}
	log.Printf("bestValve=%q, bestPressure=%v", bestValve, bestPressure)

	if bestValve == "" {
		p.elapsedTime++
		return p
	}

	newP := p.moveTo(bestValve)
	newP.openValve(bestValve)

	return newP
}

func (p *puzT) potentialPressureIfGoTo(from, to string, pathTo map[string]int) int {
	minutes := pathTo[to] * 2

	// for k, v := range pathTo {
	// 	if k == from {
	// 		continue
	// 	}
	// }

	totalFlowTime := maxTime - 1 - minutes - p.elapsedTime
	if totalFlowTime <= 0 {
		return 0
	}
	pressure := totalFlowTime * p.valves[to].flowRate
	return pressure
}

const (
	maxTime = 30
)

type puzT struct {
	elapsedTime int
	moves       []string
	opened      map[string]bool
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
	if p.elapsedTime >= maxTime {
		return p
	}

	name := p.moves[len(p.moves)-1]
	valve := p.valves[name]

	if valve.flowRate > 0 && valve.openTime > p.elapsedTime { // open the valve
		valve.openTime = p.elapsedTime
		p.opened[name] = true
		// p.printSummary()
		// fmt.Printf("You open valve %v.\n", name)
		p.elapsedTime++
		if p.elapsedTime >= maxTime || len(p.opened) == len(p.valves) {
			p.elapsedTime = maxTime
			return p
		}
	}

	var bestPuz *puzT
	var bestPressure int
	for k := range valve.neighbors {
		newP := p.moveTo(k).maxPressure()
		pressure := newP.pressure(maxTime)
		if bestPuz == nil || pressure > bestPressure {
			bestPuz = newP
			bestPressure = pressure
		}
	}

	return bestPuz
}

func (p *puzT) openValve(name string) {
	valve := p.valves[name]
	valve.openTime = p.elapsedTime
	p.opened[name] = true
	// p.printSummary()
	// fmt.Printf("You open valve %v.\n", name)
	p.elapsedTime++
}

func (p *puzT) moveTo(name string) *puzT {
	// p.printSummary()
	// fmt.Printf("You move to valve %v.\n", name)

	moves := append([]string{}, p.moves...)
	newP := &puzT{
		elapsedTime: p.elapsedTime + 1,
		moves:       append(moves, name),
		opened:      map[string]bool{},
		valves:      map[string]*valveT{},
	}
	for k := range p.opened {
		newP.opened[k] = true
	}
	for k, v := range p.valves {
		newP.valves[k] = &valveT{
			openTime:  v.openTime,
			flowRate:  v.flowRate,
			neighbors: v.neighbors, // ok to point to original map, as neighbors does not change.
		}
	}
	return newP
}

func (p *puzT) printSummary() {
	fmt.Printf("\n== Minute %v ==\n", p.elapsedTime+1)
	open := maps.Keys(p.opened)
	open = Filter(open, func(name string) bool { return p.valves[name].flowRate > 0 })
	if len(open) > 0 {
		sort.Strings(open)
		pressure := p.pressure(p.elapsedTime + 1)
		fmt.Printf("Valves %v are open, releasing %v pressure.\n", strings.Join(open, ", "), pressure)
	} else {
		fmt.Printf("No valves are open.\n")
	}
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
		moves:  []string{"AA"},
		opened: map[string]bool{},
		valves: map[string]*valveT{},
	}

	for _, line := range lines {
		m := lineRE.FindStringSubmatch(line)
		if len(m) != 4 {
			log.Fatalf("Unable to parse line: %v", line)
		}
		name := m[1]
		v := &valveT{
			openTime:  maxTime, // not open yet
			flowRate:  must.Atoi(m[2]),
			neighbors: map[string]bool{},
		}
		for _, n := range strings.Split(m[3], ", ") {
			v.neighbors[n] = true
		}
		p.valves[name] = v
		if v.flowRate == 0 {
			p.opened[name] = true // no need to open it
		}
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
