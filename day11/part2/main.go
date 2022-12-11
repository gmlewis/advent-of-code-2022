// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
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
	buf := must.ReadSplitFile(filename, "\n\n")

	puz := parsePuz(buf)
	// log.Printf("lcm=%v", puz.lcm)
	puz.simulate(10000)

	printf("Solution: %v\n", puz.solution())
}

type puzT struct {
	monkies []*monkeyT
	lcm     int
}
type monkeyT struct {
	inspections int
	items       []numType
	op          opFunc
	throw       throwFunc
}
type opFunc func(numType) numType
type throwFunc func(newValue numType) int
type numType = int

func (p *puzT) round() {
	for _, m := range p.monkies {
		// log.Printf("Monkey %v:", i)
		m.turn(p)
	}
}

func (m *monkeyT) turn(p *puzT) {
	for _, item := range m.items {
		m.inspections++
		// log.Printf("  Monkey inspects an item with a worry level of %v.", item)
		wl := m.op(item)
		if wl > p.lcm {
			wl = wl % p.lcm
		}
		// log.Printf("    Worry level is now %v.", wl)
		// log.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %v.", wl)
		toMonkey := m.throw(wl)
		// log.Printf("    Item with worry level %v is thrown to monkey %v.", wl, toMonkey)
		p.monkies[toMonkey].items = append(p.monkies[toMonkey].items, wl)
	}
	m.items = nil
}

func (p *puzT) simulate(rounds int) {
	for i := 0; i < rounds; i++ {
		p.round()
		// if i == 0 || i == 19 || (i+1)%1000 == 0 {
		// 	fmt.Printf("\n== After round %v ==\n", i+1)
		// 	for j, m := range p.monkies {
		// 		fmt.Printf("Monkey %v inspected items %v times.\n", j, m.inspections)
		// 	}
		// }
	}
}

func (p *puzT) solution() int64 {
	n := len(p.monkies)
	inspections := make([]int, 0, n)
	for _, m := range p.monkies {
		inspections = append(inspections, m.inspections)
	}
	sort.Ints(inspections)
	return int64(inspections[n-1]) * int64(inspections[n-2])
}

func parsePuz(monkeySections []string) *puzT {
	puz := &puzT{lcm: 1, monkies: make([]*monkeyT, 0, len(monkeySections))}
	for _, ms := range monkeySections {
		parts := strings.Split(ms, "\n")
		m := &monkeyT{}
		puz.monkies = append(puz.monkies, m)
		m.parseItems(parts[1])
		m.parseOperation(parts[2])
		m.parseTest(puz, parts[3:])
	}
	return puz
}

func (m *monkeyT) parseItems(s string) {
	ii := 2 + strings.Index(s, ": ")
	m.items = Map(
		strings.Split(strings.Replace(s[ii:], " ", "", -11), ","), must.Atoi)
}

func (m *monkeyT) parseOperation(s string) {
	ii := 6 + strings.Index(s, "= old ")
	parts := strings.Split(s[ii:], " ")
	var value numType
	if parts[1] != "old" {
		value = must.Atoi(parts[1])
		switch parts[0] {
		case "+":
			m.op = func(old numType) numType { return old + value }
		case "*":
			m.op = func(old numType) numType { return old * value }
		default:
			log.Fatalf("unable to parse: %v", s)
		}
		return
	}

	switch parts[0] {
	case "+":
		m.op = func(old numType) numType { return old + old }
	case "*":
		m.op = func(old numType) numType { return old * old }
	default:
		log.Fatalf("unable to parse: %v", s)
	}
}

func (m *monkeyT) parseTest(p *puzT, s []string) {
	const str = "divisible by "
	ii := len(str) + strings.Index(s[0], str)
	v := must.Atoi(s[0][ii:])
	p.lcm *= v
	ii = 1 + strings.LastIndex(s[1], " ")
	trueMonkey := must.Atoi(s[1][ii:])
	ii = 1 + strings.LastIndex(s[2], " ")
	falseMonkey := must.Atoi(s[2][ii:])
	m.throw = func(newValue numType) int {
		if newValue%v == 0 {
			return trueMonkey
		}
		return falseMonkey
	}
}
