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
	buf := must.ReadFile(filename)

	elvesCalories := strings.Split(buf, "\n\n")
	// log.Printf("%v elves", len(elvesCalories))

	elvesTotals := Map(elvesCalories, func(calLines string) int {
		cals := Map(strings.Split(calLines, "\n"), must.Atoi)
		return Reduce(cals, 0, func(acc, v int) int { return acc + v })
	})
	sort.Ints(elvesTotals)
	// log.Printf("%+v", elvesTotals)

	printf("Solution: %v\n", elvesTotals[len(elvesTotals)-1])
}
