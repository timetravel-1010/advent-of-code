package main

import (
	"fmt"
	"os"
	"strings"
)

type End struct {
	pos struct {
		y int
		x int
	}
	paths [][2]int
}

func (e *End) String() string {
	return fmt.Sprintf("&{pos:{y:%d, x:%d}, paths:%v", e.pos.y, e.pos.x, e.paths)
}

func solutionPart1(fileName string) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error reading the file")
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	tMap := make([][]rune, len(lines))

	for i, line := range lines {
		tMap[i] = []rune(line)
	}

	ends := make(map[int]map[int]map[int]map[int]bool)
	endPoss := []*End{}

	for y, row := range tMap {
		ends[y] = make(map[int]map[int]map[int]bool)
		for x, c := range row {
			if c == '9' {
				endPoss = append(endPoss, &End{pos: struct {
					y int
					x int
				}{y, x},
					paths: [][2]int{{y, x}},
				})
				ends[y][x] = make(map[int]map[int]bool)
			}
		}
	}
	scores := make(map[int]map[int]int)
	locs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for currentHeight := 57; currentHeight > 48; currentHeight-- {
		for _, end := range endPoss {
			paths := [][2]int{}
			for _, pos := range end.paths {
				if pos[0] > 0 {
					upY := pos[0] + locs[0][0]
					upX := pos[1] + locs[0][1]
					if up := int(tMap[upY][upX]); up == (currentHeight - 1) {
						if up == 48 && currentHeight == 49 &&
							!ends[end.pos.y][end.pos.x][upY][upX] {
							if _, ok := ends[end.pos.y][end.pos.x][upY]; !ok {
								ends[end.pos.y][end.pos.x][upY] = make(map[int]bool)
							}
							ends[end.pos.y][end.pos.x][upY][upX] = true
							if _, ok := scores[upY]; !ok {
								scores[upY] = make(map[int]int)
							}
							scores[upY][upX]++
						} else {
							paths = append(paths, [2]int{upY, upX})
						}
					}
				}

				if pos[1] < len(tMap[0])-1 {
					rightY := pos[0] + locs[1][0]
					rightX := pos[1] + locs[1][1]
					if right := int(tMap[rightY][rightX]); right == (currentHeight - 1) {
						if right == 48 && currentHeight == 49 &&
							!ends[end.pos.y][end.pos.x][rightY][rightX] {
							if _, ok := ends[end.pos.y][end.pos.x][rightY]; !ok {
								ends[end.pos.y][end.pos.x][rightY] = make(map[int]bool)
							}
							ends[end.pos.y][end.pos.x][rightY][rightX] = true
							if _, ok := scores[rightY]; !ok {
								scores[rightY] = make(map[int]int)
							}
							scores[rightY][rightX]++
						} else {
							paths = append(paths, [2]int{rightY, rightX})
						}
					}
				}

				if pos[0] < len(tMap)-1 {
					downY := pos[0] + locs[2][0]
					downX := pos[1] + locs[2][1]
					if down := int(tMap[downY][downX]); down == (currentHeight - 1) {
						if down == 48 && currentHeight == 49 &&
							!ends[end.pos.y][end.pos.x][downY][downX] {
							if _, ok := ends[end.pos.y][end.pos.x][downY]; !ok {
								ends[end.pos.y][end.pos.x][downY] = make(map[int]bool)
							}
							ends[end.pos.y][end.pos.x][downY][downX] = true
							if _, ok := scores[downY]; !ok {
								scores[downY] = make(map[int]int)
							}
							scores[downY][downX]++
						} else {
							paths = append(paths, [2]int{downY, downX})
						}
					}
				}

				if pos[1] > 0 {
					leftY := pos[0] + locs[3][0]
					leftX := pos[1] + locs[3][1]
					if left := int(tMap[leftY][leftX]); left == (currentHeight - 1) {
						if left == 48 && currentHeight == 49 &&
							!ends[end.pos.y][end.pos.x][leftY][leftX] {
							if _, ok := ends[end.pos.y][end.pos.x][leftY]; !ok {
								ends[end.pos.y][end.pos.x][leftY] = make(map[int]bool)
							}
							ends[end.pos.y][end.pos.x][leftY][leftX] = true
							if _, ok := scores[leftY]; !ok {
								scores[leftY] = make(map[int]int)
							}
							scores[leftY][leftX]++
						} else {
							paths = append(paths, [2]int{leftY, leftX})
						}
					}
				}

			} // end pos
			end.paths = paths
		} // endPoss
	}
	fmt.Println("trailheads:", len(scores))
	fmt.Println("scores:", scores)
	sum := 0
	for _, v := range scores {
		for _, score := range v {
			sum += score
		}
	}
	return sum
}

