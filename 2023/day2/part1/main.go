package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dries1234/aoc/2023/day2"
	"github.com/Dries1234/aoc/util"
)

var colorMapping = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	input := util.TrimInput(day2.Input)
	lines := strings.Split(input, "\n")

	var sum int

line:
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		draws := strings.Split(line, ";")
		for _, draw := range draws {
			counts := strings.Split(draw, ",")
			for _, count := range counts {
				reg, err := regexp.Compile("[^0-9]+")
				if err != nil {
					panic(err)
				}
				
				number, err := strconv.Atoi(reg.ReplaceAllString(count, ""))
				if err != nil {
					panic(err)
				}

				for _, color := range []string{"red", "green", "blue"} {
					if strings.Contains(count, color) {
						if number > colorMapping[color] {
							continue line
						}
					}
				}

			}

		}
		sum += i + 1
	}

	fmt.Println(sum)

}
