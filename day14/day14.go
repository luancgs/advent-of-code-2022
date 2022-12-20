/*
	Day 14: Regolith Reservoir
	Link: https://adventofcode.com/2022/day/14
*/

package day14

import (
	"adventofcode2022/utils/input"
	"strconv"
	"strings"
)

type Element interface{}

type Air struct {
	x, y int
}

type Rock struct {
	x, y int
}

type Sand struct {
	x, y int
}

func (s *Sand) move(area [][]Element) bool {
	_, isAirDown := area[s.y+1][s.x].(Air)
	_, isAirLeft := area[s.y+1][s.x-1].(Air)
	_, isAirRight := area[s.y+1][s.x+1].(Air)

	if isAirDown {
		s.y++
	} else if isAirLeft {
		s.x--
		s.y++
	} else if isAirRight {
		s.x++
		s.y++
	} else {
		return false
	}
	return true
}

func rangePositions(start, end int) []int {
	var output []int

	if start < end {
		for i := start; i <= end; i++ {
			output = append(output, i)
		}
	} else {
		for i := end; i <= start; i++ {
			output = append(output, i)
		}
	}

	return output
}

func getAllRocks(rockPaths [][]Rock) []Rock {
	var allRocks []Rock

	for _, rockPath := range rockPaths {
		for i := 0; i < len(rockPath)-1; i++ {
			if rockPath[i].x == rockPath[i+1].x {
				for _, y := range rangePositions(rockPath[i].y, rockPath[i+1].y) {
					allRocks = append(allRocks, Rock{rockPath[i].x, y})
				}
			} else {
				for _, x := range rangePositions(rockPath[i].x, rockPath[i+1].x) {
					allRocks = append(allRocks, Rock{x, rockPath[i].y})
				}
			}
		}
	}

	return allRocks
}

func createArea(allRocks []Rock) [][]Element {
	var greatestX, greatestY int = 0, 0

	for _, rock := range allRocks {
		if rock.x > greatestX {
			greatestX = rock.x
		}
		if rock.y > greatestY {
			greatestY = rock.y
		}
	}

	var area [][]Element

	for y := 0; y <= greatestY; y++ {
		var areaLine []Element
		for x := 0; x <= greatestX; x++ {
			areaLine = append(areaLine, Air{x, y})
		}
		area = append(area, areaLine)
	}

	for y, areaLine := range area {
		for x := range areaLine {
			for _, rock := range allRocks {
				if rock.x == x && rock.y == y {
					area[y][x] = Rock{x, y}
				}
			}
		}
	}

	return area
}

func formatInput() ([][]Element, error) {
	content := input.GetInput(14)

	rockPathsInput := strings.Split(content, "\n")
	var rockPaths [][]Rock

	for _, rockPath := range rockPathsInput {
		if rockPath == "" {
			continue
		}

		rockPathSplit := strings.Split(rockPath, " -> ")
		var rockPath []Rock

		for _, rockPosition := range rockPathSplit {
			pos := strings.Split(rockPosition, ",")

			x, err := strconv.Atoi(pos[0])

			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(pos[1])

			if err != nil {
				return nil, err
			}

			rockPath = append(rockPath, Rock{x, y})
		}

		rockPaths = append(rockPaths, rockPath)
	}

	allRocks := getAllRocks(rockPaths)
	area := createArea(allRocks)

	return area, nil
}

func Part1() (int, error) {

	cave, err := formatInput()

	if err != nil {
		return 0, err
	}

	sandOrigin := []int{500, 0}
	sand := Sand{sandOrigin[0], sandOrigin[1]}
	intoTheAbyss := false
	countSand := 0

	for !intoTheAbyss {
		sand = Sand{sandOrigin[0], sandOrigin[1]}
		var sandSettled bool = false

		for !sandSettled {
			sandSettled = !sand.move(cave)

			if sand.y == len(cave)-1 || sand.x == len(cave[0])-1 {
				intoTheAbyss = true
				break
			}
		}

		if !intoTheAbyss {
			cave[sand.y][sand.x] = sand
			countSand++
		}
	}

	return countSand, nil
}

func Part2() (int, error) {
	cave, err := formatInput()

	if err != nil {
		return 0, err
	}

	for i := 1; i <= 2; i++ {
		var caveLine []Element
		for j := 0; j < len(cave[0]); j++ {
			if i == 1 {
				caveLine = append(caveLine, Air{j, len(cave) + i})
			} else {
				caveLine = append(caveLine, Rock{j, len(cave) + i})
			}
		}
		cave = append(cave, caveLine)
	}

	sandOrigin := []int{500, 0}
	sand := Sand{sandOrigin[0], sandOrigin[1]}
	filledCave := false
	countSand := 0

	for !filledCave {
		sand = Sand{sandOrigin[0], sandOrigin[1]}
		var sandSettled bool = false

		for !sandSettled {
			sandSettled = !sand.move(cave)

			if sand.x == len(cave[0])-1 {

				for y, caveLine := range cave {
					var newCaveLine []Element
					if y == len(cave)-1 {
						newCaveLine = append(caveLine, Rock{len(caveLine), y})
					} else {
						newCaveLine = append(caveLine, Air{len(caveLine), y})
					}
					cave[y] = newCaveLine
				}
			}

			if sand.x == sandOrigin[0] && sand.y == sandOrigin[1] {
				filledCave = true
				break
			}
		}

		if !filledCave {
			cave[sand.y][sand.x] = sand
		}
		countSand++
	}

	return countSand, nil
}
