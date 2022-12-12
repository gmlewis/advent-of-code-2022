// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

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
	log.Printf("start=%v, end=%v", puz.start, puz.end)

	printf("Solution: %v\n", len(puz.grid))
}

// puzT implements the algorithm.Graph interface.
var _ algorithm.Graph[keyT, int] = &puzT{}

type keyT [2]int
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
		case 'E':
			p.end = key
		}
	}
	return acc
}

func (p *puzT) Distance(from, to keyT) int {
	return 0
}

func (p *puzT) Each(func(key keyT)) {
}

func (p *puzT) EachNeighbor(from keyT, f func(from, to keyT)) {
}
