/*
	Day 6: Day 7: No Space Left On Device
	Link: https://adventofcode.com/2022/day/7
*/

package day07

import (
	"adventofcode2022/utils/input"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	parent *dir
	name   string
	size   int
	files  []file
	dirs   map[string]*dir
}

const internalSpace int = 70000000
const updateSpace int = 30000000

func formatInput(reader input.InputReader) (dir, error) {
	content := reader.GetInput(7)

	var currentDir *dir
	root := dir{nil, "/", 0, []file{}, make(map[string]*dir)}

	terminalOutput := strings.Split(content, "\n")

	for _, terminalLine := range terminalOutput {

		if terminalLine == "" {
			continue
		}

		terminalLineSplit := strings.Split(terminalLine, " ")
		if terminalLineSplit[0] == "$" {
			if terminalLineSplit[1] == "cd" {
				if terminalLineSplit[2] == "/" {
					currentDir = &root
				} else if terminalLineSplit[2] == ".." {
					currentDir = currentDir.parent
				} else {
					currentDir = currentDir.dirs[terminalLineSplit[2]]
				}
			} else {
				continue
			}
		} else if terminalLineSplit[0] == "dir" {
			currentDir.dirs[terminalLineSplit[1]] = &dir{currentDir, terminalLineSplit[1], 0, []file{}, make(map[string]*dir)}
		} else {
			fileName := terminalLineSplit[1]
			fileSize, err := strconv.Atoi(terminalLineSplit[0])

			if err != nil {
				return dir{nil, "", 0, []file{}, make(map[string]*dir)}, err
			}

			currentDir.files = append(currentDir.files, file{fileName, fileSize})
		}
	}

	getDirsSizes(&root)

	return root, nil
}

func getDirsSizes(directory *dir) int {

	dirSize := 0

	if len(directory.dirs) != 0 {
		for _, childDirectory := range directory.dirs {
			dirSize += getDirsSizes(childDirectory)
		}
	}

	for _, file := range directory.files {
		dirSize += file.size
	}

	directory.size = dirSize

	return dirSize
}

func getPart1Total(directory *dir, sum *int) {
	if len(directory.dirs) != 0 {
		for _, childDirectory := range directory.dirs {
			getPart1Total(childDirectory, sum)
		}
	}

	if directory.size <= 100000 {
		*sum += directory.size
	}
}

func getPart2PossibleDirs(directory *dir, possibleDirsArray *[]dir, unusedSpace int) {
	if len(directory.dirs) != 0 {
		for _, childDirectory := range directory.dirs {
			if childDirectory.size+unusedSpace >= updateSpace {
				getPart2PossibleDirs(childDirectory, possibleDirsArray, unusedSpace)
			}
		}
	}

	if directory.size+unusedSpace >= updateSpace {
		*possibleDirsArray = append(*possibleDirsArray, *directory)
	}
}

func Part1(reader input.InputReader) (int, error) {
	root, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	totalSum := 0

	getPart1Total(&root, &totalSum)

	return totalSum, nil

}

func Part2(reader input.InputReader) (int, error) {
	root, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	unusedSpace := internalSpace - root.size
	var possibleDirs []dir

	if unusedSpace >= updateSpace {
		return 0, nil
	} else {
		getPart2PossibleDirs(&root, &possibleDirs, unusedSpace)
		sort.Slice(possibleDirs, func(i, j int) bool {
			return possibleDirs[i].size < possibleDirs[j].size
		})

		return possibleDirs[0].size, nil
	}
}
