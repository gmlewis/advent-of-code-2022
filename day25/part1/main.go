// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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

	var sum int
	for _, line := range lines {
		sum += snafu2dec(line)
	}
	solution := dec2snafu(sum)

	printf("Solution: %v\n", solution)
}

func snafu2dec(s string) int {
	var v int
	slen := len(s)
	for i, r := range s {
		pow := int(math.Pow(5, float64(slen-i-1)))
		switch r {
		case '2':
			v += 2 * pow
		case '1':
			v += pow
		case '0':
		case '-':
			v -= pow
		case '=':
			v -= 2 * pow
		}
	}
	return v
}

func dec2snafu(v int) string {
	slen := 1
	for {
		if v <= snafu2dec(strings.Repeat("2", slen)) {
			break
		}
		slen++
	}

	var result string
	for slen >= 0 {
		slen--
		var suffix string
		if slen > 0 {
			suffix = strings.Repeat("=", slen)
		}
		switch {
		case v >= snafu2dec(result+"2"+suffix):
			result += "2"
		case v >= snafu2dec(result+"1"+suffix):
			result += "1"
		case v >= snafu2dec(result+"0"+suffix):
			result += "0"
		case v >= snafu2dec(result+"-"+suffix):
			result += "-"
		case v >= snafu2dec(result+"="+suffix):
			result += "="
		}
	}
	return result
}
