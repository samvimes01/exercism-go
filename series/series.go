package series

func All(n int, s string) []string {
	if n > len(s) || n < 0 {
		return nil
	}
	series := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if i + n > len(s) {
			break
		}
		serie := s[i : i+n]
		series = append(series, serie)
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) || n < 0 {
		return "", false
	}
	return s[:n], true
}