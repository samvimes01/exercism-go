package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
func Keep[T any](list []T, f func(T) bool)[]T {
  filtered := make([]T, 0)
  for i, v := range list {
    if f(v) {
      filtered = append(filtered, list[i])
    }
  }
  return filtered
}

func Discard[T any](list []T, f func(T) bool)[]T {
  filtered := make([]T, 0)
  for i, v := range list {
    if !f(v) {
      filtered = append(filtered, list[i])
    }
  }
  return filtered
}