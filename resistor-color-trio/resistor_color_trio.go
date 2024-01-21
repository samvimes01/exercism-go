package resistorcolortrio

import (
	"strconv"
)

var colorCode = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	powten, suffix := ColorSuffix(colors[2])
	first := colorCode[colors[0]]
	second := colorCode[colors[1]]
	num := 0
	if first > 0 && second > 0 {
		num = first * 10 + second
	} else if first > 0 {
		num = first
	} else if second > 0 {
		num = second
	}
	return strconv.Itoa(num * powten) + " " + suffix + "ohms"
}

// ColorCode returns the resistance value of the given color.
func ColorSuffix(color string) (int, string) {
	switch color {
	case "black":
		return 1, ""
	case "brown":
		return 10, ""
	case "red":
		return 1, "kilo"
	case "orange":
		return 1, "kilo"
	case "yellow":
		return 10, "kilo"
	case "green":
		return 100, "kilo"
	case "blue":
		return 1, "mega"
	case "violet":
		return 10, "mega"
	case "grey":
		return 100, "mega"
	case "white":
		return 1, "giga"
	}
	return 1, ""
}
