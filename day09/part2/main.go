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

	// 2600 is too low.

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
		tail.x() == newHead.x()-2 && tail.y() == newHead.y():
		*tail = keyT{newHead.x() - 1, newHead.y()}
	case
		tail.x() == newHead.x()+2 && tail.y() == newHead.y():
		*tail = keyT{newHead.x() + 1, newHead.y()}
	case
		tail.x() == newHead.x() && tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x(), newHead.y() - 1}
	case
		tail.x() == newHead.x() && tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x(), newHead.y() + 1}

	case
		tail.y() == newHead.y()+2:
		*tail = keyT{newHead.x(), newHead.y() + 1}
	case
		tail.y() == newHead.y()-2:
		*tail = keyT{newHead.x(), newHead.y() - 1}

	case
		tail.x() == newHead.x()+2:
		*tail = keyT{newHead.x() + 1, newHead.y()}
	case
		tail.x() == newHead.x()-2:
		*tail = keyT{newHead.x() - 1, newHead.y()}

	// case // ???
	// 	tail.x() == newHead.x()+3:
	// 	tail = keyT{newHead.x() + 2, newHead.y()}
	// case
	// 	tail.x() == newHead.x()-3:
	// 	tail = keyT{newHead.x() - 2, newHead.y()}
	//
	// case // ???
	// 	tail.x() == newHead.x()+4:
	// 	tail = keyT{newHead.x() + 3, newHead.y()}
	// case
	// 	tail.x() == newHead.x()-4:
	// 	tail = keyT{newHead.x() - 3, newHead.y()}

	default:
		log.Fatalf("Unhandled case: H(%v) - T(%v)", newHead, tail)
	}
}
