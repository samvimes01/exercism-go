package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, ch := range log {
		switch {
		case ch == []rune("❗")[0]:
			return "recommendation"
		case ch == []rune("🔍")[0]:
			return "search"
		case ch == []rune("☀")[0]:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// runes := make([]rune, utf8.RuneCountInString(log))
	// for i, ch := range log {
	// 	if ch == oldRune {
	// 		runes[i] = newRune
	// 	} else {
	// 		runes[i] = ch
	// 	}
	// }
	// return string(runes) / doesn't work when rine replaced wih single char
	r := strings.NewReplacer(string(oldRune), string(newRune))
	return r.Replace(log)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
