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

		case p.grid[next] == 0 && p.wrappedSpace(pos) == '.': // available
			p.path[pos.key] = pos.symbol()
			pos = p.wrappedPos(pos)
			p.path[pos.key] = pos.symbol()

		case p.grid[next] == 0 && p.wrappedSpace(pos) == '#': // blocked
			return pos

		default:
			log.Fatalf("unhandled case: move(%v, %v)", pos, steps)
		}
	}
	return pos
}

func (p *puzT) wrappedPos(pos posT) posT {
	x, y := pos.key[0], pos.key[1]

	// TODO: make this work for any puzzle input - maybe?!?  maybe not.
	if p.maxCol < 20 { // example1.txt - K=4
		switch {
		case y < 5 && pos.dx < 0: // face 1 to face 3 - left to down
			return posT{dx: 0, dy: 1, key: keyT{4 + y, 5}}
		case y < 5 && pos.dx > 0: // face 1 to face 6 - right to left
			return posT{dx: -1, dy: 0, key: keyT{p.maxCol, 13 - y}}
		case y < 5: // face 1 to face 2 - up to down
			return posT{dx: 0, dy: 1, key: keyT{5 - (x - 8), 5}}
		case y > 8 && x > 12 && pos.dy < 0: // face 6 to face 4 - up to left
			return posT{dx: -1, dy: 0, key: keyT{12, 9 - (x - 12)}}
		case y > 8 && x > 12 && pos.dy > 0: // face 6 to face 2 - down to right
			return posT{dx: 1, dy: 0, key: keyT{1, 9 - (x - 12)}}
		case y > 8 && x > 12: // face 6 to face 1 - right to left
			return posT{dx: -1, dy: 0, key: keyT{12, 5 - (y - 8)}}
		case y > 8 && pos.dx < 0: // face 5 to face 3 - left to up
			return posT{dx: 0, dy: -1, key: keyT{9 - (y - 8), 8}}
		case y > 8: // face 5 to face 2 - down to up
			// log.Printf("face 5 to face 2 - down to up - (%v,%v) => (%v,8)", x, y, 5-(x-8))
			return posT{dx: 0, dy: -1, key: keyT{5 - (x - 8), 8}}
		case x < 5 && pos.dx < 0: // face 2 to face 6 - left to up
			return posT{dx: 0, dy: -1, key: keyT{13 - (y - 4), 12}}
		case x < 5 && pos.dy < 0: // face 2 to face 1 - up to down
			return posT{dx: 0, dy: 1, key: keyT{13 - x, 1}}
		case x < 5: // face 2 to face 5 - down to up
			return posT{dx: 0, dy: -1, key: keyT{13 - x, 12}}
		case x > 8: // face 4 to face 6 - right to down
			// log.Printf("face 4 to face 6 - right to down - (%v,%v) => (%v,9)", x, y, 17-(y-4))
			return posT{dx: 0, dy: 1, key: keyT{17 - (y - 4), 9}}
		case pos.dy < 0: // face 3 to face 1 - up to right
			// log.Printf("face 3 to face 1 - up to right - (%v,%v) => (8,%v)", x, y, 8+(x-4))
			return posT{dx: 1, dy: 0, key: keyT{9, (x - 4)}}
		default: // face 3 to face 5 - down to right
			return posT{dx: 1, dy: 0, key: keyT{9, 13 - (x - 4)}}
		}
	}

	// input.txt - K=50
	//  16
	//  4
	// 35
	// 2
	switch {
	case y > 150 && pos.dx > 0: // face 2 to face 5 - right to up
		return posT{dx: 0, dy: -1, key: keyT{50 + (y - 150), 150}}
	case y > 150 && pos.dx < 0: // face 2 to face 1 - left to down
		return posT{dx: 0, dy: 1, key: keyT{50 + (y - 150), 1}}
	case y > 150: // face 2 to face 6 - down to down
		return posT{dx: 0, dy: 1, key: keyT{100 + x, 1}}
	case y > 100 && pos.dx > 0: // face 5 to face 6 - right to left
		return posT{dx: -1, dy: 0, key: keyT{150, 51 - (y - 100)}}
	case y > 100 && pos.dx < 0: // face 3 to face 1 - left to right
		return posT{dx: 1, dy: 0, key: keyT{51, 51 - (y - 100)}}
	case y > 100 && pos.dy > 0: // face 5 to face 2 - down to left
		return posT{dx: -1, dy: 0, key: keyT{50, 100 + x}}
	case y > 100: // face 3 to face 4 - up to right
		return posT{dx: 1, dy: 0, key: keyT{51, 50 + x}}
	case y > 50 && pos.dx > 0: // face 4 to face 6 - right to up
		return posT{dx: 0, dy: -1, key: keyT{100 + (y - 50), 50}}
	case y > 50: // face 4 to face 3 - left to down
		return posT{dx: 0, dy: 1, key: keyT{y - 50, 101}}
	case pos.dx > 0: // face 6 to face 5 - right to left
		return posT{dx: -1, dy: 0, key: keyT{100, 151 - y}}
	case pos.dx < 0: // face 1 to face 3 - left to right
		return posT{dx: 1, dy: 0, key: keyT{1, 151 - y}}
	case x > 100 && pos.dy > 0: // face 6 to face 4 - down to left
		return posT{dx: -1, dy: 0, key: keyT{100, x - 50}}
	case x > 100 && pos.dy < 0: // face 6 to face 2 - up to up
		return posT{dx: 0, dy: -1, key: keyT{x - 100, 200}}
	case pos.dx < 0: // face 1 to face 3 - left to right
		return posT{dx: 1, dy: 0, key: keyT{1, 151 - y}}
	default: // face 1 to face 2 - up to right
		return posT{dx: 1, dy: 0, key: keyT{1, x + 100}}
	}

	log.Fatalf("unhandled: pos=%v", pos)
	return posT{}
}

func (p *puzT) wrappedSpace(pos posT) rune {
	pos = p.wrappedPos(pos)
	r := p.grid[pos.key]
	if r != '.' && r != '#' {
		log.Fatalf("programming error: wrappedSpace(%v) = %v (%c)", pos, r, r)
	}
	return r
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
