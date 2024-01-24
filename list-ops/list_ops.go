package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	l := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	l := make(IntList, 0)
	for _, v := range s {
		l = append(l, fn(v))
	}
	return l
}

func (s IntList) Reverse() IntList {
	l := make(IntList, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		l = append(l, s[i])
	}
	return l
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	l := make(IntList, 0)
	for _, v := range lists {
		l = append(l, v...)
	}
	return append(s, l...)
}
