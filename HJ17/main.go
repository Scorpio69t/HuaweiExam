package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInstructions(input string) (int, int) {
	x, y := 0, 0

	instructions := strings.Split(input, ";")

	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}

		if ok, direction, distance := isValidInstruction(instruction); !ok {
			continue
		} else {
			switch direction {
			case 'A':
				x -= distance
			case 'D':
				x += distance
			case 'W':
				y += distance
			case 'S':
				y -= distance
			}
		}
	}

	return x, y
}

func isValidInstruction(instruction string) (bool, byte, int) {
	direction := byte(' ')
	distance := 0

	if len(instruction) < 2 || len(instruction) > 4 {
		return false, direction, distance
	}

	direction = instruction[0]
	if direction != 'A' && direction != 'D' &&
		direction != 'S' && direction != 'W' {
		return false, direction, distance
	}

	distanceStr := instruction[1:]

	for _, char := range distanceStr {
		if char < '0' || char > '9' {
			return false, direction, distance
		}
	}

	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return false, direction, distance
	}

	if distance <= 0 || distance >= 100 {
		return false, direction, distance
	}

	return true, direction, distance
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			continue
		}

		x, y := processInstructions(input)
		fmt.Printf("%d,%d", x, y)
	}
}
