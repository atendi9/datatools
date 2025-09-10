package datatools

import (
	"math"
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	t.Run("Simple positive vector", func(t *testing.T) {
		vec := []float64{3, 4}
		got := Normalize(vec)
		want := []float64{0.6, 0.8} // (3/5, 4/5)
		if !reflect.DeepEqual(RoundSlice(got, 2), want) {
			t.Errorf("Normalize(%v) = %v, want %v", vec, got, want)
		}
	})

	t.Run("Negative values", func(t *testing.T) {
		vec := []float64{-3, -4}
		got := Normalize(vec)
		want := []float64{-0.6, -0.8}
		if !reflect.DeepEqual(RoundSlice(got, 2), want) {
			t.Errorf("Normalize(%v) = %v, want %v", vec, got, want)
		}
	})

	t.Run("Already normalized", func(t *testing.T) {
		vec := []float64{0, 1}
		got := Normalize(vec)
		want := []float64{0, 1}
		if !reflect.DeepEqual(RoundSlice(got, 2), want) {
			t.Errorf("Normalize(%v) = %v, want %v", vec, got, want)
		}
	})

	t.Run("Zero vector", func(t *testing.T) {
		vec := []float64{0, 0, 0}
		got := Normalize(vec)
		// Vai gerar divisão por zero => NaN
		if !math.IsNaN(got[0]) {
			t.Errorf("Normalize(%v) = %v, expected NaN values", vec, got)
		}
	})

	t.Run("Single element vector", func(t *testing.T) {
		vec := []float64{5}
		got := Normalize(vec)
		want := []float64{1}
		if !reflect.DeepEqual(RoundSlice(got, 2), want) {
			t.Errorf("Normalize(%v) = %v, want %v", vec, got, want)
		}
	})
}
