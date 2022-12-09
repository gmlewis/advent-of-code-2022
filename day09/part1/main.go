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

	puz := puzT{visited: map[keyT]bool{}}
	puz.makeMoves(lines)

	printf("Solution: %v\n", len(puz.visited))
}

type keyT [2]int

func (k keyT) x() int { return k[0] }
func (k keyT) y() int { return k[1] }

type puzT struct {
	head    keyT
	tail    keyT
	visited map[keyT]bool
}

func (p *puzT) makeMoves(moves []string) {
	p.visited[p.tail] = true
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
			p.move(dx, dy)
			p.visited[p.tail] = true
		}
	}
}

func (p *puzT) move(dx, dy int) {
	newHead := keyT{p.head.x() + dx, p.head.y() + dy}
	switch {
	// just move the head.
	case
		p.tail.x() == newHead.x() && p.tail.y() == newHead.y(),
		p.tail.x() == newHead.x()+1 && p.tail.y() == newHead.y(),
		p.tail.x() == newHead.x()-1 && p.tail.y() == newHead.y(),
		p.tail.x() == newHead.x() && p.tail.y() == newHead.y()+1,
		p.tail.x() == newHead.x() && p.tail.y() == newHead.y()-1,
		p.tail.x() == newHead.x()+1 && p.tail.y() == newHead.y()+1,
		p.tail.x() == newHead.x()-1 && p.tail.y() == newHead.y()-1,
		p.tail.x() == newHead.x()+1 && p.tail.y() == newHead.y()-1,
		p.tail.x() == newHead.x()-1 && p.tail.y() == newHead.y()+1:
		p.head = newHead
	case
		p.tail.x() == newHead.x()-2 && p.tail.y() == newHead.y():
		p.head = newHead
		p.tail = keyT{newHead.x() - 1, newHead.y()}
	case
		p.tail.x() == newHead.x()+2 && p.tail.y() == newHead.y():
		p.head = newHead
		p.tail = keyT{newHead.x() + 1, newHead.y()}
	case
		p.tail.x() == newHead.x() && p.tail.y() == newHead.y()-2:
		p.head = newHead
		p.tail = keyT{newHead.x(), newHead.y() - 1}
	case
		p.tail.x() == newHead.x() && p.tail.y() == newHead.y()+2:
		p.head = newHead
		p.tail = keyT{newHead.x(), newHead.y() + 1}

	case
		p.tail.y() == newHead.y()+2:
		p.head = newHead
		p.tail = keyT{newHead.x(), newHead.y() + 1}
	case
		p.tail.y() == newHead.y()-2:
		p.head = newHead
		p.tail = keyT{newHead.x(), newHead.y() - 1}

	case
		p.tail.x() == newHead.x()+2:
		p.head = newHead
		p.tail = keyT{newHead.x() + 1, newHead.y()}
	case
		p.tail.x() == newHead.x()-2:
		p.head = newHead
		p.tail = keyT{newHead.x() - 1, newHead.y()}

	default:
		log.Fatalf("Unhandled case: H(%v) - T(%v)", newHead, p.tail)
	}
}
