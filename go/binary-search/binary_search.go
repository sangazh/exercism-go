package binarysearch

func SearchInts(slice []int, target int) (idx int) {
	if len(slice) == 0 {
		return -1
	}

	if len(slice) == 1 {
		if slice[0] == target {
			return 0
		}
		return -1
	}

	mid := len(slice) / 2
	if slice[mid] == target {
		return mid
	} else if slice[mid] < target {
		idx = SearchInts(slice[mid:], target)
		if idx < 0 {
			return idx
		} else {
			return idx + mid
		}
	}
	return SearchInts(slice[:mid], target)

}
