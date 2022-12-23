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

const totalRounds = 10

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	puz := ReduceWithIndex(lines, &puzT{grid: map[keyT]rune{}}, parseLine)
	// log.Printf("puzzle:\n%v", puz)

	for i := 0; i < totalRounds; i++ {
		var ok bool
		if puz, ok = puz.round(i); !ok {
			break
		}
	}

	printf("Solution: %v\n", puz.spaces())
}

type keyT [2]int
type puzT struct {
	minX int
	maxX int
	minY int
	maxY int
	grid map[keyT]rune
}

func (p *puzT) round(roundNum int) (*puzT, bool) {
	newP := &puzT{minX: p.maxX, minY: p.maxY, grid: map[keyT]rune{}}

	// first pass
	wishMove := map[keyT]keyT{}      // map[from]to
	possibleMoves := map[keyT]keyT{} // map[to]from
	for from := range p.grid {
		to, ok := p.elfMove(from, roundNum)
		if !ok {
			continue
		}
		if f1, ok := possibleMoves[to]; ok {
			delete(wishMove, f1)
			continue
		}
		possibleMoves[to] = from
		wishMove[from] = to
	}

	if len(wishMove) == 0 {
		return p, false
	}

	// second pass
	for from := range p.grid {
		to, ok := wishMove[from]
		if !ok {
			to = from
		}

		newP.grid[to] = '#'
		if to[0] < newP.minX {
			newP.minX = to[0]
		}
		if to[0] > newP.maxX {
			newP.maxX = to[0]
		}
		if to[1] < newP.minY {
			newP.minY = to[1]
		}
		if to[1] > newP.maxY {
			newP.maxY = to[1]
		}
	}

	// log.Printf("newP:\n%v", newP)
	return newP, true
}

var (
	rounds = [][]keyT{
		{keyT{-1, -1}, keyT{0, -1}, keyT{1, -1}}, // NW, N, NE
		{keyT{-1, 1}, keyT{0, 1}, keyT{1, 1}},    // SW, S, SE
		{keyT{-1, -1}, keyT{-1, 0}, keyT{-1, 1}}, // NW, W, SW
		{keyT{1, -1}, keyT{1, 0}, keyT{1, 1}},    // NE, E, SE
	}
)

func (p *puzT) elfMove(from keyT, roundNum int) (keyT, bool) {
	var shouldMove bool
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if p.grid[keyT{from[0] + dx, from[1] + dy}] == '#' {
				shouldMove = true
				break
			}
		}
	}

	if !shouldMove {
		return keyT{}, false
	}

	for i := 0; i < 4; i++ {
		round := rounds[(i+roundNum)%4]
		to := keyT{from[0] + round[1][0], from[1] + round[1][1]}
		if p.grid[keyT{from[0] + round[0][0], from[1] + round[0][1]}] != '#' &&
			p.grid[to] != '#' &&
			p.grid[keyT{from[0] + round[2][0], from[1] + round[2][1]}] != '#' {
			return to, true
		}
	}

	return keyT{}, false
}

func parseLine(y int, line string, acc *puzT) *puzT {
	parseRune := func(x int, r rune, p *puzT) *puzT {
		if r == '#' {
			p.grid[keyT{x, y}] = r
			if x > p.maxX {
				p.maxX = x
			}
			if y > p.maxY {
				p.maxY = y
			}
		}
		return p
	}
	return ReduceWithIndex([]rune(line), acc, parseRune)
}

func (p *puzT) spaces() int {
	var spaces int
	for y := p.minY; y <= p.maxY; y++ {
		for x := p.minX; x <= p.maxX; x++ {
			if p.grid[keyT{x, y}] != '#' {
				spaces++
			}
		}
	}
	return spaces
}

func (p *puzT) String() string {
	var lines []string
	for y := p.minY; y <= p.maxY; y++ {
		var line string
		for x := p.minX; x <= p.maxX; x++ {
			if p.grid[keyT{x, y}] == '#' {
				line += "#"
			} else {
				line += "."
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
