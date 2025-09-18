package math

// Average returns the average value in the slice xs.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Max returns the maximum value in the slice xs.
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	max := xs[0]
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum value in the slice xs.
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	min := xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}
