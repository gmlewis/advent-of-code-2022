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
	zeroIndex, mixed := mix(nums)
	v1 := mixed[(zeroIndex+1000)%len(mixed)]
	v2 := mixed[(zeroIndex+2000)%len(mixed)]
	v3 := mixed[(zeroIndex+3000)%len(mixed)]

	printf("Solution: %v\n", v1+v2+v3)
}

func mix(nums []int) (int, []int) {
	original := append([]int{}, nums...)

	indices := nums2map(nums)

	for _, v := range original {
		shift(v, indices)
	}

	return map2nums(indices)
}

func shift(num int, indices map[int]int) {
	oldIndex := indices[num]
	newIndex := num + oldIndex
	for newIndex <= 0 {
		newIndex += len(indices) - 1
	}
	for newIndex >= len(indices) {
		newIndex -= len(indices) + 1
	}

	// var newIndex int
	// switch {
	// case num+oldIndex <= 0:
	// 	newIndex = (num + oldIndex - 1 + len(indices)) % len(indices)
	// case num+oldIndex >= len(indices):
	// 	newIndex = (num + oldIndex + 1) % len(indices)
	// default:
	// 	newIndex = num + oldIndex
	// }

	if newIndex < 0 || newIndex >= len(indices) {
		log.Fatalf("shift(num=%v, indices=%v), oi=%v, ni=%v", num, len(indices), oldIndex, newIndex)
	}

	if oldIndex == newIndex {
		return
	}
	// log.Printf("shift(num=%v, indices=%v), oi=%v, ni=%v", num, indices, oldIndex, newIndex)
	// log.Printf("moving %v from index %v to index %v", num, oldIndex, newIndex)

	dx := newIndex - oldIndex
	if dx > 0 {
		for k, v := range indices {
			if v > oldIndex && v <= newIndex {
				// log.Printf("moving %v from index %v to index %v", k, v, v-1)
				indices[k]--

				if indices[k] < 0 || indices[k] >= len(indices) {
					log.Fatalf("1: k=%v, indices[k]=%v, shift(num=%v, indices=%v), oi=%v, ni=%v", k, indices[k], num, len(indices), oldIndex, newIndex)
				}

			}
		}
		indices[num] = newIndex
		return
	}

	for k, v := range indices {
		if v < oldIndex && v >= newIndex {
			// log.Printf("moving %v from index %v to index %v", k, v, v+1)
			indices[k]++

			if indices[k] < 0 || indices[k] >= len(indices) {
				log.Fatalf("2: k=%v, indices[k]=%v, shift(num=%v, indices=%v), oi=%v, ni=%v", k, indices[k], num, len(indices), oldIndex, newIndex)
			}

		}
	}
	indices[num] = newIndex
}

func map2nums(indices map[int]int) (int, []int) {
	nums := make([]int, len(indices))
	var zeroIndex int
	for k, v := range indices {
		nums[v] = k
		if k == 0 {
			zeroIndex = v
		}
	}
	return zeroIndex, nums
}

func nums2map(nums []int) map[int]int {
	indices := map[int]int{}
	for i, v := range nums {
		indices[v] = i
	}
	return indices
}
