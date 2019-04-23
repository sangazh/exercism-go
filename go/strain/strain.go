package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(f func(int) bool) (result Ints) {
	if len(i) == 0 {
		return nil
	}
	result = make([]int, 0)
	for _, item := range i {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
func (i Ints) Discard(f func(int) bool) (result Ints) {
	if len(i) == 0 {
		return nil
	}
	result = make([]int, 0)
	for _, item := range i {
		if !f(item) {
			result = append(result, item)
		}
	}
	return result
}
func (l Lists) Keep(f func([]int) bool) (result Lists) {
	if len(l) == 0 {
		return nil
	}
	result = make([][]int, 0)
	for _, item := range l {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
func (s Strings) Keep(f func(string) bool) (result Strings) {
	if len(s) == 0 {
		return nil
	}

	result = make([]string, 0)
	for _, item := range s {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
