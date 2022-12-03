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

	games := strings.Split(buf, "\n")
	// log.Printf("%v games", len(games))

	scores := Map(games, scoreGame)
	// log.Printf("%+v", scores)

	totalScore := Reduce(scores, 0, func(acc, v int) int { return acc + v })
	printf("Solution: %v\n", totalScore)
}

func scoreGame(game string) int {
	switch game {
	case "A X":
		return 1 + 3
	case "A Y":
		return 2 + 6
	case "A Z":
		return 3 + 0
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 1 + 6
	case "C Y":
		return 2 + 0
	case "C Z":
		return 3 + 3
	default:
		log.Fatalf("bad game: %q", game)
	}
	return 0
}
