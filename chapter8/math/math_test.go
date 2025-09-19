package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var avg_tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var max_tests = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
}

var min_tests = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, -1},
}

func TestAverage(t *testing.T) {
	for _, pair := range avg_tests {
		avg := Average(pair.values)
		if avg != pair.average {
			t.Error("Average of ", pair.values, "should be ", pair.average, "but was ", avg)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range max_tests {
		max := Max(pair.values)
		if max != pair.average {
			t.Error("Max of ", pair.values, "should be ", pair.average, "but was ", max)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range min_tests {
		min := Min(pair.values)
		if min != pair.average {
			t.Error("Min of ", pair.values, "should be ", pair.average, "but was ", min)
		}
	}
}
