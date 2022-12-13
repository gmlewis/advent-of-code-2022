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
	pairs := must.ReadSplitFile(filename, "\n\n")

	var sum int
	for i, pair := range pairs {
		if isInOrder(pair) {
			sum += (i + 1)
		}
	}

	printf("Solution: %v\n", sum)
}

type packetT struct {
	v   *int
	els []*packetT
}

func isInOrder(pair string) bool {
	// parts := strings.Split(pair, "\n")
	// p1 := parsePacket(parts[0])
	// p2 := parsePacket(parts[1])
	return true
}

func parsePacket(s string) *packetT {
	p := &packetT{}
	if next, _ := p.parseNext(s); next != "" {
		log.Fatalf("parsePacket('%v'): next has value: %q\n%v", s, next, p)
	}
	return p
}

var (
	numRE = regexp.MustCompile(`^(\d+)`)
)

func (p *packetT) String() string {
}

func (p *packetT) parseNext(s string) (next string, eol bool) {
	switch {
	case s == "":
		return "", false
	case strings.HasPrefix(s, ","):
		return s[1:], false
	case strings.HasPrefix(s, "]"):
		return s[1:], true
	case strings.HasPrefix(s, "["):
		next = s[1:]
		for {
			sp := &packetT{}
			next, eol = sp.parseNext(next)
			p.els = append(p.els, sp)
			if eol {
				return next, false
			}
			switch {
			case strings.HasPrefix(next, ","):
				next = next[1:]
			case strings.HasPrefix(next, "]"):
				next = next[1:]
				return next, true
			default:
				log.Fatalf("parseNext('%v') expected comma or ']', but got %q", s, next)
			}
		}
	default:
		m := numRE.FindStringSubmatch(s)
		if len(m) == 2 {
			v := must.Atoi(m[1])
			p.v = &v
			return s[len(m[1]):], false
		}
		log.Fatalf("parseNext unhandled case: %q", s)
	}
	return "error", false
}
