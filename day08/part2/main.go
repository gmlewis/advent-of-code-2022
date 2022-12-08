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
	// log.Printf("max interior height: %v", puz.maxInteriorHeight)
	// log.Printf("puz:\n%v", puz)
	solution := puz.highestScenicScore()

	printf("Solution: %v\n", solution)
}

type keyT [2]int

func (k keyT) x() int { return k[0] }
func (k keyT) y() int { return k[1] }

type puzT struct {
	width    int
	height   int
	m        map[keyT]int
	interior map[keyT]int
}

func (p *puzT) highestScenicScore() int {
	var highestScore int
	for key, h := range p.interior {
		score := p.scenicScore(key, h)
		if score > highestScore {
			highestScore = score
		}
	}
	return highestScore
}

func (p *puzT) scenicScore(key keyT, h int) int {
	r := p.sibblingDist(key, h, 1, 0)
	t := p.sibblingDist(key, h, 0, -1)
	l := p.sibblingDist(key, h, -1, 0)
	d := p.sibblingDist(key, h, 0, 1)
	return r * t * l * d
}

func (p *puzT) sibblingDist(key keyT, h, dx, dy int) int {
	var dist int
	for x, y := key.x()+dx, key.y()+dy; x >= 0 && y >= 0 && x < p.width && y < p.height; x, y = x+dx, y+dy {
		dist++
		k := keyT{x, y}
		if p.m[k] >= h {
			break
		}
	}
	// log.Printf("sibblingDist(%v,%v,%v) = %v", key, dx, dy, dist)
	return dist
}

func (p *puzT) String() string {
	var lines []string
	for y := 0; y < p.height; y++ {
		var line string
		for x := 0; x < p.width; x++ {
			key := keyT{x, y}
			if v, ok := p.interior[key]; ok {
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
			visible++
		}
	}
	return visible
}

func parsePuzzle(lines []string) *puzT {
	puz := &puzT{
		width:    len(lines[0]),
		height:   len(lines),
		m:        map[keyT]int{},
		interior: map[keyT]int{},
	}

	for y, line := range lines {
		for x, r := range line {
			h := must.Atoi(string(r))
			key := keyT{x, y}
			puz.m[key] = h
			if x > 0 && y > 0 && x < puz.width-1 && y < puz.height-1 {
				puz.interior[key] = h
			}
		}
	}

	return puz
}
