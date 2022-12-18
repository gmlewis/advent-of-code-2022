// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/gmlewis/advent-of-code-2021/stream"
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

	puz := &puzT{grid: map[keyT]bool{}}
	puz = Reduce(lines, puz, parseLine)
	surfaceArea := puz.surfaceArea()

	printf("Solution: %v\n", surfaceArea)
}

type keyT [3]int
type puzT struct {
	grid map[keyT]bool
}

func (p *puzT) cubeSurfaceArea(ch chan<- int, key keyT) {
	var surfaceArea int
	f := func(k keyT) {
		if !p.grid[k] {
			surfaceArea++
		}
	}
	f(keyT{key[0] - 1, key[1], key[2]})
	f(keyT{key[0] + 1, key[1], key[2]})
	f(keyT{key[0], key[1] - 1, key[2]})
	f(keyT{key[0], key[1] + 1, key[2]})
	f(keyT{key[0], key[1], key[2] - 1})
	f(keyT{key[0], key[1], key[2] + 1})
	ch <- surfaceArea
}

func (p *puzT) surfaceArea() int {
	ch := make(chan int, 1000)

	var wg sync.WaitGroup
	wg.Add(len(p.grid))
	for k := range p.grid {
		go func(k keyT) {
			p.cubeSurfaceArea(ch, k)
			wg.Done()
		}(k)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return stream.Sum(ch)
}

func parseLine(line string, acc *puzT) *puzT {
	parts := strings.Split(line, ",")
	vs := Map(parts, must.Atoi)
	acc.grid[keyT{vs[0], vs[1], vs[2]}] = true
	return acc
}
