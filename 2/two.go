package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, _ := os.Open("input.txt")
	// readFile, _ := os.Open("test_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		gameInfo := strings.Split(line, ":")

		// gameNumber, _ := strconv.Atoi(strings.TrimPrefix(gameInfo[0], "Game "))
		turns := strings.Split(strings.Trim(gameInfo[1], " "), ";")
		var gameLimits = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, turn := range turns {
			sets := strings.Split(strings.Trim(turn, " "), ",")
			for _, set := range sets {
				splitTurn := strings.Split(strings.Trim(set, " "), " ")
				//0: amount
				amount, _ := strconv.Atoi(splitTurn[0])
				//1: color
				color := splitTurn[1]
				if gameLimits[color] < amount {
					gameLimits[color] = amount
				}
			}

		}
		var power = 1
		for _, v := range gameLimits {
			power *= v
		}
		total += power
	}
	readFile.Close()
	fmt.Println("Total: ", total)
}
