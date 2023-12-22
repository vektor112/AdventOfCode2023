package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vektor112/AdventOfCode2023/internal/day1"
)

const (
  DAY_1 = "1"
)

type ChallengeRunnerFunction func(scanner []string) string

var challenges = map[string]ChallengeRunnerFunction{
  DAY_1: day1.RunCalculation,
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getInput() string {
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter the day of challenge (1, 2, 3 ...):")

  var input string

	for scanner.Scan() {
		input = scanner.Text()
    break
	}
	err := scanner.Err()
  check(err)

  return input
}

func collectLines(input string) []string {
  currentDir, err := os.Getwd()
  check(err)
  file, err := os.Open(currentDir + "/example/" + input)
  check(err)
  defer file.Close()
  
  var lines []string
  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

  return lines
}

func main() {
  input := getInput()
  
  foundChallenge := challenges[input]
  
  if foundChallenge == nil {
    fmt.Println("This day is not supported.")
    return
  } else {
    fmt.Println("Load the example/", input)
    lines := collectLines(input)
    fmt.Println("Running the", input, "day challenge...")
    fmt.Println("The result is:", foundChallenge(lines))
  }
}
