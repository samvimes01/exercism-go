package allergies

var allergies = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

var keys = [8]uint{1, 2, 4, 8, 16, 32, 64, 128}

// Allergies ...
func Allergies(score uint) []string {
	result := make([]string, 0)

	for _, key := range keys {
		if key&score > 0 {
			result = append(result, allergies[key])
		}
	}

	return result
}

// AllergicTo ...
func AllergicTo(score uint, substance string) bool {
	for _, key := range keys {
		if key&score > 0 && allergies[key] == substance {
			return true
		}
	}

	return false
}
