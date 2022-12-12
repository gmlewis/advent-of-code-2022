// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"

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

	puz := parseLines(lines)
	// log.Printf("start=%v, end=%v", puz.start, puz.end)
	distances := algorithm.Dijkstra[keyT, int](puz, puz.start, &puz.end, math.MaxInt32)

	printf("Solution: %v\n", distances[puz.end])
}

// puzT implements the algorithm.Graph interface.
var _ algorithm.Graph[keyT, int] = &puzT{}

type keyT [2]int

func (k keyT) x() int { return k[0] }
func (k keyT) y() int { return k[1] }

type puzT struct {
	start keyT
	end   keyT
	grid  map[keyT]rune
}

func parseLines(lines []string) *puzT {
	puz := &puzT{}
	puz.grid = ReduceWithIndex(lines, map[keyT]rune{}, puz.parseLine)
	return puz
}

func (p *puzT) parseLine(y int, line string, acc map[keyT]rune) map[keyT]rune {
	for x, r := range line {
		key := keyT{x, y}
		acc[key] = r
		switch r {
		case 'S':
			p.start = key
			acc[key] = 'a'
		case 'E':
			p.end = key
			acc[key] = 'z'
		}
	}
	return acc
}

func (p *puzT) Distance(from, to keyT) int {
	limit := 1 + p.grid[from]
	if v, ok := p.grid[to]; !ok || v > limit {
		return math.MaxInt32
	}

	dx := to.x() - from.x()
	if dx < 0 {
		dx = -dx
	}
	dy := to.y() - from.y()
	if dy < 0 {
		dy = -dy
	}
	// log.Printf("Distance(from=%v=%c, to=%v=%c): %v", from, p.grid[from], to, p.grid[to], dx+dy)
	return dx + dy
}

func (p *puzT) Each(f func(key keyT)) {
	for k := range p.grid {
		f(k)
	}
}

func (p *puzT) EachNeighbor(from keyT, f func(from, to keyT)) {
	// log.Printf("EachNeighbor(from=%v=%c)", from, p.grid[from])
	up := keyT{from.x(), from.y() - 1}
	if _, ok := p.grid[up]; ok {
		f(from, up)
	}
	down := keyT{from.x(), from.y() + 1}
	if _, ok := p.grid[down]; ok {
		f(from, down)
	}
	left := keyT{from.x() - 1, from.y()}
	if _, ok := p.grid[left]; ok {
		f(from, left)
	}
	right := keyT{from.x() + 1, from.y()}
	if _, ok := p.grid[right]; ok {
		f(from, right)
	}
}
