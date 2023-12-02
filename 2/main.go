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

	lines := strings.Split(string(b), "\n")
	sumIDs := 0
	sumMins := 0

	for i, game := range lines {
		game = game[strings.Index(game, ":")+1:]

		var (
			valid = true
			minR  int
			minG  int
			minB  int
		)

		rounds := strings.Split(game, ";")
		for _, round := range rounds {

			var r, g, b int

			draws := strings.Split(round, ",")
			for _, draw := range draws {
				draw := strings.TrimSpace(draw)

				parts := strings.Split(draw, " ")
				if len(parts) == 2 {
					switch parts[1] {
					case "red":
						r, _ = strconv.Atoi(parts[0])
					case "green":
						g, _ = strconv.Atoi(parts[0])
					case "blue":
						b, _ = strconv.Atoi(parts[0])
					}
				}
			}

			if r > 12 ||
				g > 13 ||
				b > 14 {
				valid = false
			}

			minR = max(r, minR)
			minG = max(g, minG)
			minB = max(b, minB)
		}

		if valid {
			sumIDs += i + 1
		}
		sumMins += minR * minG * minB
	}
	fmt.Println(sumIDs)
	fmt.Println(sumMins)
}
