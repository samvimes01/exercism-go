package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
func Keep[T any](list []T, filter func(T) bool)[]T {
  if len(list) == 0 {
    return list
  }
  filtered := make([]T, 0, len(list))
  for i, v := range list {
    if filter(v) {
      filtered = append(filtered, list[i])
    }
  }
  return filtered
}

func Discard[T any](list []T, filter func(T) bool)[]T {
  return Keep(list, func (el T) bool { return !filter(el) })
}