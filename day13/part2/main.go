// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
	lines := must.ReadFileLines(filename)
	lines = append(lines, "[[2]]", "[[6]]")
	lines = Filter(lines, func(s string) bool { return s != "" })
	packets := Map(lines, func(s string) *packetT { return parsePacket(s) })

	sort.Slice(packets, func(a, b int) bool { return compare(packets[a], packets[b]) < 0 })
	decoderKey := 1
	for i, p := range packets {
		s := p.String()
		switch s {
		case "[[2]]":
			decoderKey *= (i + 1)
		case "[[6]]":
			decoderKey *= (i + 1)
			break
		}
	}

	printf("Solution: %v\n", decoderKey)
}

type packetT struct {
	v   *int
	els []*packetT
}

func isInOrder(pair string) bool {
	parts := strings.Split(pair, "\n")
	p1 := parsePacket(parts[0])
	p2 := parsePacket(parts[1])
	return compare(p1, p2) <= 0
}

func compare(p1, p2 *packetT) int {
	switch {
	// If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
	case p1 == nil && p2 == nil:
		return 0
	// If the left list runs out of items first, the inputs are in the right order.
	case p1 == nil:
		return -1
	// If the right list runs out of items first, the inputs are not in the right order.
	case p2 == nil:
		return 1

	// If both values are integers, the lower integer should come first.
	// If the left integer is lower than the right integer, the inputs are in the right order.
	case p1.v != nil && p2.v != nil && *p1.v < *p2.v:
		return -1
	// If the left integer is higher than the right integer, the inputs are not in the right order.
	case p1.v != nil && p2.v != nil && *p1.v > *p2.v:
		return 1
	// Otherwise, the inputs are the same integer; continue checking the next part of the input.
	case p1.v != nil && p2.v != nil && *p1.v == *p2.v:
		return 0

	// If both values are lists, compare the first value of each list, then the second value, and so on.
	case p1.v == nil && p2.v == nil:
		n := len(p1.els)
		if len(p2.els) > n {
			n = len(p2.els)
		}

		for i := 0; i < n; i++ {
			var el1, el2 *packetT
			if i < len(p1.els) {
				el1 = p1.els[i]
			}
			if i < len(p2.els) {
				el2 = p2.els[i]
			}
			cmp := compare(el1, el2)
			if cmp != 0 {
				return cmp
			}
		}

		if len(p1.els) < len(p2.els) {
			return -1
		}
		if len(p1.els) > len(p2.els) {
			return 1
		}
		return 0

	// If exactly one value is an integer, convert the integer to a list which contains that integer
	// as its only value, then retry the comparison.
	// For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2);
	// the result is then found by instead comparing [0,0,0] and [2].
	case p1.v == nil: // p1 is a list. Convert p2 to list.
		right := &packetT{els: []*packetT{p2}}
		return compare(p1, right)

	case p2.v == nil: // p2 is a list. Convert p1 to list.
		left := &packetT{els: []*packetT{p1}}
		return compare(left, p2)

	default:
		log.Fatalf("Unhandled case: packetsEqual:\np1=%#v\np2=%#v", p1, p2)
		return 1
	}
}

func (p *packetT) String() string {
	if p == nil {
		return "<nil>"
	}
	switch {
	case p.v != nil:
		return fmt.Sprintf("%v", *p.v)
	default:
		parts := make([]string, 0, len(p.els))
		for _, el := range p.els {
			parts = append(parts, el.String())
		}
		return fmt.Sprintf("[%v]", strings.Join(parts, ","))
	}
}

var (
	numRE = regexp.MustCompile(`^(\d+)`)
)

func parsePacket(s string) *packetT {
	p, next := parseNext(nil, s)
	if next != "" {
		log.Fatalf("parsePacket('%v'): next has value: %q\n%v", s, next, p)
	}
	return p
}

func parseNext(p *packetT, s string) (root *packetT, next string) {
	m := numRE.FindStringSubmatch(s)
	if len(m) == 2 {
		v := must.Atoi(m[1])
		if p == nil {
			log.Fatalf("parseNext('%v') expected p!=nil", s)
		}
		p.els = append(p.els, &packetT{v: &v})
		next = s[len(m[1]):]
		return p, next
	}

	switch {
	case strings.HasPrefix(s, "["):
		child := &packetT{}
		if p == nil {
			p = child
		} else {
			p.els = append(p.els, child)
		}
		next = s[1:]
		for next != "" {
			if strings.HasPrefix(next, "]") {
				return p, next[1:]
			}
			_, next = parseNext(child, next)
			if strings.HasPrefix(next, ",") {
				next = next[1:]
			}
		}
	}

	return p, next
}
