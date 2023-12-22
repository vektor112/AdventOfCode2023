package day2

import (
	"strconv"
	"strings"

	"github.com/vektor112/AdventOfCode2023/internal/utils"
)

const (
	RED_COLOR   = "red"
	BLUE_COLOR  = "blue"
	GREEN_COLOR = "green"

	RED_LIMIT   = 12
	BLUE_LIMIT  = 14
	GREEN_LIMIT = 13
)

func addOrDefault(color string, count int, cubeColorCounter *map[string]int) {
	if _, ok := (*cubeColorCounter)[color]; !ok {
		(*cubeColorCounter)[color] = 0
	}

	if (*cubeColorCounter)[color] < count {
		(*cubeColorCounter)[color] = count
	}
}

func trimStrings(items *[]string) {
	for index := 0; index < len(*items); index++ {
		(*items)[index] = strings.TrimLeft((*items)[index], " ")
	}
}

func RunCalculation(lines []string) string {

	var possibleGameIds []int
	var powers []int

	for _, line := range lines {
		mainParts := strings.Split(line, ":")

		gameId, _ := strconv.Atoi(strings.Split(mainParts[0], " ")[1])
		sets := strings.Split(mainParts[1], ";")

		cubeColorCounter := make(map[string]int)

		for _, set := range sets {
			cubeColorCounts := strings.Split(set, ",")
			trimStrings(&cubeColorCounts)

			for _, cubeColorCount := range cubeColorCounts {
				splittedcubeColorCount := strings.Split(cubeColorCount, " ")
				count, _ := strconv.Atoi(splittedcubeColorCount[0])
				color := splittedcubeColorCount[1]

				addOrDefault(color, count, &cubeColorCounter)
			}
		}

		if cubeColorCounter[RED_COLOR] <= RED_LIMIT &&
			cubeColorCounter[BLUE_COLOR] <= BLUE_LIMIT &&
			cubeColorCounter[GREEN_COLOR] <= GREEN_LIMIT {
			possibleGameIds = append(possibleGameIds, gameId)
		}

		powers = append(
			powers,
			cubeColorCounter[RED_COLOR]*cubeColorCounter[BLUE_COLOR]*cubeColorCounter[GREEN_COLOR],
		)
	}

	// Part 1
	// return strconv.Itoa(utils.SumIntegers(possibleGameIds))

	// Part 2
	return strconv.Itoa(utils.SumIntegers(powers))
}
