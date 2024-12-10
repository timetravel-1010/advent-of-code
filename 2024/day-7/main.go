package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solutionPart1(fileName string) int {
	result := 0
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		vals := strings.Split(scanner.Text(), ":")
		testVal, _ := strconv.Atoi(vals[0])
		nums := strings.Fields(vals[1])

		l, _ := strconv.Atoi(nums[0])
		pos := []int{l}
		lastRhs := nums[len(nums)-1]
	outer:
		for _, rhStr := range nums[1:] {
			rhs, _ := strconv.Atoi(rhStr)
			length := len(pos)
			for i := 0; i < length; i++ {
				lhs := pos[i]
				if tempAdd := lhs + rhs; tempAdd < testVal {
					pos[i] = tempAdd
				} else if rhStr == lastRhs {
					if tempAdd == testVal {
						result += tempAdd
						break outer
					}
				} else {
					pos[i] = tempAdd
					continue // check this, maybe a multiplication doesnt change
				}

				if tempMul := lhs * rhs; tempMul < testVal {
					pos = append(pos, tempMul)
				} else if rhStr == lastRhs {
					if tempMul == testVal {
						result += tempMul
						break outer
					}
				} else {
					pos = append(pos, tempMul)
				}
			}
		}
	}
	return result
}

func solutionPart2(fileName string) int {
	result := 0
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		vals := strings.Split(scanner.Text(), ":")
		testVal, _ := strconv.Atoi(vals[0])
		nums := strings.Fields(vals[1])

		l, _ := strconv.Atoi(nums[0])
		pos := []int{l}
		lastRhs := nums[len(nums)-1]
	outer:
		for _, rhStr := range nums[1:] {
			rhs, _ := strconv.Atoi(rhStr)
			length := len(pos)
			for i := 0; i < length; i++ {
				lhs := pos[i]

				if tempAdd := lhs + rhs; tempAdd < testVal {
					pos[i] = tempAdd
				} else if rhStr == lastRhs {
					if tempAdd == testVal {
						result += tempAdd
						break outer
					}
				} else {
					pos[i] = tempAdd
					continue // check this, maybe a multiplication doesnt change
				}

				if tempMul := lhs * rhs; tempMul < testVal {
					pos = append(pos, tempMul)
				} else if rhStr == lastRhs {
					if tempMul == testVal {
						result += tempMul
						break outer
					}
				} else {
					pos = append(pos, tempMul)
				}

				con := strconv.Itoa(lhs) + rhStr
				tempCon, _ := strconv.Atoi(con)
				if tempCon < testVal {
					pos = append(pos, tempCon)
				} else if rhStr == lastRhs {
					if tempCon == testVal {
						result += tempCon
						break outer
					}
				} else {
					pos = append(pos, tempCon)
				}
			}
		}
	}
	return result
}
func main() {
	fmt.Printf("sol1: %d\n", solutionPart1("test.txt"))
	fmt.Printf("sol2: %d\n", solutionPart2("input.txt"))
}
