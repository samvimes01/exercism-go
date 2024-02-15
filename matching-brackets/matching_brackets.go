package brackets

func Bracket(input string) bool {
	dict := make(map[rune]int)

	for _, char := range input {
		dict[char]++
	}
	if dict['('] != dict[')'] || dict['{'] != dict['}'] || dict['['] != dict[']'] {
		return false
	}
	return true
}
