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
	"github.com/gmlewis/advent-of-code-2021/stream"
)

var logf = log.Printf
var printf = fmt.Printf

var (
	debug = flag.Bool("debug", false, "Print debugging")
)

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	puz := parsePuzzle(lines)
	puz = puz.maxPressure()

	printf("Solution: %v\n", puz.pressure(maxTime))
}

func (p *puzT) maxPressure() *puzT {
	var bestPuz *puzT
	var bestPressure int
	for order := range p.allPossibleOrders() {
		// log.Printf("Order: %v", strings.Join(order, ", "))
		puz, pressure := p.checkPressure(bestPressure, order)
		// log.Printf("Order: %v - Final pressure %v", strings.Join(order, ", "), pressure)
		if bestPuz == nil || pressure > bestPressure {
			// log.Printf("New winning order: %v - %v", pressure, strings.Join(order, ", "))
			bestPuz = puz
			bestPressure = pressure
		}
	}
	return bestPuz
}

func (p *puzT) checkPressure(bestPressure int, order []string) (*puzT, int) {
	newP := p.copy()
	for _, name := range order {
		newP.makeMoveTo(name)
	}

	// for newP.elapsedTime <= maxTime {
	// 	newP.printSummary()
	// 	newP.elapsedTime++
	// }

	return newP, newP.pressure(maxTime)
}

func (p *puzT) allPossibleOrders() <-chan []string {
	keys := maps.Keys(p.needsToOpen)
	sort.Slice(keys, func(a, b int) bool { return p.valves[keys[a]].flowRate > p.valves[keys[b]].flowRate })
	// log.Printf("keys: %v", keys)
	return stream.PermutationsOf(keys)
}

// func (p *puzT) maxPressure() *puzT {
// 	if p.elapsedTime >= maxTime || len(p.needsToOpen) == 0 {
// 		p.elapsedTime++
// 		return p
// 	}
//
// 	var bestPressure int
// 	var bestPuz *puzT
// 	for n := range p.needsToOpen {
// 		newP := p.makeMoveTo(n).maxPressure()
// 		pressure := newP.pressure(maxTime)
// 		if bestPuz == nil || pressure > bestPressure {
// 			bestPuz = newP
// 			bestPressure = pressure
// 		}
// 	}
//
// 	return bestPuz
// }

// func (p *puzT) bestMove() *puzT {
// 	if p.elapsedTime >= maxTime || len(p.needsToOpen) == 0 {
// 		p.elapsedTime++
// 		p.printSummary()
// 		return p
// 	}
//
// 	name := p.moves[len(p.moves)-1]
//
// 	distances := algorithm.Dijkstra[string, int](p, name, nil, math.MaxInt)
// 	log.Printf("distances from %q: %#v", name, distances)
//
// 	log.Printf("valves that still need to be opened: %+v", maps.Keys(p.needsToOpen))
//
// 	var bestValve string
// 	var bestPressure int
// 	var bestPathTo map[string]int
// 	for n := range p.needsToOpen {
// 		pathTo := algorithm.PathTo[string, int](p, name, n, math.MaxInt)
// 		log.Printf("Path from %q to %q: %+v", name, n, pathTo)
//
// 		pressure := p.potentialPressureIfGoTo(name, n, pathTo)
// 		if pressure == 0 {
// 			continue
// 		}
//
// 		if bestValve == "" || bestPressure < pressure {
// 			bestValve = n
// 			bestPressure = pressure
// 			bestPathTo = pathTo
// 		}
// 	}
// 	log.Printf("bestValve=%q, bestPressure=%v\n\n", bestValve, bestPressure)
//
// 	if bestValve == "" {
// 		p.elapsedTime++
// 		return p
// 	}
//
// 	return p.followPathTo(bestPathTo)
// }

func (p *puzT) makeMoveTo(to string) {
	name := p.moves[len(p.moves)-1]
	pathTo := algorithm.PathTo[string, int](p, name, to, math.MaxInt)
	// log.Printf("makeMoveTo: from=%q, to=%q, pathTo=%#v", name, to, pathTo)
	p.followPathTo(pathTo)
}

