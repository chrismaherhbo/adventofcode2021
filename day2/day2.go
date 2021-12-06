package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Up
	Down
)

type movement struct {
	direction Direction
	magnitude int
}

type coordinates struct {
	x   int
	y   int
	aim int
}

func readMovements(file string) []movement {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(body), "\n")

	var move *movement
	var moves []movement
	var number int

	for _, line := range lines {
		move = new(movement)

		fields := strings.Fields(line)
		number, err = strconv.Atoi(string(fields[1]))
		if err != nil {
			fmt.Println(err)
		}
		move.magnitude = number

		switch fields[0] {
		case "forward":
			move.direction = Forward
		case "down":
			move.direction = Down
		case "up":
			move.direction = Up
		}

		moves = append(moves, *move)
	}

	return moves
}

func calculatePosition(moves []movement) coordinates {
	var coords coordinates
	for _, move := range moves {
		switch move.direction {
		case Forward:
			coords.x += move.magnitude
			coords.y += coords.aim * move.magnitude
		case Up:
			coords.aim -= move.magnitude
		case Down:
			coords.aim += move.magnitude
		}
	}

	return coords
}

func main() {
	movements := readMovements("day2.txt")
	coordinates := calculatePosition(movements)
	fmt.Println(coordinates.x * coordinates.y)
}
