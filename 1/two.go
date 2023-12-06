package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// readFile, _ := os.Open("test_input.txt")
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var firstDigit rune
		var lastDigit rune
		for position, c := range line {
			c = isDigitWord(c, position, line, "")
			if unicode.IsDigit(c) {
				if reflect.ValueOf(firstDigit).IsZero() {
					firstDigit = c
					lastDigit = c
				} else {
					lastDigit = c
				}
			}
		}

		number, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		fmt.Println("Number: ", number)
		total += number
		fmt.Println("Current Total: ", total)
	}
	readFile.Close()
	fmt.Println(total)
}

func isDigitWord(r rune, position int, line string, currentString string) (retValue rune) {
	var myMap = map[rune]string{
		'1': "one",
		'2': "two",
		'3': "three",
		'4': "four",
		'5': "five",
		'6': "six",
		'7': "seven",
		'8': "eight",
		'9': "nine"}

	currentString = currentString + string(r)
	for value, word := range myMap {
		if currentString == word {
			return value

		} else if strings.HasPrefix(word, currentString) && position+1 < len(line) {
			return isDigitWord(rune(line[position+1]), position+1, line, currentString)

		}
	}

	return r
}