func solutionPart2(fileName string) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error reading the file")
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	tMap := make([][]rune, len(lines))

	for i, line := range lines {
		tMap[i] = []rune(line)
	}

	ends := make(map[int]map[int]map[int]map[int]int)
	endPoss := []*End{}

	for y, row := range tMap {
		ends[y] = make(map[int]map[int]map[int]int)
		for x, c := range row {
			if c == '9' {
				endPoss = append(endPoss, &End{pos: struct {
					y int
					x int
				}{y, x},
					paths: [][2]int{{y, x}},
				})
				ends[y][x] = make(map[int]map[int]int)
			}
		}
	}
	scores := make(map[int]map[int]int)
	locs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for currentHeight := 57; currentHeight > 48; currentHeight-- {
		for _, end := range endPoss {
			paths := [][2]int{}
			for _, pos := range end.paths {
				if pos[0] > 0 {
					upY := pos[0] + locs[0][0]
					upX := pos[1] + locs[0][1]
					if up := int(tMap[upY][upX]); up == (currentHeight - 1) {
						if up == 48 && currentHeight == 49 {
							//!ends[end.pos.y][end.pos.x][upY][upX] {
							if _, ok := ends[end.pos.y][end.pos.x][upY]; !ok {
								ends[end.pos.y][end.pos.x][upY] = make(map[int]int)
							}
							ends[end.pos.y][end.pos.x][upY][upX]++
							if _, ok := scores[upY]; !ok {
								scores[upY] = make(map[int]int)
							}
							scores[upY][upX]++
						} else {
							paths = append(paths, [2]int{upY, upX})
						}
					}
				}

				if pos[1] < len(tMap[0])-1 {
					rightY := pos[0] + locs[1][0]
					rightX := pos[1] + locs[1][1]
					if right := int(tMap[rightY][rightX]); right == (currentHeight - 1) {
						if right == 48 && currentHeight == 49 {
							//!ends[end.pos.y][end.pos.x][rightY][rightX] {
							if _, ok := ends[end.pos.y][end.pos.x][rightY]; !ok {
								ends[end.pos.y][end.pos.x][rightY] = make(map[int]int)
							}
							ends[end.pos.y][end.pos.x][rightY][rightX]++
							if _, ok := scores[rightY]; !ok {
								scores[rightY] = make(map[int]int)
							}
							scores[rightY][rightX]++
						} else {
							paths = append(paths, [2]int{rightY, rightX})
						}
					}
				}

				if pos[0] < len(tMap)-1 {
					downY := pos[0] + locs[2][0]
					downX := pos[1] + locs[2][1]
					if down := int(tMap[downY][downX]); down == (currentHeight - 1) {
						if down == 48 && currentHeight == 49 {
							//!ends[end.pos.y][end.pos.x][downY][downX] {
							if _, ok := ends[end.pos.y][end.pos.x][downY]; !ok {
								ends[end.pos.y][end.pos.x][downY] = make(map[int]int)
							}
							ends[end.pos.y][end.pos.x][downY][downX]++
							if _, ok := scores[downY]; !ok {
								scores[downY] = make(map[int]int)
							}
							scores[downY][downX]++
						} else {
							paths = append(paths, [2]int{downY, downX})
						}
					}
				}

				if pos[1] > 0 {
					leftY := pos[0] + locs[3][0]
					leftX := pos[1] + locs[3][1]
					if left := int(tMap[leftY][leftX]); left == (currentHeight - 1) {
						if left == 48 && currentHeight == 49 {
							//!ends[end.pos.y][end.pos.x][leftY][leftX] {
							if _, ok := ends[end.pos.y][end.pos.x][leftY]; !ok {
								ends[end.pos.y][end.pos.x][leftY] = make(map[int]int)
							}
							ends[end.pos.y][end.pos.x][leftY][leftX]++
							if _, ok := scores[leftY]; !ok {
								scores[leftY] = make(map[int]int)
							}
							scores[leftY][leftX]++
						} else {
							paths = append(paths, [2]int{leftY, leftX})
						}
					}
				}

			} // end pos
			end.paths = paths
		} // endPoss
	}
	sum := 0
	for _, v := range scores {
		for _, score := range v {
			sum += score
		}
	}
	return sum
}
func main() {
	//fmt.Println("sol1:", solutionPart1("input.txt"))
	fmt.Println("sol2:", solutionPart2("input.txt"))
}
