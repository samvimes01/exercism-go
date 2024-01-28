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
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}