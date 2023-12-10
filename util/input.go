package util

import (
	"strings"
)

func TrimInput(input string) string {
	
	out := strings.TrimRight(input, "\n")
	if len(out) == 0 {
		panic("empty input.txt file!")
	}
	 return out
}
