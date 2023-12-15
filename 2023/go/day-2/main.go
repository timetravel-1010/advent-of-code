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
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	regexps = map[string]string{
		"red":   `(\d+) red`,
		"green": `(\d+) green`,
		"blue":  `(\d+) blue`,
	}
)

type Game struct {
	ID    int
	cubes map[string]int
}

// parseGameInfo returns a map with the color name as key and the sum of occurrences as value.
func parseGameInfo(gameInfo string) Game {
	re := regexp.MustCompile(`Game (\d+):`)
	match := re.FindStringSubmatch(gameInfo)
	gameID, err := strconv.Atoi(match[1])
	game := Game{ID: gameID, cubes: map[string]int{}}
	if err != nil {
		panic("an error has occurred while converting gameID to int")
	}

	// probably i don't need to split the sets...
	sets := strings.Split(strings.Split(gameInfo, ":")[1], ";")
	for _, s := range sets {
		for color, re := range regexps {
			match := regexp.MustCompile(re).FindStringSubmatch(s)
			if match == nil {
				continue
			}
			num, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}
			if num > game.cubes[color] {
				game.cubes[color] = num
			}
		}
	}
	return game
}

func isPossible(cubes map[string]int) bool {
	for color, cubeCount := range cubes {
		if cubeCount > confs[color] {
			return false
		}
	}
	return true
}

func getPower(cubes map[string]int) int {
	total := 1
	for _, v := range cubes {
		total *= v
	}
	return total
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
		// if isPossible(game.cubes) {
		// 	sum += game.ID
		// }
		sum += getPower(game.cubes)
		lines++
	}
	fmt.Println(lines, "lines read.")
	return sum
}

func main() {
	//input := "Game 112: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	fmt.Println(getSum("./day-2/input.txt"))
	// fmt.Println(getSum("input-2.txt"))
	// fmt.Println(getSum("test.txt"))
}
