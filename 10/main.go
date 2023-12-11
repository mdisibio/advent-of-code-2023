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
	i := readInput("test1.txt")
	part1(i)

	i = readInput("test2.txt")
	part1(i)

	i = readInput("input.txt")
	part1(i)
}

func readInput(name string) []string {
	b, _ := os.ReadFile(name)
	lines := strings.Split(string(b), "\n")
	return lines
}

func part1(input []string) {
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

func visit(input []string, p point, dist int, previouslyVisited map[point]int) {
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
