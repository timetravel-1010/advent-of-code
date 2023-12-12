package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	confs = map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	regexps    = map[string]string{
		"red":   `(\d+) red`,
		"green": `(\d+) green`,
		"blue":  `(\d+) blue`,
	}
)

type Game struct {
	ID int
	cubes map[string]int
}

// parseGameInfo returns a map with the color name as key and the sum of occurrences as value.
func parseGameInfo(gameInfo string) Game {
	re := regexp.MustCompile(`Game (\d+):`)
	match := re.FindStringSubmatch(gameInfo)
	fmt.Println("match", match)
	gameID, err := strconv.Atoi(match[1])
	game := Game{ID: gameID, cubes: map[string]int{}}
	if err != nil {
		panic("an error has occurred while converting gameID to int")
	}

	// probably i don't need to split the sets...
	sets := strings.Split(strings.Split(gameInfo, ":")[1], ";")
	//fmt.Println(sets)
	for _, s := range sets {
		fmt.Println("set:", s)
		for color, re := range regexps {
			match := regexp.MustCompile(re).FindStringSubmatch(s)
			if match == nil {
				continue
			}
			// fmt.Println("color:", color, "val:", match[1])
			num, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("error cause: c=",match[1])
				continue
			}
			if num > game.cubes[color] {
				game.cubes[color] = num
			}
		}
	}
	fmt.Println(game.cubes)
	return game
}

func isPossible(cubes map[string]int) bool {
	for color, cubeCount := range cubes {
		if cubeCount > confs[color] {
			fmt.Println("because: color", color, ", count:", cubeCount, ", out of", confs[color])
			return false
		}
	}
	return true
}



func getSum(fileName string) int {
	file, err := os.Open(fileName)

	if err != nil {
		panic("an error has occurred!")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	lines := 0
	for scanner.Scan() {
		game := parseGameInfo(scanner.Text())
		// fmt.Println("id:", game.ID)
		// fmt.Println("cubes:", game.cubes)
		if isPossible(game.cubes) {
			// fmt.Println("id:", game.ID)
			sum += game.ID
		}
		lines++
	}
	fmt.Println(lines, "lines read.")
	return sum
}

func main() {
	//input := "Game 112: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	fmt.Println(getSum("input.txt"))
	fmt.Println(getSum("input-2.txt"))
	fmt.Println(getSum("test.txt"))
}
