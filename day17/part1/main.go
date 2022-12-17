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
	for i := 0; i < totalRocks; i++ {
		puz.dropRock(i)
		if i < 10 {
			fmt.Printf("%v\n\n", puz)
		}
	}

	printf("Solution: %v\n", puz.height)
}

const (
	startXOffset = 2
	startYOffset = 4
	chamberWidth = 7
	totalRocks   = 2022
)

type keyT [2]int
type puzT struct {
	height   int
	gasIndex int
	gas      []rune
	grid     map[keyT]rune
}

// type rockFunc func(p *puzT, startX int)
type RenderRock func(p *puzT) bool
type rockFunc func(x, y int, render bool) RenderRock
type rockDef []keyT

var (
	rockFuncs = []rockFunc{rock1, rock2, rock3, rock4, rock5}
	rock1def  = []keyT{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	rock2def  = []keyT{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
	rock3def  = []keyT{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	rock4def  = []keyT{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	rock5def  = []keyT{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	rockSizes = []keyT{{4, 1}, {3, 3}, {3, 3}, {1, 4}, {2, 2}}
)

func (p *puzT) checkRock(x, y int, rockdef rockDef, render bool) bool {
	for _, k := range rockdef {
		if r := p.grid[keyT{x + k[0], y + k[1]}]; r != 0 {
			return false // collision
		}
	}
	if render {
		for _, k := range rockdef {
			p.grid[keyT{x + k[0], y + k[1]}] = '#'
		}
	}
	return true
}

func rock1(x, y int, render bool) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock1def, render) }
}

func rock2(x, y int, render bool) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock2def, render) }
}

func rock3(x, y int, render bool) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock3def, render) }
}

func rock4(x, y int, render bool) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock4def, render) }
}

func rock5(x, y int, render bool) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock5def, render) }
}

func (p *puzT) dropSpecificRock(x int, rockSize keyT, fn rockFunc) {
	if x > chamberWidth-rockSize[0] {
		x = chamberWidth - rockSize[0]
	}
	y := p.height + 1

	for y > 0 && fn(x, y-1, false)(p) {
		y--
		dx := p.getGasDx(1)
		if x+dx >= 0 && x+dx+rockSize[0] < chamberWidth && fn(x+dx, y, false)(p) {
			x += dx
		}
	}

	fn(x, y, true)(p)

	newHeight := y + rockSize[1]
	if newHeight > p.height {
		p.height = newHeight
	}
}

func (p *puzT) dropRock(n int) {
	rockNum := n % 5
	dxFreefall := p.getGasDx(startYOffset - 1)
	startX := startXOffset + dxFreefall
	if startX < 0 {
		startX = 0
	}

	p.dropSpecificRock(startX, rockSizes[rockNum], rockFuncs[rockNum])
}

func (p *puzT) getGasDx(n int) int {
	var dx int
	for i := 0; i < n; i++ {
		r := p.gas[p.gasIndex]
		p.gasIndex = (p.gasIndex + 1) % len(p.gas)
		switch r {
		case '>':
			dx++
		case '<':
			dx--
		default:
			log.Fatalf("Bad gas: '%c'", r)
		}
	}
	return dx
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
