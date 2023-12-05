package main

import (
	"fmt"
	"os"
	"strings"
)

func read(name string) []string {
	b, _ := os.ReadFile(name)
	return strings.Split(string(b), "\n")
}

func main() {
	do(read("test.txt"))
	do(read("input.txt"))
}

func do(input []string) {
	total := 0
	copiesWon := make([]int, len(input))
	for i, line := range input {
		var (
			line    = line[strings.Index(line, ":")+1:]
			parts   = strings.Split(line, " ")
			winning = map[string]struct{}{}
			scoring = false
			matches = 0
		)
		for _, n := range parts {
			switch n {
			case "":
			case "|":
				scoring = true
			default:
				if scoring {
					if _, ok := winning[n]; ok {
						matches++
					}
				} else {
					winning[n] = struct{}{}
				}
			}
		}
		if matches > 0 {
			total += 1 << (matches - 1)
			for j := 1; j <= matches && (i+j < len(input)); j++ {
				copiesWon[i+j] += copiesWon[i] + 1
			}
		}
	}
	fmt.Println(total)
	total = 0
	for _, x := range copiesWon {
		total += x
	}
	fmt.Println(total + len(input))
}
