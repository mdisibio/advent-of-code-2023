package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}

	var (
		data      = strings.Split(string(b), "\n")
		ids       = map[coord]string{}
		symbols   = map[coord]rune{}
		currCoord = coord{}
		currID    = ""
	)

	endCurr := func() {
		if currID != "" {
			ids[currCoord] = currID
			currID = ""
		}
	}

	for y, line := range data {
		for x, ch := range line {
			switch ch {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if currID == "" {
					currCoord.x = x
					currCoord.y = y
				}
				currID += string(ch)
			case '.':
				endCurr()
			default:
				symbols[coord{x, y}] = ch
				endCurr()
			}
		}
		endCurr()
	}

	// Part 1 - All IDs attached to at least 1 symbol
	sum := 0
	for p, id := range ids {
		// Look around the perimeter for symbols
		for x := p.x - 1; x <= p.x+len(id); x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				if _, ok := symbols[coord{x, y}]; ok {
					v, _ := strconv.Atoi(id)
					sum += v
					break
				}
			}
		}
	}
	fmt.Println("Sum of valid part ids", sum)

	// Part 2: look for all gears
	sum = 0
	for p, sym := range symbols {
		if sym != '*' {
			continue
		}

		attached := []int{}
		// Look around for part IDs. Asssumes at most 3 digit numbers
		for x := p.x - 3; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				if id, ok := ids[coord{x, y}]; ok {
					// We found an ID within range
					// but double check it actually is attached
					// This is shorter than 3 digit numbers
					if x+len(id) <= p.x-1 {
						// Too short ex:
						// 12.....
						// ...*...
						continue
					}
					v, _ := strconv.Atoi(id)
					attached = append(attached, v)
				}
			}
		}
		if len(attached) == 2 {
			sum += attached[0] * attached[1]
		}
	}
	fmt.Println("Sum of gear ratios:", sum)
}
