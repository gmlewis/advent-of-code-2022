// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math/bits"
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

	printf("Solution: %v\n", len(puz.grid))
}

const (
	blizUp blizT = 1 << iota
	blizRight
	blizDown
	blizLeft
)

type keyT [2]int
type blizT byte
type puzT struct {
	width  int
	height int
	pos    keyT
	grid   map[keyT]blizT
}

func (p *puzT) blowWindNextStep() {
	g := map[keyT]blizT{}
	for k, v := range p.grid {
		if v&blizUp == blizUp {
			key := keyT{k[0], (k[1] - 1 + p.height) % p.height}
			g[key] |= blizUp
		}
		if v&blizRight == blizRight {
			key := keyT{(k[0] + 1) % p.width, k[1]}
			g[key] |= blizRight
		}
		if v&blizDown == blizDown {
			key := keyT{k[0], (k[1] + 1) % p.height}
			g[key] |= blizDown
		}
		if v&blizLeft == blizLeft {
			key := keyT{(k[0] - 1 + p.width) % p.width, k[1]}
			g[key] |= blizLeft
		}
	}
	p.grid = g
}

func parsePuzzle(lines []string) *puzT {
	p := &puzT{pos: keyT{0, -1}, grid: map[keyT]blizT{}}
	ReduceWithIndex(lines, p, parseLine)
	return p
}

func parseLine(y int, line string, p *puzT) *puzT {
	if line[2] == '#' { // skip line
		p.width = len(line) - 2
		return p
	}
	p.height = y

	parseRune := func(x int, r rune, p *puzT) *puzT {
		if r == '#' {
			return p
		}
		k := keyT{x - 1, y - 1}
		switch r {
		case '.':
		case '^':
			p.grid[k] = blizUp
		case '>':
			p.grid[k] = blizRight
		case 'v':
			p.grid[k] = blizDown
		case '<':
			p.grid[k] = blizLeft
		default:
			log.Fatalf("unknown rune %v", r)
		}
		return p
	}

	ReduceWithIndex([]rune(line), p, parseRune)
	return p
}

func (p *puzT) String() string {
	lines := []string{fmt.Sprintf("#.%v", strings.Repeat("#", p.width-1))}
	for y := 0; y < p.height; y++ {
		var line string
		for x := 0; x < p.width; x++ {
			k := keyT{x, y}
			v := bits.OnesCount(uint(p.grid[k]))
			switch {
			case v == 0 && p.pos == k:
				line += "E"
			case v == 0:
				line += "."
			case v == 2:
				line += "2"
			case v == 3:
				line += "3"
			case v == 4:
				line += "4"
			case v == 1 && p.grid[k] == blizUp:
				line += "^"
			case v == 1 && p.grid[k] == blizRight:
				line += ">"
			case v == 1 && p.grid[k] == blizDown:
				line += "v"
			case v == 1 && p.grid[k] == blizLeft:
				line += "<"
			default:
				log.Printf("unhandled v=%v, k=%v, grid[k]=%v", v, k, p.grid[k])
			}
		}
		lines = append(lines, line)
	}
	lines = append(lines, fmt.Sprintf("%v.#", strings.Repeat("#", p.width-1)))
	return strings.Join(lines, "#\n#")
}
