package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{values: []float64{1, 2}, average: 1.5},
	{values: []float64{}, average: 0},
	{values: []float64{-1, 1}, average: 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
