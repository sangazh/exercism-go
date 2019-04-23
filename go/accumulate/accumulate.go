package accumulate

//Accumulate convert listA by f function to result
func Accumulate(listA []string, f func(string) string) (result []string) {
	result = make([]string, len(listA))
	for i, a := range listA {
		result[i] = f(a)
	}
	return result
}
