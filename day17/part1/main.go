// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

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
	buf := must.ReadFile(filename)

	puz := parsePuzzle(buf)
	for i := 0; i < numRocks; i++ {
		puz.dropRock(i)
		if i < 10 {
			fmt.Printf("%v\n\n", puz)
		}
	}

	printf("Solution: %v\n", puz.height)
}

const (
	startOffset  = 2
	chamberWidth = 7
	numRocks     = 2022
)

type keyT [2]int
type puzT struct {
	time   int
	height int
	gas    []rune
	grid   map[keyT]rune
}

func (p *puzT) dropRock(n int) {
	// rockNum := n % 5
}

func parsePuzzle(buf string) *puzT {
	return &puzT{
		gas:  []rune(strings.TrimSpace(buf)),
		grid: map[keyT]rune{},
	}
}

func (p *puzT) String() string {
	var lines []string
	for y := p.height - 1; y >= 0; y-- {
		line := "|"
		for x := 0; x < chamberWidth; x++ {
			r := p.grid[keyT{x, y}]
			if r == 0 {
				r = '.'
			}
			line += string(r)
		}
		line += "|"
		lines = append(lines, line)
	}
	lines = append(lines, "+-------+")
	return strings.Join(lines, "\n")
}
