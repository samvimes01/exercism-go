package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	first := ColorCode(colors[0])
	second := ColorCode(colors[1])
	return first * 10 + second
}

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
	for i, element := range Colors() {
		if element == color {
			return i
		}
	}
	return -1
}