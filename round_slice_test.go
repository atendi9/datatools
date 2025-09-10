package datatools

import (
	"reflect"
	"testing"
)

func TestRoundSlice(t *testing.T) {
	t.Run("Round positive numbers with precision 2", func(t *testing.T) {
		input := []float64{1.2345, 2.3456, 3.9999}
		got := RoundSlice(input, 2)
		want := []float64{1.23, 2.35, 4.00}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Round negative numbers with precision 3", func(t *testing.T) {
		input := []float64{-1.2345, -2.3456, -3.9999}
		got := RoundSlice(input, 3)
		want := []float64{-1.235, -2.346, -4.000}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Round with precision 0", func(t *testing.T) {
		input := []float64{1.2, 2.5, 3.7}
		got := RoundSlice(input, 0)
		want := []float64{1, 3, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Empty slice", func(t *testing.T) {
		input := []float64{}
		got := RoundSlice(input, 2)
		want := []float64{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
