// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gmlewis/advent-of-code-2021/algorithm"
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

	puz := &puzT{grid: map[keyT][]keyT{}, air: map[keyT][]keyT{}}
	puz = Reduce(lines, puz, parseLine)
	puz = puz.findNeighbors()

	log.Printf("puz.grid: %v keys", len(puz.grid))
	log.Printf("puz.air:  %v keys", len(puz.air))

	cc := algorithm.ConnectedComponents[keyT](puz)
	for k, v := range cc {
		log.Printf("cc[%v]: %v - %v", k, len(v), v)
	}

	printf("Solution: %v\n", len(cc))
}

type keyT [3]int
type puzT struct {
	grid map[keyT][]keyT
	air  map[keyT][]keyT
}
type neighborsT struct {
	key       keyT
	neighbors []keyT
	air       []keyT
}

func (p *puzT) neighbors(ch chan<- *neighborsT, key keyT) {
	var neighbors []keyT
	var air []keyT
	f := func(k keyT) {
		if _, ok := p.grid[k]; ok {
			neighbors = append(neighbors, k)
		} else {
			air = append(air, k)
		}
	}
	f(keyT{key[0] - 1, key[1], key[2]})
	f(keyT{key[0] + 1, key[1], key[2]})
	f(keyT{key[0], key[1] - 1, key[2]})
	f(keyT{key[0], key[1] + 1, key[2]})
	f(keyT{key[0], key[1], key[2] - 1})
	f(keyT{key[0], key[1], key[2] + 1})
	ch <- &neighborsT{key: key, neighbors: neighbors, air: air}
}

func (p *puzT) findNeighbors() *puzT {
	ch := make(chan *neighborsT, 1000)

	var wg sync.WaitGroup
	wg.Add(len(p.grid))
	for k := range p.grid {
		go func(k keyT) {
			p.neighbors(ch, k)
			wg.Done()
		}(k)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	puz := &puzT{grid: map[keyT][]keyT{}, air: map[keyT][]keyT{}}
	for n := range ch {
		if len(n.neighbors) > 0 {
			puz.grid[n.key] = n.neighbors
		}
		for _, airKey := range n.air {
			puz.air[airKey] = []keyT{}
		}
	}

	return puz
}

func parseLine(line string, acc *puzT) *puzT {
	parts := strings.Split(line, ",")
	vs := Map(parts, must.Atoi)
	acc.grid[keyT{vs[0], vs[1], vs[2]}] = []keyT{}
	return acc
}

func (p *puzT) Less(k1, k2 keyT) bool {
	if k1[2] == k2[2] {
		if k1[1] == k2[1] {
			return k1[0] < k2[0]
		}
		return k1[1] < k2[1]
	}
	return k1[2] < k2[2]
}

func (p *puzT) Each(f func(k keyT)) {
	for k := range p.air {
		f(k)
	}
}

func (p *puzT) EachNeighbor(key keyT, eachFn func(from, to keyT)) {
	f := func(k keyT) {
		if _, ok := p.air[k]; ok {
			eachFn(key, k)
		}
	}
	f(keyT{key[0] - 1, key[1], key[2]})
	f(keyT{key[0] + 1, key[1], key[2]})
	f(keyT{key[0], key[1] - 1, key[2]})
	f(keyT{key[0], key[1] + 1, key[2]})
	f(keyT{key[0], key[1], key[2] - 1})
	f(keyT{key[0], key[1], key[2] + 1})
}
