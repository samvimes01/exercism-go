package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return RelationEqual
	} else if isSublist(l1, l2) {
		return RelationSublist
	} else if isSublist(l2, l1) {
		return RelationSuperlist
	} else {
		return RelationUnequal
	}
}

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func isSublist(assertShort, assertLong []int) bool {
	if len(assertShort) > len(assertLong) {
		return false
	}	
	if len(assertShort) == 0 {
		return true
	}
	equals := false
	// for each number in a long list check each number in short one
	// if equal - just return, if non equal after n-th try - continue outer cycle
	// outer cycle  limits to diff since inner cover diff
	for longInd := 0; longInd <= len(assertLong)-len(assertShort); longInd++ {
		for i := 0; i < len(assertShort); i++ {
			if assertShort[i] == assertLong[longInd+i] {
				equals = true
			} else {
				equals = false
				break
			}
		}
		if equals {
			return equals
		}
	}
	return false
}
