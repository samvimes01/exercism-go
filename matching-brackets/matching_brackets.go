package brackets

func Bracket(input string) bool {
  openBrackStack := make([]rune, 0, len(input))
  for _, l := range input {
    switch {
    case l == '(' || l == '{' || l == '[':
      openBrackStack = append(openBrackStack, l)
    case l == ')' || l == '}' || l == ']':
      opnLen := len(openBrackStack)
      if opnLen > 0 {
        if !isCorresponding(openBrackStack[opnLen - 1], l) {
          return false
        }
        openBrackStack = openBrackStack[:opnLen-1]
      } else {
        return false
      }
    }
  }
  return len(openBrackStack) == 0
}

func isCorresponding(a rune, b rune) bool {
  switch {
  case a == '(' && b == ')':
    return true
  case a == '{' && b == '}':
    return true
  case a == '[' && b == ']':
    return true
  default:
    return false
  }
}