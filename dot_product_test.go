package datatools

import "testing"

func TestDotProduct(t *testing.T) {
	t.Run("positive vectors", func(t *testing.T) {
		vec1 := []float64{1, 2, 3}
		vec2 := []float64{4, 5, 6}
		want := 32.0 // 1*4 + 2*5 + 3*6
		got := DotProduct(vec1, vec2)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("vectors with negatives", func(t *testing.T) {
		vec1 := []float64{-1, 2, -3}
		vec2 := []float64{4, -5, 6}
		want := -32.0 // (-1*4) + (2*-5) + (-3*6)
		got := DotProduct(vec1, vec2)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty vectors", func(t *testing.T) {
		vec1 := []float64{}
		vec2 := []float64{}
		want := 0.0
		got := DotProduct(vec1, vec2)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("single element vectors", func(t *testing.T) {
		vec1 := []float64{2}
		vec2 := []float64{5}
		want := 10.0
		got := DotProduct(vec1, vec2)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("decimal values", func(t *testing.T) {
		vec1 := []float64{1.5, 2.5}
		vec2 := []float64{3.0, 4.0}
		want := 14.5 // 1.5*3 + 2.5*4
		got := DotProduct(vec1, vec2)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
