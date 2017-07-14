package main

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestMatrixDense(t *testing.T) {
	is := is.New(t)

	d := &Dense{r: 2, c: 3, data: []float64{1, 2, 3, 4, 5, 6}}

	is.Equal(d.Get(0, 0), 1)
	is.Equal(d.Get(0, 1), 2)
	is.Equal(d.Get(0, 2), 3)
	is.Equal(d.Get(1, 0), 4)
	is.Equal(d.Get(1, 1), 5)
	is.Equal(d.Get(1, 2), 6)
}
