package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		var first, last string
		for i := 0; i < len(l); i++ {
			if l[i] >= '0' && l[i] <= '9' {
				first = l[i : i+1]
				break
			}
		}
		for i := len(l) - 1; i >= 0; i-- {
			if l[i] >= '0' && l[i] <= '9' {
				last = l[i : i+1]
				break
			}
		}
		v, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	fmt.Println("Total is:", sum)
}
