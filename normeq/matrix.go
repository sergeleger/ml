package main

// Define a matrix interface
type Matrix interface {
	R() int
	C() int
	Get(r, c int) float64
	Set(r, c int, v float64)
}

// Dense defines a dense matrix
type Dense struct {
	r, c int
	data []float64
}

func (d *Dense) R() int {
	return d.r
}

func (d *Dense) C() int {
	return d.c
}

func (d *Dense) Get(r, c int) float64 {
	return d.data[r*d.c+c]
}

func (d *Dense) Set(r, c int, v float64) {
	d.data[r*d.c+c] = v
}

func NewDense(r, c int) *Dense {
	return &Dense{r: r, c: c, data: make([]float64, r*c)}
}

func NewDenseWithData(r, c int, data []float64) *Dense {
	d := &Dense{r: r, c: c}
	if r*c <= len(data) {
		d.data = data[0 : r*c]
	} else {
		d.data = make([]float64, r*c)
		copy(d.data, data)
	}

	return d

}
