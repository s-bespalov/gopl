package maxmin

func max(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vavalues ...int) int {
	if len(vavalues) == 0 {
		return 0
	}
	min := vavalues[0]
	for _, val := range vavalues {
		if val < min {
			min = val
		}
	}
	return min
}
