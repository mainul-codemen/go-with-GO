package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{
		values:  []float64{1, 2},
		average: 1.5,
	},
	{
		values:  []float64{1, 2, 3},
		average: 2.0,
	},
	{
		values:  []float64{10, 20, 10, 5, 5},
		average: 10.0,
	},
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
