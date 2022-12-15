// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
	lines := must.ReadFileLines(filename)

	puz := parsePuzzle(lines)

	maxValue := 20
	if strings.Contains(filename, "input.txt") {
		maxValue = 4000000
	}

	beacon := puz.beaconPosition(maxValue)

	solution := beacon.x()*4000000 + beacon.y()
	printf("Solution: %v\n", solution)
}

type keyT [2]int

func (k keyT) x() int { return k[0] }
func (k keyT) y() int { return k[1] }

type puzT struct {
	xmin    int
	xmax    int
	ymin    int
	ymax    int
	grid    map[keyT]rune
	sensors map[keyT]*pairT // keyed by sensor location
}
type pairT struct {
	beacon keyT
	dx     int // positive
	dy     int // positive
}

func (p *puzT) beaconPosition(maxValue int) keyT {
	for y := 0; y <= maxValue; y++ {
		if beacon, ok := p.possiblePosition(y, maxValue); ok {
			return beacon
		}
	}
	log.Fatalf("beacon position not found")
	return keyT{}
}

func (p *puzT) possiblePosition(y, maxValue int) (keyT, bool) {
	var notPossible []keyT
	for s, pair := range p.sensors {
		if xRange, ok := pair.notPossibleXRange(s, y, maxValue); ok {
			notPossible = append(notPossible, xRange)
		}
	}
	if len(notPossible) == 0 {
		log.Fatalf("notPossible=nil at y=%v", y)
	}
	sort.Slice(notPossible, func(a, b int) bool { return notPossible[a].x() < notPossible[b].x() })
	// log.Printf("y=%v, notPossible=%+v", y, notPossible)

	if notPossible[0][0] > 0 {
		return keyT{0, y}, true
	}

	xMax := notPossible[0][1]
	for _, xRange := range notPossible[1:] {
		if xRange[0] > xMax {
			return keyT{xRange[0] - 1, y}, true
		}
		if xRange[1] > xMax {
			xMax = xRange[1]
		}
	}

	return keyT{}, false
}

func (p *pairT) notPossibleXRange(sensor keyT, y, maxValue int) (keyT, bool) {
	dy := sensor.y() - y
	if dy < 0 {
		dy = -dy
	}

	dx := p.dy + p.dx - dy // remaining manhattan distance after checking y delta
	if dx < 0 {
		return keyT{}, false // sensor doesn't reach at this y value.
	}

	xRange := keyT{sensor.x() - dx, sensor.x() + dx}
	if xRange[1] < 0 || xRange[0] > maxValue {
		return keyT{}, false // out of bounds
	}

	if xRange[0] < 0 {
		xRange[0] = 0
	}
	if xRange[1] > maxValue {
		xRange[1] = maxValue
	}

	return xRange, true
}

// func (p *puzT) noBeacons(y int) int {
// 	var impossibleBeacons int
// 	// log.Printf("Searching from x=%v to x=%v", p.xmin, p.xmax)
// 	for x := p.xmin; x <= p.xmax; x++ {
// 		if !p.isBeaconPossible(x, y) {
// 			impossibleBeacons++
// 		}
// 	}
// 	return impossibleBeacons
// }
//
// func (p *puzT) isBeaconPossible(x, y int) bool {
// 	if p.grid[keyT{x, y}] == 'B' {
// 		return true
// 	}
// 	for s, pair := range p.sensors {
// 		if !pair.isBeaconPossible(s, x, y) {
// 			// log.Printf("beacon impossible at (%v,%v) because of sensor %v", x, y, s)
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func (p *pairT) isBeaconPossible(sensor keyT, x, y int) bool {
// 	dx := sensor.x() - x
// 	if dx < 0 {
// 		dx = -dx
// 	}
// 	dy := sensor.y() - y
// 	if dy < 0 {
// 		dy = -dy
// 	}
// 	possible := dx+dy > p.dx+p.dy
// 	return possible
// }

var lineRE = regexp.MustCompile(`=(-?\d+)`)

func parsePuzzle(lines []string) *puzT {
	p := &puzT{grid: map[keyT]rune{}, sensors: map[keyT]*pairT{}}
	for _, line := range lines {
		m := lineRE.FindAllStringSubmatch(line, -1)
		if len(m) != 4 {
			log.Fatalf("Unable to parse line: %v, %#v", line, m)
		}
		sx := must.Atoi(m[0][1])
		sy := must.Atoi(m[1][1])
		sensor := keyT{sx, sy}
		bx := must.Atoi(m[2][1])
		by := must.Atoi(m[3][1])
		beacon := keyT{bx, by}
		dx, dy := sx-bx, sy-by
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		p.grid[sensor] = 'S'
		p.grid[beacon] = 'B'
		p.updateBounds(sx-dx, sy-dy)
		p.updateBounds(sx+dx, sy+dy)
		p.updateBounds(bx-dx, by-dy)
		p.updateBounds(bx+dx, by+dy)
		p.sensors[sensor] = &pairT{
			beacon: beacon,
			dx:     dx,
			dy:     dy,
		}
	}
	return p
}

func (p *puzT) updateBounds(x, y int) {
	if x < p.xmin {
		p.xmin = x
	}
	if x > p.xmax {
		p.xmax = x
	}
	if y < p.ymin {
		p.ymin = y
	}
	if y > p.ymax {
		p.ymax = y
	}
}
