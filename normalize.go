package datatools

import "math"

// Normalize computes the Euclidean norm of a slice and returns the normalized slice.
func Normalize(vec []float64) []float64 {
	normVal := norm(vec)
	normalizedVec := make([]float64, len(vec))
	for i, v := range vec {
		normalizedVec[i] = v / normVal
	}
	return normalizedVec
}

// norm computes the Euclidean norm of a slice.
func norm(vec []float64) float64 {
	var sum float64
	for _, v := range vec {
		sum += v * v
	}
	return math.Sqrt(sum)
}
