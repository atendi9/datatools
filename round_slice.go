package datatools

import "math"

// RoundSlice helper function to round results (comparison with floats).
func RoundSlice(vec []float64, precision int) []float64 {
	result := make([]float64, len(vec))
	factor := math.Pow(10, float64(precision))
	for i, v := range vec {
		result[i] = math.Round(v*factor) / factor
	}
	return result
}
