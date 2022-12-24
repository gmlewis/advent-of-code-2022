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
	moves1 := puz.bestPathFromTo(keyT{0, -1}, keyT{puz.width - 1, puz.height})
	moves2 := puz.bestPathFromTo(keyT{puz.width - 1, puz.height}, keyT{0, -1})
	moves3 := puz.bestPathFromTo(keyT{0, -1}, keyT{puz.width - 1, puz.height})

	printf("Solution: %v\n", moves1+moves2+moves3)
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
	minutes int
	width   int
	height  int
	grid    map[keyT]blizT
}
type allPathsT [][]keyT

func (p *puzT) bestPathFromTo(from, to keyT) int {
	allPaths := allPathsT{[]keyT{from}}
	for {
		p.minutes++
		p.blowWindNextStep()

		var newPaths allPathsT
		alreadySuggested := map[keyT]bool{}
		for _, path := range allPaths {
			pos := path[0]
			options := p.allPossibleMovesFrom(pos, to)
			// log.Printf("Minute %v - pos=%v - %v options for path #%v: %+v", p.minutes, pos, len(options), i+1, options)
			for _, option := range options {
				if alreadySuggested[option] {
					continue
				}
				alreadySuggested[option] = true
				newPath := append([]keyT{option}, path...)
				if option == to {
					// log.Printf("FOUND SOLUTION after %v minutes!", p.minutes)
					return len(newPath) - 1
				}
				newPaths = append(newPaths, newPath)
			}
		}
		allPaths = newPaths
	}

	return 0
}

func (p *puzT) allPossibleMovesFrom(pos, goal keyT) []keyT {
	if pos == goal { // already solved - no options
		return nil
	}
	if pos[0] == goal[0] && (pos[1] == goal[1]-1 || pos[1] == goal[1]+1) { // at exit - one option
		return []keyT{goal}
	}

	var options []keyT
	if p.grid[pos] == 0 { // stay put
		options = append(options, pos)
	}

	f := func(k keyT) {
		if p.grid[k] == 0 && k[1] >= 0 && k[1] < p.height && k[0] >= 0 && k[0] < p.width {
			options = append(options, k)
		}
	}

	f(keyT{pos[0] + 1, pos[1]})
	f(keyT{pos[0], pos[1] + 1})
	f(keyT{pos[0] - 1, pos[1]})
	f(keyT{pos[0], pos[1] - 1})
	return options
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
	p := &puzT{grid: map[keyT]blizT{}}
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
