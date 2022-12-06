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

	pairs := must.ReadFileLines(filename)
	fullOverlaps := Filter(pairs, findOverlap)

	printf("Solution: %v\n", len(fullOverlaps))
}

func findOverlap(s string) bool {
	parts := strings.Split(s, ",")
	range1 := strings.Split(parts[0], "-")
	range2 := strings.Split(parts[1], "-")
	x1 := must.Atoi(range1[0])
	x2 := must.Atoi(range1[1])
	x3 := must.Atoi(range2[0])
	x4 := must.Atoi(range2[1])
	return (x1 >= x3 && x2 <= x4) || (x3 >= x1 && x4 <= x2)
}
