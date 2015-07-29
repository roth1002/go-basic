package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type addtestpair struct {
	values []float64
	sum    float64
}

var addtests = []addtestpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
	{[]float64{1.1, 2.3}, 3.4},
}

func TestAdd(t *testing.T) {
	for _, pair := range addtests {
		v := Add(pair.values)
		assert.Equal(t, v, pair.sum, "they should be equal")
	}
}
