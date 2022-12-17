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
type rockFunc func(p *puzT, startX int)
type rockDef []keyT

var (
	dropRockN = []rockFunc{dropRock1, dropRock2, dropRock3, dropRock4, dropRock5}
	rock1def  = []keyT{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	rock2def  = []keyT{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
	rock3def  = []keyT{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	rock4def  = []keyT{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	rock5def  = []keyT{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
)

type RenderRock func(p *puzT) bool

func (p *puzT) checkRock(x, y int, rockdef rockDef) bool {
	for _, k := range rockdef {
		if r := p.grid[keyT{x + k[0], y + k[1]}]; r != 0 {
			return false // collision
		}
	}
	for _, k := range rockdef {
		p.grid[keyT{x + k[0], y + k[1]}] = '#'
	}
	return true
}

func rock1(x, y int) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock1def) }
}

func rock2(x, y int) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock2def) }
}

func rock3(x, y int) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock3def) }
}

func rock4(x, y int) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock4def) }
}

func rock5(x, y int) RenderRock {
	return func(p *puzT) bool { return p.checkRock(x, y, rock5def) }
}

func (p *puzT) dropSpecificRock(x, width, height int, rockFunc func(x, y int) RenderRock) {
	if x > chamberWidth-width {
		x = chamberWidth - width
	}
	y := p.height

	noCollision := func() bool {
		return false
	}

	for y > 0 && noCollision() {
		y--
	}

	rockFunc(x, y)(p)

	p.height += height
}

func dropRock1(p *puzT, x int) {
	// ####
	p.dropSpecificRock(x, 4, 1, rock1)
}

func dropRock2(p *puzT, x int) {
	// .#.
	// ###
	// .#.
	p.dropSpecificRock(x, 3, 3, rock2)
}

func dropRock3(p *puzT, x int) {
	// ..#
	// ..#
	// ###
	p.dropSpecificRock(x, 3, 3, rock3)
}

func dropRock4(p *puzT, x int) {
	// #
	// #
	// #
	// #
	p.dropSpecificRock(x, 1, 4, rock4)
}

func dropRock5(p *puzT, x int) {
	// ##
	// ##
	p.dropSpecificRock(x, 2, 2, rock5)
}

func (p *puzT) dropRock(n int) {
	rockNum := n % 5
	dxFreefall := p.getGasDx(startYOffset)
	startX := startXOffset + dxFreefall
	if startX < 0 {
		startX = 0
	}
	dropRockN[rockNum](p, startX)
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
