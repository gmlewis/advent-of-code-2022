// -*- compile-command: "go run main.go ../example1.txt ../example2.txt ../input.txt"; -*-

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

const numKnots = 10

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	puz := puzT{visited: map[keyT]bool{}}
	puz.makeMoves(lines)

	printf("Solution: %v\n", len(puz.visited))
}

type keyT [2]int

func (k keyT) x() int { return k[0] }
func (k keyT) y() int { return k[1] }

type puzT struct {
	knots   [numKnots]keyT
	visited map[keyT]bool
}

func (p *puzT) makeMoves(moves []string) {
	p.visited[p.knots[numKnots-1]] = true
	for _, move := range moves {
		parts := strings.Split(move, " ")
		d := must.Atoi(parts[1])
		var dx, dy int
		switch parts[0] {
		case "U":
			dy = -1
		case "D":
			dy = 1
		case "L":
			dx = -1
		case "R":
			dx = 1
		default:
			log.Fatalf("bad dir: %v", parts[0])
		}
		for i := 0; i < d; i++ {
			for knot := 0; knot < numKnots-1; knot++ {
				p.move(knot, dx, dy)
			}
			p.visited[p.knots[numKnots-1]] = true
			// log.Printf("after move %q - puz:\n%v", move, p)
		}
	}
}

func (p *puzT) move(knot, dx, dy int) {
	newHead := p.knots[knot]
	if knot == 0 {
		newHead = keyT{newHead.x() + dx, newHead.y() + dy}
		p.knots[0] = newHead
	}
	tail := &p.knots[knot+1]
	switch {
	// just move the head.
	case
		tail.x() == newHead.x() && tail.y() == newHead.y(),
		tail.x() == newHead.x()+1 && tail.y() == newHead.y(),
		tail.x() == newHead.x()-1 && tail.y() == newHead.y(),
		tail.x() == newHead.x() && tail.y() == newHead.y()+1,
		tail.x() == newHead.x() && tail.y() == newHead.y()-1,
		tail.x() == newHead.x()+1 && tail.y() == newHead.y()+1,
		tail.x() == newHead.x()-1 && tail.y() == newHead.y()-1,
		tail.x() == newHead.x()+1 && tail.y() == newHead.y()-1,
		tail.x() == newHead.x()-1 && tail.y() == newHead.y()+1:
	case
		tail.x() == newHead.x()-2 && tail.y() == newHead.y()+1,
		tail.x() == newHead.x()-2 && tail.y() == newHead.y(),
		tail.x() == newHead.x()-2 && tail.y() == newHead.y()-1:
		*tail = keyT{newHead.x() - 1, newHead.y()}
	case
		tail.x() == newHead.x()+2 && tail.y() == newHead.y()+1,
		tail.x() == newHead.x()+2 && tail.y() == newHead.y(),
		tail.x() == newHead.x()+2 && tail.y() == newHead.y()-1:
		*tail = keyT{newHead.x() + 1, newHead.y()}
	case
		tail.x() == newHead.x()+1 && tail.y() == newHead.y()-2,
		tail.x() == newHead.x() && tail.y() == newHead.y()-2,
		tail.x() == newHead.x()-1 && tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x(), newHead.y() - 1}
	case
		tail.x() == newHead.x()+1 && tail.y() == newHead.y()+2,
		tail.x() == newHead.x() && tail.y() == newHead.y()+2,
		tail.x() == newHead.x()-1 && tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x(), newHead.y() + 1}

	case
		tail.x() == newHead.x()-2 && tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x() - 1, newHead.y() - 1}
	case
		tail.x() == newHead.x()-2 && tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x() - 1, newHead.y() + 1}
	case
		tail.x() == newHead.x()+2 && tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x() + 1, newHead.y() - 1}
	case
		tail.x() == newHead.x()+2 && tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x() + 1, newHead.y() + 1}
	case
		tail.x() == newHead.x()-2 && tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x() - 1, newHead.y() - 1}
	case
		tail.x() == newHead.x()+2 && tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x() + 1, newHead.y() - 1}
	case
		tail.x() == newHead.x()-2 && tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x() - 1, newHead.y() + 1}
	case
		tail.x() == newHead.x()+2 && tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x() + 1, newHead.y() + 1}

	default:
		log.Fatalf("Unhandled case: H(%v) - T(%v)", newHead, tail)
	}
}

func (p *puzT) String() string {
	knot := func(x, y int) string {
		key := keyT{x, y}
		for i, k := range p.knots {
			if k != key {
				continue
			}
			if i == 0 {
				return "H"
			}
			return fmt.Sprintf("%v", i)
		}
		return "."
	}

	var lines []string
	bounds := p.findBounds()
	for y := bounds[0].y(); y <= bounds[1].y(); y++ {
		var line string
		for x := bounds[0].x(); x <= bounds[1].x(); x++ {
			line += knot(x, y)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (p *puzT) findBounds() [2]keyT {
	bounds := [2]keyT{}
	for _, key := range p.knots {
		if key.x() < bounds[0].x() {
			bounds[0][0] = key.x()
		}
		if key.y() < bounds[0].y() {
			bounds[0][1] = key.y()
		}
		if key.x() > bounds[1].x() {
			bounds[1][0] = key.x()
		}
		if key.y() > bounds[1].y() {
			bounds[1][1] = key.y()
		}
	}
	return bounds
}

/*
2022/12/09 13:50:21 after move "U 4" - puz:
....H
4321.
2022/12/09 13:50:21 after move "U 4" - puz:
....H
.4321
5....
2022/12/09 13:50:21 after move "U 4" - puz:
....H
....1
.432.
5....
2022/12/09 13:50:21 after move "U 4" - puz:
....H
....1
..432
.65..
7....

SHOULD BE:

....H
....1
..432
.5...
6....  (6 covers 7, 8, 9, s)

*/
