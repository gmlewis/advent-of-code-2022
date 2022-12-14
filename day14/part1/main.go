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

const xStart = 500

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	puz := parseLines(lines)
	grains := puz.dropGrains()

	printf("Solution: %v\n", grains)
}

type keyT [2]int

func (k keyT) x() int { return k[0] }
func (k keyT) y() int { return k[1] }

type puzT struct {
	xmin int
	xmax int
	ymax int
	grid map[keyT]rune
}

func (p *puzT) dropGrains() int {
	var grains int
	for p.dropOneGrain() {
		grains++
	}
	return grains
}

func (p *puzT) dropOneGrain() bool {
	key := keyT{xStart, 0}
	dKey := func() keyT { return keyT{key.x(), key.y() + 1} }
	dlKey := func() keyT { return keyT{key.x() - 1, key.y() + 1} }
	drKey := func() keyT { return keyT{key.x() + 1, key.y() + 1} }
	down := func() rune { return p.grid[dKey()] }
	downLeft := func() rune { return p.grid[dlKey()] }
	downRight := func() rune { return p.grid[drKey()] }

	for {
		d := down()
		dl := downLeft()
		dr := downRight()
		switch {
		case key.x() < p.xmin, key.x() > p.xmax, key.y() > p.ymax:
			return false
		case d == 0:
			key = dKey()
		case dl != 0 && dr != 0:
			p.grid[key] = 'o'
			return true
		case dl == 0:
			key = dlKey()
		case dl != 0 && dr == 0:
			key = drKey()
		default:
			log.Fatalf("unhandled case: key=%v, d=%v, dl=%v, dr=%v", key, d, dl, dr)
		}
	}
}

func parseLines(lines []string) *puzT {
	p := &puzT{xmin: xStart, xmax: xStart, grid: map[keyT]rune{}}

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		p.drawRocks(parts)
	}

	return p
}

func (p *puzT) drawRocks(rocks []string) {
	var lastKey keyT
	for i, rock := range rocks {
		parts := strings.Split(rock, ",")
		x := must.Atoi(parts[0])
		y := must.Atoi(parts[1])
		if x < p.xmin {
			p.xmin = x
		}
		if x > p.xmax {
			p.xmax = x
		}
		if y > p.ymax {
			p.ymax = y
		}
		key := keyT{x, y}
		if i == 0 {
			lastKey = key
			continue
		}
		p.drawLine(lastKey, key, '#')
		lastKey = key
	}
}

func cmp(v int) int {
	switch {
	case v < 0:
		return -1
	case v > 0:
		return 1
	default:
		return 0
	}
}

func (p *puzT) drawLine(from, to keyT, r rune) {
	dx := cmp(to.x() - from.x())
	dy := cmp(to.y() - from.y())

	p.grid[to] = r
	for from != to {
		p.grid[from] = r
		from[0] += dx
		from[1] += dy
	}
}
