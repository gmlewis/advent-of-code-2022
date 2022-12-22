// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
	parts := must.ReadSplitFile(filename, "\n\n")

	puz := parsePuzzle(strings.Split(parts[0], "\n"), parts[1])
	// fmt.Printf("BEFORE puzzle:\n%v\n", puz)
	p := puz.solve()
	// fmt.Printf("AFTER puzzle:\n%v\n", puz)

	printf("Solution: %v\n", p.password())
}

type posT struct {
	key keyT // x (col), y (row)
	dx  int
	dy  int
}
type keyT [2]int
type puzT struct {
	dirs string

	maxCol int
	maxRow int

	leftX  map[int]int
	rightX map[int]int
	topY   map[int]int
	botY   map[int]int
	grid   map[keyT]rune

	// debugging
	path map[keyT]rune
}

var digitsRE = regexp.MustCompile(`^(\d+)`)

func (p *puzT) solve() posT {
	pos := posT{key: keyT{p.leftX[1], 1}, dx: 1, dy: 0}

	dirs := p.dirs
	for {
		m := digitsRE.FindStringSubmatch(dirs)
		if len(m) != 2 {
			log.Fatalf("unable to parse: '%v'", dirs)
		}
		v := must.Atoi(m[1])
		pos = p.move(pos, v)
		// fmt.Printf("\nafter p.move(%v, %v):\n%v", pos, v, p)

		dirs = dirs[len(m[1]):]
		if len(dirs) == 0 {
			break
		}
		switch dirs[0] {
		case 'R':
			pos = posT{key: pos.key, dx: -pos.dy, dy: pos.dx}
			p.path[pos.key] = pos.symbol()
			// fmt.Printf("\nafter turn 'R':\n%v", p)
		case 'L':
			pos = posT{key: pos.key, dx: pos.dy, dy: -pos.dx}
			p.path[pos.key] = pos.symbol()
			// fmt.Printf("\nafter turn 'L':\n%v", p)
		default:
			log.Fatalf("expected turn, got: '%v'", dirs)
		}
		dirs = dirs[1:]
	}

	return pos
}

func (p *puzT) move(pos posT, steps int) posT {
	for i := 0; i < steps; i++ {
		x, y := pos.key[0], pos.key[1]
		next := keyT{x + pos.dx, y + pos.dy}
		switch {
		case p.grid[next] == '#':
			return pos
		case p.grid[next] == '.':
			p.path[pos.key] = pos.symbol()
			pos.key = next
			p.path[pos.key] = pos.symbol()

		case p.grid[next] == 0 && pos.dx > 0 && p.grid[keyT{p.leftX[y], y}] == '.': // right wrap available
			p.path[pos.key] = pos.symbol()
			pos.key = keyT{p.leftX[y], y}
			p.path[pos.key] = pos.symbol()
		case p.grid[next] == 0 && pos.dx > 0 && p.grid[keyT{p.leftX[y], y}] == '#': // right wrap blocked
			return pos

		case p.grid[next] == 0 && pos.dx < 0 && p.grid[keyT{p.rightX[y], y}] == '.': // left wrap available
			p.path[pos.key] = pos.symbol()
			pos.key = keyT{p.rightX[y], y}
			p.path[pos.key] = pos.symbol()
		case p.grid[next] == 0 && pos.dx < 0 && p.grid[keyT{p.rightX[y], y}] == '#': // left wrap blocked
			return pos

		case p.grid[next] == 0 && pos.dy > 0 && p.grid[keyT{x, p.topY[x]}] == '.': // down wrap available
			p.path[pos.key] = pos.symbol()
			pos.key = keyT{x, p.topY[x]}
			p.path[pos.key] = pos.symbol()
		case p.grid[next] == 0 && pos.dy > 0 && p.grid[keyT{x, p.topY[x]}] == '#': // down wrap blocked
			return pos

		case p.grid[next] == 0 && pos.dy < 0 && p.grid[keyT{x, p.botY[x]}] == '.': // up wrap available
			p.path[pos.key] = pos.symbol()
			pos.key = keyT{x, p.botY[x]}
			p.path[pos.key] = pos.symbol()
		case p.grid[next] == 0 && pos.dy < 0 && p.grid[keyT{x, p.botY[x]}] == '#': // up wrap available
			return pos
		default:
			log.Fatalf("unhandled case: move(%v, %v)", pos, steps)
		}
	}
	return pos
}

func (p posT) symbol() rune {
	switch {
	case p.dx > 0 && p.dy == 0:
		return '>'
	case p.dx < 0 && p.dy == 0:
		return '<'
	case p.dx == 0 && p.dy > 0:
		return 'v'
	case p.dx == 0 && p.dy < 0:
		return '^'
	default:
		log.Fatalf("bad pos: dx=%v, dy=%v", p.dx, p.dy)
	}
	return 0
}

func (p *puzT) String() string {
	lines := []string{fmt.Sprintf("max: {%v,%v}", p.maxCol, p.maxRow)}

	for y := 1; y <= p.maxRow; y++ {
		var line string
		for x := 1; x <= p.maxCol; x++ {
			if r := p.path[keyT{x, y}]; r != 0 {
				line += string(r)
			} else {
				line += " "
			}
		}
		lines = append(lines, fmt.Sprintf("%v - leftX[%v]=%v, rightX[%v]=%v", line, y, p.leftX[y], y, p.rightX[y]))
	}

	for x := 1; x <= p.maxCol; x++ {
		lines = append(lines, fmt.Sprintf("topY[%v]=%v, botY[%v]=%v", x, p.topY[x], x, p.botY[x]))
	}

	return strings.Join(lines, "\n")
}

func parsePuzzle(lines []string, dirs string) *puzT {
	puz := &puzT{
		dirs:   dirs,
		leftX:  map[int]int{},
		rightX: map[int]int{},
		topY:   map[int]int{},
		botY:   map[int]int{},
		grid:   map[keyT]rune{},
		path:   map[keyT]rune{}, // debugging
	}

	for y, line := range lines {
		row := y + 1
		var validMap bool
		for x, r := range line {
			col := x + 1
			switch {
			case r == ' ':
			default:
				key := keyT{col, row}
				puz.grid[key] = r
				puz.path[key] = r

				puz.rightX[row] = col
				puz.botY[col] = row
				if puz.topY[col] == 0 {
					puz.topY[col] = row
				}
				if !validMap {
					validMap = true
					puz.leftX[row] = col
				}
				if col > puz.maxCol {
					puz.maxCol = col
				}
				if row > puz.maxRow {
					puz.maxRow = row
				}
			}
		}
	}

	return puz
}

func (p posT) password() int {
	var facing int
	switch {
	case p.dx > 0 && p.dy == 0: // right, facing=0
	case p.dx == 0 && p.dy > 0: // down
		facing = 1
	case p.dx < 0 && p.dy == 0: // left
		facing = 2
	case p.dx == 0 && p.dy < 0: // up
		facing = 3
	}
	return 1000*p.key[1] + 4*p.key[0] + facing
}
