package sublist

type Relation string

type code int

const (
	equal     code = 1
	unequal   code = -1
	sublist   code = 2
	superlist code = -2
)

var relationMap = map[code]Relation{
	equal:     "equal",
	unequal:   "unequal",
	sublist:   "sublist",
	superlist: "superlist",
}

func (c code) String() Relation {
	return relationMap[c]
}

func (c code) reverse() code {
	return -c
}

// Sublist determine the given two lists' relation, sublist, superlist, equal or unequal
func Sublist(listA, listB []int) Relation {
	if len(listA) == 0 && len(listB) == 0 {
		return equal.String()
	}

	if len(listA) == 0 && len(listB) > 0 {
		return sublist.String()
	}

	if len(listA) > 0 && len(listB) == 0 {
		return superlist.String()
	}

	if len(listA) == len(listB) {
		return isEqual(listA, listB).String()
	}

	if len(listA) < len(listB) {
		return isSublist(listA, listB).String()
	}

	if len(listA) > len(listB) {
		if code := isSublist(listB, listA); code != unequal {
			return code.reverse().String()
		} else {
			return unequal.String()
		}
	}

	return equal.String()
}

//determine if first list is sublist of second list
func isSublist(listA, listB []int) code {
	if len(listA) < len(listB) {
		if listA[0] != listB[0] {
			return isSublist(listA, listB[1:])
		}

		equalFlag := true
		for i, a := range listA {
			if a != listB[i] {
				equalFlag = false
			}
		}
		if equalFlag {
			return sublist
		}
		return isSublist(listA, listB[1:])
	}

	if isEqual(listA, listB) > 0 {
		return sublist
	} else {
		return unequal
	}

}

// determine if two lists equal
func isEqual(listA, listB []int) code {
	for i, a := range listA {
		if a != listB[i] {
			return unequal
		}
	}
	return equal
}
