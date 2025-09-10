package datatools

// DotProduct computes the dot product of two slices.
func DotProduct(vec1, vec2 []float64) float64 {
	var sum float64
	for i, value := range vec1 {
		sum += value * vec2[i]
	}
	return sum
}