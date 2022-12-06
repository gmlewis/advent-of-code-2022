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

	sacks := must.ReadFileLines(filename)
	dups := Map(sacks, findDups)
	priorities := Map(dups, itemPriority)
	sum := Sum(priorities)

	printf("Solution: %v\n", sum)
}

func findDups(sack string) string {
	n := len(sack) / 2
	in1, in2 := map[rune]int{}, map[rune]int{}
	for i, r := range sack {
		if i < n {
			in1[r]++
			continue
		}
		in2[r]++
		if _, ok := in1[r]; ok {
			return string(r)
		}
	}
	log.Fatalf("Could not find dup in %q", sack)
	return ""
}

const priorities = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func itemPriority(item string) int {
	return strings.Index(priorities, item)
}
