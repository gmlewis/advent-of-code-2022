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

	elvesCalories := must.ReadSplitFile(filename, "\n\n")
	// log.Printf("%v elves", len(elvesCalories))

	elvesTotals := Map(elvesCalories, func(calLines string) int {
		cals := Map(strings.Split(calLines, "\n"), must.Atoi)
		return Sum(cals)
	})
	sort.Ints(elvesTotals)
	// log.Printf("%+v", elvesTotals)

	printf("Solution: %v\n", elvesTotals[len(elvesTotals)-1])
}
