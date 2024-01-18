package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	items map[string]struct{}
}

func New() Set {
	return Set{items: make(map[string]struct{})}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, v := range l {
		set.items[v] = struct{}{}
	}
	return set
}

func (s Set) String() string {
	sb := strings.Builder{}
	sb.WriteString("{")
	i := 0
	for k := range s.items {
		sb.WriteString(`"` + k + `"`)
		i += 1
		if i < len(s.items) {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (s Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.items[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.items[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	if len(s1.items) == 0 && len(s2.items) == 0 {
		return true
	}
	if len(s1.items) != 0 && len(s2.items) == 0 {
		return false
	}
	isSubset := true
  for k := range s1.items {
		if _, ok := s2.items[k]; !ok {
			isSubset = false
			break
		}
	}
	return isSubset
}

func Disjoint(s1, s2 Set) bool {
	isDisjoint := true
  for k := range s1.items {
		if _, ok := s2.items[k]; ok {
			isDisjoint = false
			break
		}
	}
	return isDisjoint
}

func Equal(s1, s2 Set) bool {
	if len(s1.items) != len(s2.items) {
		return false
	}
	isEqual := true
	for k := range s1.items {
		if _, ok := s2.items[k]; !ok {
			isEqual = false
			break
		}
	}
	return isEqual
}

func Intersection(s1, s2 Set) Set {
	set := New()
	for k := range s1.items {
		if _, ok := s2.items[k]; ok {
			set.Add(k)
		}
	}
	return set
}

func Difference(s1, s2 Set) Set {
	set := New()
	if len(s1.items) == 0 {
		return set
	}
	biggerSet := s1
	smallerSet := s2
	if len(s2.items) > len(s1.items) {
		biggerSet = s2
		smallerSet = s1
	}
	for k := range biggerSet.items {
		if _, ok := smallerSet.items[k]; !ok {
			set.Add(k)
		}
	}
	return set
}

func Union(s1, s2 Set) Set {
	set := New()
	for k := range s1.items {
		set.Add(k)
	}
	for k := range s2.items {
		set.Add(k)
	}
	return set
}
