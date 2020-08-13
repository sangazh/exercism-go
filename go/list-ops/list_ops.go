package listops

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int
type IntList []int

func (l IntList) Foldr(fn binFunc, init int) int {
	if l.Length() == 0 {
		return init
	}

	for i := l.Length() - 1; i >= 0; i-- {
		init = fn(l[i], init)
	}

	return init
}

func (l IntList) Foldl(fn binFunc, init int) int {
	if l.Length() == 0 {
		return init
	}

	for _, n := range l {
		init = fn(init, n)
	}

	return init
}

func (l IntList) Filter(fn predFunc) (result IntList) {
	if l.Length() == 0 {
		return IntList{}
	}
	for _, n := range l {
		if fn(n) {
			result = append(result, n)
		}
	}
	return
}
func (l IntList) Length() int {
	return len(l)
}
func (l IntList) Map(fn unaryFunc) IntList {
	if l.Length() == 0 {
		return IntList{}
	}

	result := make([]int, l.Length())
	for i, n := range l {
		result[i] = fn(n)
	}

	return result
}
func (l IntList) Reverse() IntList {
	if l.Length() == 0 {
		return IntList{}
	}
	result := make([]int, l.Length())
	for i, n := range l {
		result[l.Length()-i-1] = n
	}

	return result
}
func (l IntList) Append(list IntList) IntList {
	if l.Length() == 0 && list.Length() == 0 {
		return IntList{}
	}

	result := make([]int, l.Length()+list.Length())
	for i, n := range l {
		result[i] = n
	}

	for j, n := range list {
		result[l.Length()+j] = n
	}

	return result
}

func (l IntList) Concat(args []IntList) IntList {
	for _, list := range args {
		l = l.Append(list)
	}

	return l
}
