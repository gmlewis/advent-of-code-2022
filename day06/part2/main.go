// -*- compile-command: "go run main.go ../example*.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

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

	parts := ChunkEvery([]rune(buf), 14, 1)
	index := 14 + FindFirst(parts, allUnique)

	printf("Solution: %v\n", index)
}

func allUnique(s []rune) bool {
	m := map[rune]bool{}
	for _, r := range s {
		if m[r] {
			return false
		}
		m[r] = true
	}
	log.Printf("found: %v", string(s))
	return true
}