func (p *puzT) followPathTo(bestPathTo map[string]int) {
	if len(bestPathTo) == 0 {
		return
	}

	path := maps.Keys(bestPathTo)
	sort.Slice(path, func(a, b int) bool { return bestPathTo[path[a]] < bestPathTo[path[b]] })

	for _, n := range path[1:] {
		p.moveTo(n)
		// Don't automatically open a valve just because you are there already.
	}
	p.openValve(path[len(path)-1])
}

func (p *puzT) potentialPressureIfGoTo(from, to string, pathTo map[string]int) int {
	minutes := pathTo[to] // this includes opening the valve

	// for k, v := range pathTo {
	// 	if k == from {
	// 		continue
	// 	}
	// }

	totalFlowTime := maxTime - minutes - p.elapsedTime
	if totalFlowTime <= 0 {
		return 0
	}
	pressure := totalFlowTime * p.valves[to].flowRate
	// log.Printf("potential added pressure if %v were next: %v (totalFlowTime=%v, flowRate=%v)", to, pressure, totalFlowTime, p.valves[to].flowRate)
	return pressure
}

const (
	maxTime = 30
)

type puzT struct {
	elapsedTime int
	moves       []string
	needsToOpen map[string]bool
	valves      map[string]*valveT
}

// puzT implements the algorithm.Graph interface.
var _ algorithm.Graph[string, int] = &puzT{}

type valveT struct {
	openTime  int
	flowRate  int
	neighbors map[string]bool
}

// func (p *puzT) maxPressure() *puzT {
// 	if p.elapsedTime >= maxTime {
// 		return p
// 	}
//
// 	name := p.moves[len(p.moves)-1]
// 	valve := p.valves[name]
//
// 	if valve.flowRate > 0 && valve.openTime > p.elapsedTime { // open the valve
// 		p.openValve(name)
// 		if p.elapsedTime >= maxTime || len(p.needsToOpen) == 0 {
// 			p.elapsedTime = maxTime
// 			return p
// 		}
// 	}
//
// 	var bestPuz *puzT
// 	var bestPressure int
// 	for k := range valve.neighbors {
// 		newP := p.moveTo(k).maxPressure()
// 		pressure := newP.pressure(maxTime)
// 		if bestPuz == nil || pressure > bestPressure {
// 			bestPuz = newP
// 			bestPressure = pressure
// 		}
// 	}
//
// 	return bestPuz
// }

func (p *puzT) openValve(name string) {
	p.elapsedTime++
	if *debug {
		p.printSummary()
		fmt.Printf("You open valve %v.\n\n", name)
	}

	valve := p.valves[name]
	valve.openTime = p.elapsedTime
	delete(p.needsToOpen, name)
}

func (p *puzT) copy() *puzT {
	moves := append([]string{}, p.moves...)
	newP := &puzT{
		elapsedTime: p.elapsedTime,
		moves:       moves,
		needsToOpen: map[string]bool{},
		valves:      map[string]*valveT{},
	}
	for k := range p.needsToOpen {
		newP.needsToOpen[k] = true
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

func (p *puzT) moveTo(name string) {
	p.moves = append(p.moves, name)
	p.elapsedTime++
}

func (p *puzT) printSummary() {
	fmt.Printf("\n== Minute %v ==\n", p.elapsedTime)
	open := Filter(maps.Keys(p.valves), func(name string) bool { return p.valves[name].flowRate > 0 && !p.needsToOpen[name] })
	if len(open) > 0 {
		sort.Strings(open)
		pressure := Reduce(open, 0, func(n string, acc int) int { return acc + p.valves[n].flowRate })
		totalPressure := p.pressure(p.elapsedTime)
		fmt.Printf("Valves %v are open, releasing %v pressure (%v total).\n", strings.Join(open, ", "), pressure, totalPressure)
	} else {
		fmt.Printf("No valves are open.\n")
	}

	fmt.Printf("Moves: %v\n", strings.Join(p.moves, ", "))
	stillClosed := maps.Keys(p.needsToOpen)
	if len(stillClosed) > 0 {
		fmt.Printf("Valves %v are still closed.\n", strings.Join(stillClosed, ", "))
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
		moves:       []string{"AA"},
		needsToOpen: map[string]bool{},
		valves:      map[string]*valveT{},
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
		if v.flowRate > 0 {
			p.needsToOpen[name] = true
		}
	}

	return p
}

func (p *puzT) Distance(from, to string) int {
	if !p.valves[from].neighbors[to] {
		return math.MaxInt
	}
	return 1
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
