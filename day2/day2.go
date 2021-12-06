package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	Forward = iota
	Up
	Down
)

type Movement struct {
	Direction int
	Magnitude int
}

type Coordinates struct {
	x int
	y int
}

func readMovements(file string) []Movement {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(body), "\n")

	var movement *Movement
	var movements []Movement
	var number int

	for _, line := range lines {
		movement = new(Movement)

		fields := strings.Fields(line)
		number, err = strconv.Atoi(string(fields[1]))
		if err != nil {
			fmt.Println(err)
		}
		movement.Magnitude = number

		switch fields[0] {
		case "forward":
			movement.Direction = Forward
		case "down":
			movement.Direction = Down
		case "up":
			movement.Direction = Up
		}

		movements = append(movements, *movement)
	}

	return movements
}

func calculatePosition(movements []Movement) Coordinates {
	var coordinates Coordinates
	for _, movement := range movements {
		switch movement.Direction {
		case 0:
			coordinates.x += movement.Magnitude
		case 1:
			coordinates.y -= movement.Magnitude
		case 2:
			coordinates.y += movement.Magnitude
		}
	}

	return coordinates
}

func main() {
	movements := readMovements("day2.txt")
	coordinates := calculatePosition(movements)
	fmt.Println(coordinates.x * coordinates.y)
}
