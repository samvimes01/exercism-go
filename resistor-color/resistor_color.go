package resistorcolor

import "strings"

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	c := 0
	col := strings.ToLower(color)
	switch col {
	case "black":
		c = 0
	case "brown":
		c = 1
	case "red":
		c = 2
	case "orange":
		c = 3
	case "yellow":
		c = 4
	case "green":
		c = 5
	case "blue":
		c = 6
	case "violet":
		c = 7
	case "grey":
		c = 8
	case "white":
		c = 9
	}
	return c
}
