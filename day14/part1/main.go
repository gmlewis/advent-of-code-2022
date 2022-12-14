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

	printf("Solution: %v\n", len(puz.grid))
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

func (p *puzT) drawLine(from, to keyT, r rune) {
}
