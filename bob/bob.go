package bob

import "strings"

func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	if trimmed == "" {
		return "Fine. Be that way!"
	}
	isQuestion := trimmed[len(trimmed)-1] == '?'
	isYelling := trimmed == strings.ToUpper(trimmed) && strings.ToLower(trimmed) != trimmed

	if isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion {
		return "Sure."
	}
	if isYelling {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
