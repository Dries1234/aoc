package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dries1234/aoc/2023/day2"
	"github.com/Dries1234/aoc/util"
)

func main() {
	input := util.TrimInput(day2.Input)
	lines := strings.Split(input, "\n")

	var powersum int

	for _, line := range lines {
		minimum := make(map[string]int)

		line = strings.Split(line, ":")[1]
		draws := strings.Split(line, ";")

		for _, draw := range draws {
			counts := strings.Split(draw, ",")

		count:
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
						if _, ok := minimum[color]; !ok || number > minimum[color] {
							minimum[color] = number
						}
						continue count
					}
				}

			}
		}
		powersum += minimum["red"] * minimum["blue"] * minimum["green"]
	}
	fmt.Println(powersum)
}
