package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cylonsam/advent-of-code/2015/auxlib"
)

type Command string

const (
	TurnOn  Command = "turn on"
	TurnOff Command = "turn off"
	Toggle  Command = "toggle"
)

func main() {
	filePath := "/home/samlas/workspace/advent-of-code/2015/06/input.txt"

	inputs := auxlib.ReadLinesTextFile(filePath)
	SolveDay6(inputs, auxlib.Part2)
}

func SolveDay6(inputs []string, challengePart auxlib.ChallengePart) int {
	if challengePart == auxlib.Part1 {
		result := solvePart1(inputs)
		fmt.Println(result)
		return result
	} else {
		result := solvePart2(inputs)
		fmt.Println(result)
		return result
	}

}

func solvePart1(inputs []string) int {
	var command Command
	grid := buildGrid()

	for i := range len(inputs) {
		rawInput := inputs[i]
		cleanInput := ""
		if strings.Contains(rawInput, string(TurnOn)) {
			command = TurnOn
			cleanInput, _ = strings.CutPrefix(rawInput, string(TurnOn))
		} else if strings.Contains(rawInput, string(TurnOff)) {
			command = TurnOff
			cleanInput, _ = strings.CutPrefix(rawInput, string(TurnOff))
		} else {
			command = Toggle
			cleanInput, _ = strings.CutPrefix(rawInput, string(Toggle))
		}
		// fmt.Println(command)

		inputArgs := strings.Split(strings.TrimSpace(cleanInput), " ")
		// fmt.Println(inputArgs)
		arg1, arg2 := strings.Split(inputArgs[0], ","), strings.Split(inputArgs[2], ",")
		startRow, _ := strconv.Atoi(arg1[0])
		startColumn, _ := strconv.Atoi(arg1[1])
		endRow, _ := strconv.Atoi(arg2[0])
		endColumn, _ := strconv.Atoi(arg2[1])
		// fmt.Println(arg1, arg2)

		switch command {
		case TurnOn:
			for i := startRow; i <= endRow; i++ {
				for j := startColumn; j <= endColumn; j++ {
					grid[i][j] = 1
				}
			}
		case TurnOff:
			for i := startRow; i <= endRow; i++ {
				for j := startColumn; j <= endColumn; j++ {
					grid[i][j] = 0
				}
			}
		case Toggle:
			for i := startRow; i <= endRow; i++ {
				for j := startColumn; j <= endColumn; j++ {
					if grid[i][j] == 0 {
						grid[i][j] = 1
					} else {
						grid[i][j] = 0
					}
				}
			}
		}

	}
	return countLitLights(grid)
}

func solvePart2(inputs []string) int {
	var command Command
	grid := buildGrid()

	for i := range len(inputs) {
		rawInput := inputs[i]
		cleanInput := ""
		if strings.Contains(rawInput, string(TurnOn)) {
			command = TurnOn
			cleanInput, _ = strings.CutPrefix(rawInput, string(TurnOn))
		} else if strings.Contains(rawInput, string(TurnOff)) {
			command = TurnOff
			cleanInput, _ = strings.CutPrefix(rawInput, string(TurnOff))
		} else {
			command = Toggle
			cleanInput, _ = strings.CutPrefix(rawInput, string(Toggle))
		}
		// fmt.Println(command)

		inputArgs := strings.Split(strings.TrimSpace(cleanInput), " ")
		// fmt.Println(inputArgs)
		arg1, arg2 := strings.Split(inputArgs[0], ","), strings.Split(inputArgs[2], ",")
		startRow, _ := strconv.Atoi(arg1[0])
		startColumn, _ := strconv.Atoi(arg1[1])
		endRow, _ := strconv.Atoi(arg2[0])
		endColumn, _ := strconv.Atoi(arg2[1])
		// fmt.Println(arg1, arg2)

		switch command {
		case TurnOn:
			for i := startRow; i <= endRow; i++ {
				for j := startColumn; j <= endColumn; j++ {
					grid[i][j]++
				}
			}
		case TurnOff:
			for i := startRow; i <= endRow; i++ {
				for j := startColumn; j <= endColumn; j++ {
					if grid[i][j] <= 0 {
						grid[i][j] = 0
					} else {
						grid[i][j]--
					}

				}
			}
		case Toggle:
			for i := startRow; i <= endRow; i++ {
				for j := startColumn; j <= endColumn; j++ {
					grid[i][j] += 2
				}
			}
		}

	}
	return countBrightness(grid)
}

func buildGrid() [1000][1000]int {
	var grid [1000][1000]int
	for i := range 1000 {
		for j := range 1000 {
			grid[i][j] = 0
		}
	}
	return grid
}

func countLitLights(grid [1000][1000]int) int {
	litLights := 0
	for i := range 1000 {
		for j := range 1000 {
			if grid[i][j] == 1 {
				litLights++
			}
		}
	}
	return litLights
}

func countBrightness(grid [1000][1000]int) int {
	brightness := 0
	for i := range 1000 {
		for j := range 1000 {
			brightness += grid[i][j]
		}
	}
	return brightness
}
