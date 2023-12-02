package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input3.txt")
	if err != nil {
		panic(err)
	}

	lookups := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	lookup := func(l string, pos int) (int, bool) {
		for k, v := range lookups {
			if strings.HasPrefix(l[pos:], k) {
				return v, true
			}
		}
		return 0, false
	}

	var (
		lines = strings.Split(string(b), "\n")
		sum   int
		first int
		last  int
		ok    bool
	)

	for _, l := range lines {
		for i := 0; i < len(l); i++ {
			if first, ok = lookup(l, i); ok {
				break
			}
		}
		for i := len(l) - 1; i >= 0; i-- {
			if last, ok = lookup(l, i); ok {
				break
			}
		}

		sum += first*10 + last
	}
	fmt.Println("Total is:", sum)
}
