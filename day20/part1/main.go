// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

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
	lines := must.ReadFileLines(filename)

	nums := Map(lines, must.Atoi)
	// log.Printf("%v nums=%v ...", len(nums), nums[:7])
	zeroIndex, mixed := mix(nums)
	// log.Printf("zeroIndex=%v, mixed[zeroIndex]=%v", zeroIndex, mixed[zeroIndex])
	v1 := mixed[(zeroIndex+1000)%len(mixed)]
	v2 := mixed[(zeroIndex+2000)%len(mixed)]
	v3 := mixed[(zeroIndex+3000)%len(mixed)]

	printf("Solution: %v\n", v1+v2+v3)
}

func mix(nums []int) (int, []int) {
	indices := make([]int, len(nums))
	for i := range nums {
		indices[i] = i
	}

	for i, v := range nums {
		shift(i, v, indices)
	}

	return reorder(nums, indices)
}

func shift(index, num int, indices []int) {
	if num == 0 {
		return
	}
	oldIndex := indices[index]
	newIndex := num + oldIndex
	// log.Printf("shift(index=%v, num=%v, len(indices)=%v), oldIndex=%v, newIndex=%v", index, num, len(indices), oldIndex, newIndex)
	for newIndex <= 0 {
		newIndex += len(indices) - 1
	}
	for newIndex >= len(indices) {
		newIndex -= (len(indices) - 1)
	}
	if oldIndex == newIndex {
		return
	}
	// log.Printf("newIndex=%v", newIndex)

	dx := newIndex - oldIndex
	// log.Printf("moving %v from index %v to index %v (dx=%v)", num, oldIndex, newIndex, dx)
	if dx > 0 {
		for i, v := range indices {
			if v > oldIndex && v <= newIndex {
				indices[i]--
				if indices[i] < 0 || indices[i] >= len(indices) {
					log.Fatalf("A: ERROR: indices[%v]=%v", i, indices[i])
				}
			}
		}
		indices[index] = newIndex
		if indices[index] < 0 || indices[index] >= len(indices) {
			log.Fatalf("B: ERROR: indices[%v]=%v", index, indices[index])
		}
	} else {
		for i, v := range indices {
			if v < oldIndex && v >= newIndex {
				indices[i]++
				if indices[i] < 0 || indices[i] >= len(indices) {
					log.Fatalf("C: ERROR: indices[%v]=%v", i, indices[i])
				}
			}
		}
		indices[index] = newIndex
		if indices[index] < 0 || indices[index] >= len(indices) {
			log.Fatalf("D: ERROR: indices[%v]=%v", index, indices[index])
		}
	}
}

func reorder(nums, indices []int) (int, []int) {
	result := make([]int, len(nums))
	var zeroIndex int
	for i, v := range nums {
		result[indices[i]] = v
		if v == 0 {
			zeroIndex = indices[i]
		}
	}
	return zeroIndex, result
}
