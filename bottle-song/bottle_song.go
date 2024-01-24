package bottlesong

import (
	"fmt"
	"strings"
)
var nums = map[int]string{
	0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var verses []string
	for i := startBottles; i > startBottles - takeDown; i-- {
		if i < startBottles {
			verses = append(verses, "")
		}
		verses = append(verses, verse(i)...)
	}
	return verses
}

func verse(n int) []string {
	var s = map[bool]string{true: "s", false: ""}
	lines := make([]string, 4)
	btl := "bottle"
	lines[0] = fmt.Sprintf("%s green %s hanging on the wall,", nums[n], btl + s[n > 1])
	lines[1] = fmt.Sprintf("%s green %s hanging on the wall,", nums[n], btl + s[n > 1])
	lines[2] = "And if one green bottle should accidentally fall,"
	lines[3] = fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(nums[n-1]), btl + s[n != 2])

	return lines
}
