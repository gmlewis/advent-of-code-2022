// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
	_, expr := puz.monkey("root")
	// log.Printf("expr=%v", expr)

	// first equation is a special case, want lhs&rhs to be identical.
	m := line1RE.FindStringSubmatch(expr)
	if len(m) != 4 || m[2] != "+" {
		log.Fatalf("unable to parse original expression: %v", expr)
	}
	rhs := must.Atoi(m[3])

	solution := solve(m[1], rhs)

	printf("Solution: %v\n", solution)
}

var line1RE = regexp.MustCompile(`^\((\S+)\)(\+|-|\*|/)\((\d+)\)$`)
var line2RE = regexp.MustCompile(`^\((\d+)\)(\+|-|\*|/)\((\S+)\)$`)

func solve(expr string, rhs int) int {
	if expr == "(humn)" {
		return rhs
	}

	m := line1RE.FindStringSubmatch(expr)
	if len(m) != 4 {
		m = line2RE.FindStringSubmatch(expr)
		if len(m) != 4 {
			log.Fatalf("unable to parse expr: %v", expr)
		}
	}

	v1, err1 := strconv.Atoi(m[1])
	v2, err2 := strconv.Atoi(m[3])
	op := m[2]
	// log.Printf("m[1]=%v op=%v, m[3]=%v, rhs=%v", m[1], op, m[3], rhs)

	switch {
	case op == "+" && err2 == nil:
		return solve(m[1], rhs-v2)
	case op == "+" && err1 == nil:
		return solve(m[3], rhs-v1)
	case op == "-" && err2 == nil:
		return solve(m[1], rhs+v2)
	case op == "-" && err1 == nil:
		return solve(m[3], v1-rhs)
	case op == "*" && err2 == nil:
		return solve(m[1], rhs/v2)
	case op == "*" && err1 == nil:
		return solve(m[3], rhs/v1)
	case op == "/" && err2 == nil:
		return solve(m[1], rhs*v2)
	default:
		log.Fatalf("unknown op %q in line %v", m[2], expr)
	}

	return 0
}

type monkeyT struct {
	value *int
	f     func(p *puzT) (int, string)
}
type puzT struct {
	monkies map[string]*monkeyT
}

func (p *puzT) monkey(name string) (int, string) {
	if name == "humn" {
		return 0, "(humn)"
	}

	m := p.monkies[name]
	if m.value != nil {
		return *m.value, ""
	}
	v, expr := m.f(p)
	if expr != "" {
		return 0, expr
	}
	m.value = &v
	return v, expr
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

	expr := func(v1, v2, result int, e1, op, e2 string) (int, string) {
		switch {
		case e1 == "" && e2 == "":
			return result, ""
		case e1 != "" && e2 != "":
			return 0, fmt.Sprintf("(%v)%v(%v)", e1, op, e2)
		case e1 != "":
			return 0, fmt.Sprintf("(%v)%v(%v)", e1, op, v2)
		default:
			return 0, fmt.Sprintf("(%v)%v(%v)", v1, op, e2)
		}
	}

	eqparts := strings.Split(parts[1], " ")
	switch eqparts[1] {
	case "+":
		return parts[0], &monkeyT{f: func(puz *puzT) (int, string) {
			v1, e1 := puz.monkey(eqparts[0])
			v2, e2 := puz.monkey(eqparts[2])
			return expr(v1, v2, v1+v2, e1, "+", e2)
		}}
	case "-":
		return parts[0], &monkeyT{f: func(puz *puzT) (int, string) {
			v1, e1 := puz.monkey(eqparts[0])
			v2, e2 := puz.monkey(eqparts[2])
			return expr(v1, v2, v1-v2, e1, "-", e2)
		}}
	case "*":
		return parts[0], &monkeyT{f: func(puz *puzT) (int, string) {
			v1, e1 := puz.monkey(eqparts[0])
			v2, e2 := puz.monkey(eqparts[2])
			return expr(v1, v2, v1*v2, e1, "*", e2)
		}}
	case "/":
		return parts[0], &monkeyT{f: func(puz *puzT) (int, string) {
			v1, e1 := puz.monkey(eqparts[0])
			v2, e2 := puz.monkey(eqparts[2])
			return expr(v1, v2, v1/v2, e1, "/", e2)
		}}
	default:
		log.Fatalf("unhandled line: %v", line)
	}
	return "", nil
}
