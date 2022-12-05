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
	"github.com/gmlewis/advent-of-code-2021/maps"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var logf = log.Printf
var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

var (
	movesRE = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
)

func process(filename string) {
	logf("Processing %v ...", filename)
	buf := must.ReadFile(filename)

	parts := strings.Split(buf, "\n\n")
	stacks := parseStacks(parts[0])
	stackKeys := maps.Keys(stacks)

	stacks = moveCrates(stacks, parts[1])

	sort.Strings(stackKeys)
	solution := Reduce(stackKeys, "", takeTop(stacks))

	printf("Solution: %v\n", solution)
}

type puzT map[string][]string

func moveCrates(puz puzT, cmds string) puzT {
	for _, cmd := range strings.Split(strings.TrimSpace(cmds), "\n") {
		m := movesRE.FindStringSubmatch(cmd)
		if len(m) != 4 {
			log.Fatalf("Bad move: %v", cmd)
		}
		num := must.Atoi(m[1])
		from := m[2]
		to := m[3]
		puz.move(num, from, to)
	}
	return puz
}

func (p puzT) move(num int, from, to string) {
	if len(p[from]) < num {
		log.Fatalf("Cannot move %v from %v (%+v) to %v (%+v)!", num, from, p[from], to, p[to])
	}
	for i := 0; i < num; i++ {
		n := len(p[from]) - 1
		crate := p[from][n]
		p[to] = append(p[to], crate)
		p[from] = p[from][0:n]
	}
}

func takeTop(stacks puzT) func(stack, acc string) string {
	return func(stack, acc string) string {
		s := stacks[stack]
		if len(s) == 0 {
			return acc
		}
		return acc + s[len(s)-1]
	}
}

func parseStacks(s string) puzT {
	result := puzT{}
	lines := strings.Split(s, "\n")
	lastLine := len(lines) - 1
	for i := lastLine - 1; i >= 0; i-- {
		line := lines[i]
		for x := 1; x < len(line); x += 4 {
			crate := line[x : x+1]
			if crate != " " {
				stack := lines[lastLine][x : x+1]
				result[stack] = append(result[stack], crate)
			}
		}
	}
	return result
}
