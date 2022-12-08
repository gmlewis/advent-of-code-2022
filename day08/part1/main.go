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
	lines := must.ReadFileLines(filename)

	puz := parsePuzzle(lines)
	puz.visibleTop()
	puz.visibleRight()
	puz.visibleBottom()
	puz.visibleLeft()
	// log.Printf("puz:\n%v", puz)

	printf("Solution: %v\n", len(puz.visible))
}

type keyT [2]int
type puzT struct {
	width   int
	height  int
	m       map[keyT]int
	visible map[keyT]int
}

func (p *puzT) String() string {
	var lines []string
	for y := 0; y < p.height; y++ {
		var line string
		for x := 0; x < p.width; x++ {
			key := keyT{x, y}
			if v, ok := p.visible[key]; ok {
				line += fmt.Sprintf("%v", v)
			} else {
				line += " "
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (p *puzT) visibleFrom(x, y, dx, dy int) int {
	var visible int
	for lastHeight := -1; x >= 0 && y >= 0 && x < p.width && y < p.height; x, y = x+dx, y+dy {
		key := keyT{x, y}
		if p.m[key] > lastHeight {
			// log.Printf("%v is visible", key)
			lastHeight = p.m[key]
			p.visible[key] = lastHeight
			visible++
		}
	}
	return visible
}

func (p *puzT) visibleTop() int {
	var visible int
	for x := 0; x < p.width; x++ {
		visible += p.visibleFrom(x, 0, 0, 1)
	}
	return visible
}

func (p *puzT) visibleRight() int {
	var visible int
	for y := 0; y < p.height; y++ {
		visible += p.visibleFrom(p.width-1, y, -1, 0)
	}
	return visible
}

func (p *puzT) visibleBottom() int {
	var visible int
	for x := 0; x < p.width; x++ {
		visible += p.visibleFrom(x, p.height-1, 0, -1)
	}
	return visible
}

func (p *puzT) visibleLeft() int {
	var visible int
	for y := 0; y < p.height; y++ {
		visible += p.visibleFrom(0, y, 1, 0)
	}
	return visible
}

func parsePuzzle(lines []string) *puzT {
	puz := &puzT{
		width:   len(lines[0]),
		height:  len(lines),
		m:       map[keyT]int{},
		visible: map[keyT]int{},
	}

	for y, line := range lines {
		for x, r := range line {
			h := must.Atoi(string(r))
			key := keyT{x, y}
			puz.m[key] = h
		}
	}

	return puz
}
