package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var limits = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	readFile, _ := os.Open("input.txt")
	// readFile, _ := os.Open("test_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		gameInfo := strings.Split(line, ":")

		gameNumber, _ := strconv.Atoi(strings.TrimPrefix(gameInfo[0], "Game "))
		turns := strings.Split(strings.Trim(gameInfo[1], " "), ";")
		validGame := true
		for _, turn := range turns {
			sets := strings.Split(strings.Trim(turn, " "), ",")
			for _, set := range sets {
				splitTurn := strings.Split(strings.Trim(set, " "), " ")
				//0: amount
				amount, _ := strconv.Atoi(splitTurn[0])
				//1: color
				color := splitTurn[1]
				if limits[color] < amount {
					validGame = false
				}
			}

		}
		if validGame {
			total += gameNumber
		}
		// fmt.Println("Current Total: ", total)
	}

	readFile.Close()
	fmt.Println("Total: ", total)
}
