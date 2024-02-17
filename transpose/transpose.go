package transpose

func Transpose(input []string) []string {
	maxLength := func(input []string) (max int) {
		for _, line := range input {
			if len(line) > max {
				max = len(line)
			}
		}
		return
	}
	max := maxLength(input)
	var output = make([]string, max)
	for i, row := range input {
		for j, char := range row {
			output[j] += string(char)
		}
		remainToMax := maxLength(input[i:])
		for j := len(row); j < remainToMax; j++ {
			output[j] += " "
		}
	}
	return output
}