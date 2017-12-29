package clr

func min(args ...float64) float64 {
	min := 255.0
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

func max(args ...float64) float64 {
	max := 0.0
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}
