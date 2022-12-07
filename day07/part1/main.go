// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
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
	cmds := must.ReadSplitFile(filename, "$ ")

	filesystem := parseCmds(cmds)
	allSizes := maps.Values(filesystem)
	smallDirs := Filter(allSizes, func(v int64) bool { return v <= 100000 })
	solution := Sum(smallDirs)

	printf("Solution: %v\n", solution)
}

type puzT map[string]int64

func parseCmds(cmds []string) puzT {
	puz := puzT{}
	pwd := "/"
	parents := map[string][]string{}

	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		switch {
		case cmd == "cd /":
			pwd = "/"
		case cmd == "cd ..":
			if len(parents[pwd]) < 1 {
				log.Fatalf("Attempt to cd above / from pwd=%v", pwd)
			}
			pwd = parents[pwd][0]
		case strings.HasPrefix(cmd, "cd "):
			newPwd := pwd + cmd[3:] + "/"
			if _, ok := puz[newPwd]; ok {
				log.Fatalf("Already visited directory '%v'!!!", newPwd)
			}
			parents[newPwd] = append([]string{pwd}, parents[pwd]...)
			pwd = newPwd
		case strings.HasPrefix(cmd, "ls\n"):
			for _, file := range strings.Split(cmd, "\n")[1:] {
				if strings.HasPrefix(file, "dir ") {
					continue
				}
				parts := strings.Split(file, " ")
				size := int64(must.Atoi(parts[0]))
				puz[pwd] += size
				for _, parent := range parents[pwd] {
					puz[parent] += size
				}
			}
		default:
			log.Fatalf("Unsupported command: '%v'", cmd)
		}
	}

	return puz
}
