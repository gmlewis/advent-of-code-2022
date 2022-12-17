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

var (
	dropRockN = []rockFunc{
		dropRock1,
		dropRock2,
		dropRock3,
		dropRock4,
		dropRock5,
	}
)

type RenderRock func(p *puzT)

func rock1(x, y int) RenderRock {
	return func(p *puzT) {
		p.grid[keyT{x, y}] = '#'
		p.grid[keyT{x + 1, y}] = '#'
		p.grid[keyT{x + 2, y}] = '#'
		p.grid[keyT{x + 3, y}] = '#'
	}
}

func rock2(x, y int) RenderRock {
	return func(p *puzT) {
		p.grid[keyT{x + 1, y}] = '#'
		p.grid[keyT{x, y + 1}] = '#'
		p.grid[keyT{x + 1, y + 1}] = '#'
		p.grid[keyT{x + 2, y + 1}] = '#'
		p.grid[keyT{x + 1, y + 2}] = '#'
	}
}

func rock3(x, y int) RenderRock {
	return func(p *puzT) {
		p.grid[keyT{x, y}] = '#'
		p.grid[keyT{x + 1, y}] = '#'
		p.grid[keyT{x + 2, y}] = '#'
		p.grid[keyT{x + 2, y + 1}] = '#'
		p.grid[keyT{x + 2, y + 2}] = '#'
	}
}

func rock4(x, y int) RenderRock {
	return func(p *puzT) {
		p.grid[keyT{x, y}] = '#'
		p.grid[keyT{x, y + 1}] = '#'
		p.grid[keyT{x, y + 2}] = '#'
		p.grid[keyT{x, y + 3}] = '#'
	}
}

func rock5(x, y int) RenderRock {
	return func(p *puzT) {
		p.grid[keyT{x, y}] = '#'
		p.grid[keyT{x, y + 1}] = '#'
		p.grid[keyT{x + 1, y}] = '#'
		p.grid[keyT{x + 1, y + 1}] = '#'
	}
}

func dropRock1(p *puzT, x int) {
	// ####
	if x > chamberWidth-4 {
		x = chamberWidth - 4
	}
	y := p.height

	noCollision := func() bool {
		return false
	}

	for y > 0 && noCollision() {
		y--
	}

	rock1(x, y)(p)

	p.height++
}

func dropRock2(p *puzT, x int) {
	// .#.
	// ###
	// .#.
	if x > chamberWidth-3 {
		x = chamberWidth - 3
	}
	y := p.height

	noCollision := func() bool {
		return false
	}

	for y > 0 && noCollision() {
		y--
	}

	rock2(x, y)(p)

	p.height += 3
}

func dropRock3(p *puzT, x int) {
	// ..#
	// ..#
	// ###
	if x > chamberWidth-3 {
		x = chamberWidth - 3
	}
	y := p.height

	noCollision := func() bool {
		return false
	}

	for y > 0 && noCollision() {
		y--
	}

	rock3(x, y)(p)

	p.height += 3
}

func dropRock4(p *puzT, x int) {
	// #
	// #
	// #
	// #
	if x > chamberWidth-1 {
		x = chamberWidth - 1
	}
	y := p.height

	noCollision := func() bool {
		return false
	}

	for y > 0 && noCollision() {
		y--
	}

	rock4(x, y)(p)

	p.height += 4
}

func dropRock5(p *puzT, x int) {
	// ##
	// ##
	if x > chamberWidth-2 {
		x = chamberWidth - 2
	}
	y := p.height

	noCollision := func() bool {
		return false
	}

	for y > 0 && noCollision() {
		y--
	}

	rock5(x, y)(p)

	p.height += 2
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
