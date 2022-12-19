// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

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
	qualities := Map(puz.blueprints, func(b *blueprintT) int { return b.quality() })
	sum := Sum(qualities)

	printf("Solution: %v\n", sum)
}

const (
	totalMinutes = 24
)

type puzT struct {
	blueprints []*blueprintT
}
type blueprintT struct {
	id int

	oreRobotCosts      *materialsT
	clayRobotCosts     *materialsT
	obsidianRobotCosts *materialsT
	geodeRobotCosts    *materialsT

	oreBots      int
	clayBots     int
	obsidianBots int
	geodeBots    int

	ore      int
	clay     int
	obsidian int
	geodes   int
}
type botFunc func(b *blueprintT)
type materialsT struct {
	ore      int
	clay     int
	obsidian int
}

func (b *blueprintT) quality() int {
	for minute := 1; minute <= totalMinutes; minute++ {
		fmt.Printf("\n== Minute %v ==\n", minute)
		botsInQueue := b.buildBots()
		b.collectResources()
		b.addBuiltBots(botsInQueue)
	}

	log.Printf("blueprint %v generated %v geodes", b.id, b.geodes)
	return b.id * b.geodes
}

func (b *blueprintT) buildBots() []botFunc {
	var newBots []botFunc
	for b.canBuild(b.geodeRobotCosts) {
		fmt.Printf("Spend %v ore and %v obsidian to start building a geode-collecting robot.\n", b.geodeRobotCosts.ore, b.geodeRobotCosts.obsidian)
		b.spendResources(b.geodeRobotCosts)
		newBots = append(newBots, buildGeodeBot)
	}
	for b.canBuild(b.obsidianRobotCosts) && !b.shouldSave(b.obsidianRobotCosts, b.geodeRobotCosts) {
		fmt.Printf("Spend %v ore and %v clay to start building an obsidian-collecting robot.\n", b.obsidianRobotCosts.ore, b.obsidianRobotCosts.clay)
		b.spendResources(b.obsidianRobotCosts)
		newBots = append(newBots, buildObsidianBot)
	}
	for b.canBuild(b.clayRobotCosts) && !b.shouldSave(b.clayRobotCosts, b.obsidianRobotCosts) {
		fmt.Printf("Spend %v ore to start building a clay-collecting robot.\n", b.clayRobotCosts.ore)
		b.spendResources(b.clayRobotCosts)
		newBots = append(newBots, buildClayBot)
	}
	for b.canBuild(b.oreRobotCosts) && !b.shouldSave(b.oreRobotCosts, b.clayRobotCosts) {
		fmt.Printf("Spend %v ore to start building an ore-collecting robot.\n", b.oreRobotCosts.ore)
		b.spendResources(b.oreRobotCosts)
		newBots = append(newBots, buildOreBot)
	}
	return newBots
}

func (b *blueprintT) addBuiltBots(inQueue []botFunc) {
	for _, f := range inQueue {
		f(b)
	}
}

func (b *blueprintT) collectResources() {
	b.ore += b.oreBots
	fmt.Printf("%[1]v ore-collecting robots collect %[1]v ore; you now have %[2]v ore.\n", b.oreBots, b.ore)
	b.clay += b.clayBots
	if b.clayBots > 0 {
		fmt.Printf("%[1]v clay-collecting robots collect %[1]v clay; you now have %[2]v clay.\n", b.clayBots, b.clay)
	}
	b.obsidian += b.obsidianBots
	if b.obsidianBots > 0 {
		fmt.Printf("%[1]v obsidian-collecting robots collect %[1]v obsidian; you now have %[2]v obsidian.\n", b.obsidianBots, b.obsidian)
	}
	b.geodes += b.geodeBots
	if b.geodeBots > 0 {
		fmt.Printf("%[1]v geode-cracking robots crack %[1]v geodes; you now have %[2]v open geodes.\n", b.geodeBots, b.geodes)
	}
}

func (b *blueprintT) shouldSave(thisCost, nextCost *materialsT) bool {
	// fmt.Printf("shouldSave: (oreBots=%v,clayBots=%v) (ore=%v,clay=%v) thisCost:(ore=%v,clay=%v), nextCost:(ore=%v,clay=%v)\n", b.oreBots, b.clayBots, b.ore, b.clay, thisCost.ore, thisCost.clay, nextCost.ore, nextCost.clay)

	return b.ore+2*b.oreBots+thisCost.ore >= nextCost.ore &&
		b.clay+2*b.clayBots+thisCost.clay >= nextCost.clay &&
		b.obsidian+2*b.obsidianBots+thisCost.obsidian >= nextCost.obsidian
}

func (b *blueprintT) canBuild(m *materialsT) bool {
	return b.ore >= m.ore && b.clay >= m.clay && b.obsidian >= m.obsidian
}

func (b *blueprintT) spendResources(m *materialsT) {
	b.ore -= m.ore
	if b.ore < 0 {
		log.Fatalf("programming error: blueprint %v ore = %v", b.id, b.ore)
	}
	b.clay -= m.clay
	if b.clay < 0 {
		log.Fatalf("programming error: blueprint %v clay = %v", b.id, b.clay)
	}
	b.obsidian -= m.obsidian
	if b.obsidian < 0 {
		log.Fatalf("programming error: blueprint %v obsidian = %v", b.id, b.obsidian)
	}
}

func buildOreBot(b *blueprintT) {
	b.oreBots++
	fmt.Printf("The new ore-collecting robot is ready; you now have %v of them.\n", b.oreBots)
}
func buildClayBot(b *blueprintT) {
	b.clayBots++
	fmt.Printf("The new clay-collecting robot is ready; you now have %v of them.\n", b.clayBots)
}
func buildObsidianBot(b *blueprintT) {
	b.obsidianBots++
	fmt.Printf("The new obsidian-collecting robot is ready; you now have %v of them.\n", b.obsidianBots)
}
func buildGeodeBot(b *blueprintT) {
	b.geodeBots++
	fmt.Printf("The new geode-cracking robot is ready; you now have %v of them.\n", b.geodeBots)
}

var lineRE = regexp.MustCompile(`Blueprint \d+: Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

func parsePuzzle(lines []string) *puzT {
	puz := &puzT{}

	for i, line := range lines {
		m := lineRE.FindStringSubmatch(line)
		if len(m) != 7 {
			log.Fatalf("Unable to parse puzzle line: %v", line)
		}
		blueprint := &blueprintT{
			id:      i + 1,
			oreBots: 1,

			oreRobotCosts:      &materialsT{ore: must.Atoi(m[1])},
			clayRobotCosts:     &materialsT{ore: must.Atoi(m[2])},
			obsidianRobotCosts: &materialsT{ore: must.Atoi(m[3]), clay: must.Atoi(m[4])},
			geodeRobotCosts:    &materialsT{ore: must.Atoi(m[5]), obsidian: must.Atoi(m[6])},
		}
		puz.blueprints = append(puz.blueprints, blueprint)
	}

	return puz
}
