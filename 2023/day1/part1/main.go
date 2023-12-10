package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dries1234/aoc/2023/day1"
	"github.com/Dries1234/aoc/util"
)


func main() {
	input := util.TrimInput(day1.Input)

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")
	var sum int
	for _, line := range lines {
		number := reg.ReplaceAllString(line, "")

		if len(number) < 2 {
			number += number
		}
		if len(number) > 2 {
			number = string(number[0]) + string(number[len(number)-1])
		}

		val, err := strconv.Atoi(number)
		if err != nil {
			panic("Can't convert to number")
		}
		sum += val
	}

	fmt.Println(sum)
}
