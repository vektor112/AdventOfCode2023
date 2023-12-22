package day1

import (
  "strconv"
  "strings"
  "unicode"

  "github.com/vektor112/AdventOfCode2023/internal/utils"
)

const (
  EMPTY = ""

  ONE_STRING = "one"
  TWO_STRING = "two"
  THREE_STRING = "three"
  FOUR_STRING = "four"
  FIVE_STRING = "five"
  SIX_STRING = "six"
  SEVEN_STRING = "seven"
  EIGHT_STRING = "eight"
  NINE_STRING = "nine"

  ONE_NUMBER = "1"
  TWO_NUMBER = "2"
  THREE_NUMBER = "3"
  FOUR_NUMBER = "4"
  FIVE_NUMBER = "5"
  SIX_NUMBER = "6"
  SEVEN_NUMBER = "7"
  EIGHT_NUMBER = "8"
  NINE_NUMBER = "9"
)

var numberWords = []string{ 
  ONE_STRING, 
  TWO_STRING, 
  THREE_STRING, 
  FOUR_STRING, 
  FIVE_STRING, 
  SIX_STRING, 
  SEVEN_STRING, 
  EIGHT_STRING, 
  NINE_STRING,
}

var wordToNumberString = map[string]string{
  ONE_STRING: ONE_NUMBER,
  TWO_STRING: TWO_NUMBER,
  THREE_STRING: THREE_NUMBER,
  FOUR_STRING: FOUR_NUMBER,
  FIVE_STRING: FIVE_NUMBER,
  SIX_STRING: SIX_NUMBER,
  SEVEN_STRING: SEVEN_NUMBER,
  EIGHT_STRING: EIGHT_NUMBER,
  NINE_STRING: NINE_NUMBER,
}

func getWordKeyFromCharacters(index, lengthOfNumberWord int, characters *[]rune) string {
  var upcomingRunes []string

  for i := index; i < index + lengthOfNumberWord; i++ {
    upcomingRunes = append(upcomingRunes, string((*characters)[i]))
  }

  return strings.Join(upcomingRunes, "")
}

func tryToFindNumberString(index int, lengthOfLine int, characters *[]rune) (string, bool) {
  for _, numberWord := range numberWords {
    lengthOfNumberWord := len(numberWord)

    if lengthOfNumberWord + index - 1 < lengthOfLine {
      wordKey := getWordKeyFromCharacters(index, lengthOfNumberWord, characters)
      numberString := wordToNumberString[wordKey]

      if wordKey == numberWord && numberString != EMPTY {
        return numberString, true
      }
    }
  }

  return "", false
}

func RunCalculation(lines []string) string {
  var numbers []int 

  for _, line := range lines {
    var firstDigit, lastDigit string

    characters := []rune(line)
    lengthOfLine := len(characters)

    for index := 0; index < lengthOfLine; index++ {
      char := characters[index]
      var digit string
      isDigitFound := false 
     
      if unicode.IsDigit(char) {
        digit = string(char)
        isDigitFound = true
      } else {
        numberString, numberStringfound := tryToFindNumberString(index, lengthOfLine, &characters)
        if numberStringfound {
          digit = numberString
          isDigitFound = true
        } 
      }

      if (isDigitFound) {
        if firstDigit == EMPTY {
          firstDigit = digit
        }

        lastDigit = digit
      }
    }

    joinedDigits := firstDigit + lastDigit
    if joinedDigits != EMPTY {
      number, _ := strconv.Atoi(joinedDigits) 
      numbers = append(numbers, number)
    }
  }

  return strconv.Itoa(utils.SumIntegers(numbers))
}
