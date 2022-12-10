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

const initialX = 1

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	puz := parseProgram(lines)
	puz.run(220)

	printf("Solution: %v\n", puz.acc)
}

type puzT struct {
	pc     int // program counter
	ops    []*opcodeT
	cycles int
	x      int
	acc    int
}

type opcodeT struct {
	name   string
	cycles int
	value  int
	f      func(*puzT, *opcodeT)
}

func (p *puzT) run(cycles int) {
	for p.cycles <= cycles {
		p.clockTick(p.ops[p.pc])
		if p.ops[p.pc].f != nil {
			p.ops[p.pc].f(p, p.ops[p.pc])
		}
		p.pc = (p.pc + 1) % len(p.ops)
	}
}

func parseProgram(lines []string) *puzT {
	p := &puzT{x: initialX, ops: make([]*opcodeT, 0, len(lines))}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		var v int
		if len(parts) > 1 {
			v = must.Atoi(parts[1])
		}
		switch parts[0] {
		case "noop":
			p.ops = append(p.ops, &opcodeT{name: "noop", cycles: 1})
		case "addx":
			p.ops = append(p.ops, &opcodeT{name: "addx", cycles: 2, f: addx, value: v})
		}
	}
	return p
}

func (p *puzT) clockTick(op *opcodeT) {
	for i := 0; i < op.cycles; i++ {
		p.cycles++
		if (p.cycles+20)%40 == 0 {
			sig := p.cycles * p.x
			p.acc += sig
			// log.Printf("cycle %v: x=%v, signal strength=%v, acc=%v", p.cycles, p.x, sig, p.acc)
		}
	}
}

func addx(p *puzT, op *opcodeT) { p.x += op.value }
