package clr

func min(args ...float64) float64 {
	var min float64
	min = 255
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

func max(args ...float64) float64 {
	var max float64
	max = 0
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func pow(x uint8, y int) int {
	r := int(x)
	i := 1
	for ; i < y; i++ {
		r *= int(x)
	}
	return r
}
