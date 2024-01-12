package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	borderGen := func(n int) string {
		return strings.Repeat("*", n)
	}
	return borderGen(numStarsPerLine) + "\n" + welcomeMsg  + "\n" + borderGen(numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msgLine := strings.Split(oldMsg, "\n")[1]
	msgLine, _ = strings.CutPrefix(msgLine, "*")
	msgLine, _ = strings.CutSuffix(msgLine, "*")
	msg := strings.TrimSpace(msgLine)
	return msg
}
