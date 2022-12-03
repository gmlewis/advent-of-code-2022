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
	buf := must.ReadFile(filename)

	sacks := strings.Split(buf, "\n")
	groups := ChunkEvery(sacks, 3, 3)
	dups := Map(groups, findDups)
	// log.Printf("dups=%+v", dups)
	priorities := Map(dups, itemPriority)
	sum := Reduce(priorities, 0, func(acc, v int) int { return acc + v })

	printf("Solution: %v\n", sum)
}

func findDups(sacks []string) string {
	f := func(sack string) map[rune]int {
		result := map[rune]int{}
		for _, r := range sack {
			result[r]++
		}
		return result
	}

	in1, in2 := f(sacks[0]), f(sacks[1])
	for _, r := range sacks[2] {
		if _, ok := in1[r]; !ok {
			continue
		}
		if _, ok := in2[r]; !ok {
			continue
		}
		return string(r)
	}
	log.Fatalf("Could not find dup in %+v", sacks)
	return ""
}

const priorities = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func itemPriority(item string) int {
	return strings.Index(priorities, item)
}
