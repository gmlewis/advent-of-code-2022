// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
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
	solution := puz.monkey("root")

	printf("Solution: %v\n", solution)
}

type monkeyT struct {
	value *int
	f     func(p *puzT) int
}
type puzT struct {
	monkies map[string]*monkeyT
}

func (p *puzT) monkey(name string) int {
	m := p.monkies[name]
	if m.value != nil {
		return *m.value
	}
	v := m.f(p)
	m.value = &v
	return v
}

func parsePuzzle(lines []string) *puzT {
	puz := &puzT{monkies: map[string]*monkeyT{}}

	for _, line := range lines {
		name, m := puz.parseLine(line)
		puz.monkies[name] = m
	}

	return puz
}

func (p *puzT) parseLine(line string) (name string, monkey *monkeyT) {
	parts := strings.Split(line, ": ")
	if v, err := strconv.Atoi(parts[1]); err == nil {
		return parts[0], &monkeyT{value: &v}
	}

	eqparts := strings.Split(parts[1], " ")
	switch eqparts[1] {
	case "+":
		return parts[0], &monkeyT{f: func(puz *puzT) int { return puz.monkey(eqparts[0]) + puz.monkey(eqparts[2]) }}
	case "-":
		return parts[0], &monkeyT{f: func(puz *puzT) int { return puz.monkey(eqparts[0]) - puz.monkey(eqparts[2]) }}
	case "*":
		return parts[0], &monkeyT{f: func(puz *puzT) int { return puz.monkey(eqparts[0]) * puz.monkey(eqparts[2]) }}
	case "/":
		return parts[0], &monkeyT{f: func(puz *puzT) int { return puz.monkey(eqparts[0]) / puz.monkey(eqparts[2]) }}
	default:
		log.Fatalf("unhandled line: %v", line)
	}
	return "", nil
}
