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
	lines := must.ReadFileLines(filename)

	puz := parsePuzzle(lines)

	rowY := 10
	if strings.Contains(filename, "input.txt") {
		rowY = 2000000
	}

	solution := puz.noBeacons(rowY)

	printf("Solution: %v\n", solution)
}

type keyT [2]int
type puzT struct {
	xmin int
	xmax int
	ymin int
	ymax int
	grid map[keyT]pairT
}
type pairT struct {
	sensor keyT
	beacon keyT
	dx     int // positive
	dy     int // positive
}

func (p *puzT) noBeacons(row int) int {
	return 0
}

var lineRE = regexp.MustCompile(`=(-?\d+)`)

func parsePuzzle(lines []string) *puzT {
	p := &puzT{grid: map[keyT]rune{}}
	for _, line := range lines {
		m := lineRE.FindAllStringSubmatch(line, -1)
		if len(m) != 4 {
			log.Fatalf("Unable to parse line: %v, %#v", line, m)
		}
		sx := must.Atoi(m[0][1])
		sy := must.Atoi(m[1][1])
		p.grid[keyT{sx, sy}] = 'S'
		bx := must.Atoi(m[2][1])
		by := must.Atoi(m[3][1])
		p.grid[keyT{bx, by}] = 'S'
	}
	return p
}
