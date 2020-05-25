package algorithm

func Sort(values []int) ([]int, error) {
	return sort(values, 0, len(values) - 1)
}

func sort(values []int, low, high int) ([]int, error) {
	if low >= high {
		return values, nil
	}
	first, last, key := low, high, values[low]
	for first < last {
		for first < last && values[last] >= key {
			last -= 1
		}
		values [first] = values[last]
		for first < last && values[first] <= key {
			first += 1
		}
		values [last] = values[first]
	}
	values[first] = key
	_, _ = sort(values, low, first-1)
	_, _ = sort(values, first + 1, high)
	return values, nil
}
