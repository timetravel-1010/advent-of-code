package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MapStruct struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type Range struct {
	lower int
	upper int
}

func getInput(fileName string) ([]Range, [][]MapStruct) {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error while opening the file!")
	}
	scanner := bufio.NewScanner(file)
	seeds := []Range{}
	scanner.Scan()
	pairs := MapToInt(strings.Fields(
		scanner.Text())[1:],
		func(x string) int {
			n, _ := strconv.Atoi(x)
			return n
		})

	// Get the seeds ranges
	for i := 0; i < len(pairs)-1; i += 2 {
		lower := pairs[i]
		seeds = append(seeds, Range{
			lower: lower,
			upper: lower + pairs[i+1],
		})
	}
	fmt.Println(seeds)
	maps := make([][]MapStruct, 7)
	isMap := false
	i := 0
	mp := MapStruct{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if isMap {
				i++
			}
			isMap = false
			continue
		}
		if !isMap {
			isMap = regexp.MustCompile(`\bmap\b`).MatchString(line)
			continue
		}
		values := MapToInt(strings.Fields(line), func(x string) int {
			n, _ := strconv.Atoi(x)
			return n
		})
		mp = MapStruct{
			destinationRangeStart: values[0],
			sourceRangeStart:      values[1],
			rangeLength:           values[2],
		}
		maps[i] = append(maps[i], mp)
	}
	fmt.Println(maps)
	defer file.Close()
	return seeds, maps
}

func getLowestLocation(fileName string) int {
	seeds, maps := getInput(fileName)
	lowestLocation := math.MaxInt

	aux := seeds
	next := aux
	for _, mp := range maps {
		for i := 0; i < len(aux); i++ {
			next[i] = ConvertRange(aux[i], mp)
		}
		aux = next
	}
	// humidity to location
	locations := []int{}
	for i := 0; i < len(aux); i++ {
		r := aux[i]
		for v := r.lower; v < r.upper+1; v++ {
			locations = append(locations, Convert(v, maps[len(maps)-1]))
		}
	}
	fmt.Println(locations)
	for _, l := range locations {
		if l < lowestLocation {
			lowestLocation = l
		}
	}
	return lowestLocation
}

func Convert(v int, mps []MapStruct) int {

	for _, mp := range mps {
		if v < (mp.sourceRangeStart+mp.rangeLength) && v >= mp.sourceRangeStart {
			acc := mp.destinationRangeStart
			for i := mp.sourceRangeStart; i < mp.sourceRangeStart+mp.rangeLength; i++ {
				if v == i {
					return acc
				}
				acc++
			}
		}
	}
	return v
}

func ConvertRange(r Range, mps []MapStruct) Range {
	v1 := r.lower
	v2 := r.upper
	newRange := r
	for _, v := range []int{v1, v2} {
		for _, mp := range mps {
			if v >= mp.sourceRangeStart && v <= (mp.sourceRangeStart+mp.rangeLength) {
				distance := v - mp.sourceRangeStart
				newRange.lower = mp.destinationRangeStart + distance
				break
			}
		}
	}
	return r
}

func MapToInt(arr []string, f func(string) int) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i] = f(v)
	}
	return newArr
}

func main() {
	fmt.Println(getLowestLocation("test.txt"))
	// getInput("test.txt")
}
