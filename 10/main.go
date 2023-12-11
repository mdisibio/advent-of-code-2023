package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	fmt.Println("test1.txt")
	i := readInput("test1.txt")
	part1(i)

	fmt.Println("test2.txt")
	i = readInput("test2.txt")
	part1(i)

	fmt.Println("test3.txt")
	i = readInput("test3.txt")
	part2(i, true)

	fmt.Println("test4.txt")
	i = readInput("test4.txt")
	part2(i, true)

	fmt.Println("input.txt")
	i = readInput("input.txt")
	part1(i)
	part2(i, false)
}

func readInput(name string) [][]rune {
	b, _ := os.ReadFile(name)
	lines := strings.Split(string(b), "\n")

	input := make([][]rune, len(lines))
	for y, line := range lines {
		input[y] = []rune(line)
	}
	return input
}

func part1(input [][]rune) {
	var start point
outer:
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				start = point{x, y}
				break outer
			}
		}
	}

	previouslyVisited := map[point]int{}
	visit(input, start, 0, previouslyVisited)

	maxDist := 0
	for _, dist := range previouslyVisited {
		maxDist = max(dist, maxDist)
	}
	fmt.Println("Max dist:", maxDist)
}

func part2(input [][]rune, debug bool) {
	if debug {
		print(input)
	}

	// Rewrite map 2x with spacers in between
	// This is how we find paths "between" pipes
	// On this higher resolution map there is no longer any "between",
	// it will be directly connected.
	input = double(input)

	if debug {
		print(input)
	}

	var start point
outer:
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				start = point{x, y}
				break outer
			}
		}
	}

	mainLoopDistances := map[point]int{}
	visit(input, start, 0, mainLoopDistances)

	// Fill in the outside from the perimeter
	for y := range input {
		for x := range input[y] {
			if x == 0 || x == len(input[y])-1 ||
				y == 0 || y == len(input)-1 {
				fillOutside(input, point{x, y}, mainLoopDistances)
			}
		}
	}

	if debug {
		print(input)
	}

	// Count enclosed spaces. Only where MOD2==0 to count original spaces
	enclosed := 0
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'O' {
				continue
			}
			if _, ok := mainLoopDistances[point{x, y}]; ok {
				continue
			}
			if y%2 == 0 && x%2 == 0 {
				enclosed++
			}
		}
	}
	fmt.Println("Enclosed spaces", enclosed)
}

func visit(input [][]rune, p point, dist int, previouslyVisited map[point]int) {
	// Skip points out of bounds
	if p.x < 0 || p.y < 0 || p.x >= len(input[p.y]) || p.y >= len(input) {
		return
	}

	// Skip if already visited and was shorter (approached via a different path).
	if prev, ok := previouslyVisited[p]; ok && prev < dist {
		return
	}

	// Add this node
	previouslyVisited[p] = dist

	switch input[p.y][p.x] {
	case 'S':
		// Check each cardinal direction for a connection
		// Left
		if p.x > 0 {
			switch input[p.y][p.x-1] {
			case '-', 'F', 'L':
				visit(input, point{p.x - 1, p.y}, dist+1, previouslyVisited)
			}
		}
		// Right
		if p.x < len(input[p.y])-1 {
			switch input[p.y][p.x+1] {
			case '-', 'J', '7':
				visit(input, point{p.x + 1, p.y}, dist+1, previouslyVisited)
			}
		}
		// Down
		if p.y < len(input)-1 {
			switch input[p.y+1][p.x] {
			case '|', 'J', 'L':
				visit(input, point{p.x, p.y + 1}, dist+1, previouslyVisited)
			}
		}
		// Up
		if p.y > 0 {
			switch input[p.y-1][p.x] {
			case '|', '7', 'F':
				visit(input, point{p.x, p.y - 1}, dist+1, previouslyVisited)
			}
		}
	case '-':
		// Left right
		visit(input, point{p.x - 1, p.y}, dist+1, previouslyVisited)
		visit(input, point{p.x + 1, p.y}, dist+1, previouslyVisited)
	case '|':
		// Up down
		visit(input, point{p.x, p.y - 1}, dist+1, previouslyVisited)
		visit(input, point{p.x, p.y + 1}, dist+1, previouslyVisited)
	case '7':
		// Left down
		visit(input, point{p.x - 1, p.y}, dist+1, previouslyVisited)
		visit(input, point{p.x, p.y + 1}, dist+1, previouslyVisited)
	case 'J':
		// Left up
		visit(input, point{p.x - 1, p.y}, dist+1, previouslyVisited)
		visit(input, point{p.x, p.y - 1}, dist+1, previouslyVisited)
	case 'L':
		// Up right
		visit(input, point{p.x, p.y - 1}, dist+1, previouslyVisited)
		visit(input, point{p.x + 1, p.y}, dist+1, previouslyVisited)
	case 'F':
		// Down right
		visit(input, point{p.x, p.y + 1}, dist+1, previouslyVisited)
		visit(input, point{p.x + 1, p.y}, dist+1, previouslyVisited)
	}
}

func double(input [][]rune) [][]rune {
	out := make([][]rune, len(input)*2-1)
	for y := range out {
		out[y] = make([]rune, len(input[0])*2-1)
	}

	for y := range input {
		for x := range input[y] {
			// The copy
			out[y*2][x*2] = input[y][x]
			if x < len(input[y])-1 {
				out[y*2][x*2+1] = doubleMap(input[y][x], input[y][x+1], 'R')
			}
			if y < len(input)-1 {
				out[y*2+1][x*2] = doubleMap(input[y][x], input[y+1][x], 'D')
			}
			// Do diagonals ever connect?
			if x < len(input[y])-1 && y < len(input)-1 {
				out[y*2+1][x*2+1] = '.'
			}
		}
	}
	return out
}

func doubleMap(a, b rune, dir rune) rune {
	switch dir {
	case 'R':
		// Things that connect when going left-to-right
		if (a == '-' || a == 'F' || a == 'L' || a == 'S') &&
			(b == '-' || b == '7' || b == 'J' || b == 'S') {
			return '-'
		}
	case 'D':
		// Things that connect when going up-to-down
		if (a == '|' || a == 'F' || a == '7' || a == 'S') &&
			(b == '|' || b == 'L' || b == 'J' || b == 'S') {
			return '|'
		}
	}
	return '.'
}

func fillOutside(input [][]rune, p point, mainLoopDistances map[point]int) {
	// Skip points out of bounds
	if p.x < 0 || p.y < 0 || p.y >= len(input) || p.x >= len(input[p.y]) {
		return
	}

	// Skip if already visited or part of the main loop
	if input[p.y][p.x] == 'O' {
		return
	}
	if _, ok := mainLoopDistances[p]; ok {
		return
	}

	input[p.y][p.x] = 'O'

	// Fill all 4 cardinal directions
	fillOutside(input, point{p.x - 1, p.y}, mainLoopDistances)
	fillOutside(input, point{p.x + 1, p.y}, mainLoopDistances)
	fillOutside(input, point{p.x, p.y - 1}, mainLoopDistances)
	fillOutside(input, point{p.x, p.y + 1}, mainLoopDistances)
}

func print(input [][]rune) {
	for _, l := range input {
		fmt.Println(string(l))
	}
}
