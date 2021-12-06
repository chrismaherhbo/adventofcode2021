package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type intList []int

func createIntList(file string) intList {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(body), "\n")

	var number int
	var list intList

	for _, line := range lines {
		number, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}

		list = append(list, number)
	}

	return list
}

func (list *intList) countIncreases() int {
	var increases int

	for i := 0; i < len(*list); i++ {
		if i > 0 && (*list)[i] > (*list)[i-1] {
			increases++
		}
	}

	return increases
}

func (list *intList) countWindowIncreases() int {
	var increases int
	var windows []int

	for i := 0; i < len(*list); i++ {
		if i > 1 {
			windows = append(windows, (*list)[i]+(*list)[i-1]+(*list)[i-2])
			if i > 2 && windows[i-2] > windows[i-3] {
				increases++
			}
		}
	}

	return increases
}

func main() {
	list := createIntList("day1.txt")
	fmt.Printf("Part 1 answer: %d\n", list.countIncreases())
	fmt.Printf("Part 2 answer: %d\n", list.countWindowIncreases())
}
