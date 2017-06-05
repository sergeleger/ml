// Implements batch-gradient-descent described in http://cs229.stanford.edu/notes/cs229-notes1.pdf.
//
// Source of data:
//   http://openclassroom.stanford.edu/MainFolder/DocumentPage.php?course=MachineLearning&doc=exercises/ex2/ex2.html
//   http://openclassroom.stanford.edu/MainFolder/DocumentPage.php?course=MachineLearning&doc=exercises/ex3/ex3.html
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/sergeleger/strconv"
)

func main() {
	// Create and parse the flags
	alpha := flag.Float64("alpha", 1.0, "learning rate")
	n := flag.Int("n", 1500, "number of iterations")
	epsilon := flag.Float64("epsilon", 1e-10, "epsilon stop condition")
	norm := flag.Bool("norm", false, "apply standard score normalisation to X")
	flag.Parse()

	// read the data from standard in.
	y, err := readY(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	x, err := readX(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	// Initial solution is set to 0
	params := make([]float64, len(x[0]))
	for i := range params {
		params[i] = 0
	}

	// Normalise if request
	if *norm {
		x = normalizeStandardScore(x)
	}

	var costs = [...]float64{cost(params, y, x), 0}
	var iter int

	for iter = 0; iter < *n && math.Abs(costs[0]-costs[1]) > *epsilon; iter++ {
		params = updateParam(params, y, x, *alpha)
		if iter == 0 {
			fmt.Println("First:", params)
		}

		costs[0], costs[1] = costs[1], cost(params, y, x)
	}

	fmt.Println("Iterations:", iter, "Final:", params)
}

func updateParam(param, y []float64, x [][]float64, learningRate float64) []float64 {
	newParam := make([]float64, len(param))

	for j := range param {
		newParam[j] = param[j] - learningRate*derivedCost(param, y, x, j)/float64(len(y))
	}

	return newParam
}

// derivedCost is the derived cost function
func derivedCost(param, y []float64, x [][]float64, j int) (dLeastSq float64) {
	for i := range x {
		dLeastSq += (rowCost(param, x[i]) - y[i]) * x[i][j]
	}
	return dLeastSq
}

// cost implements the least-square cost.
func cost(param, y []float64, x [][]float64) (leastSq float64) {
	for i := range x {
		rowScore := rowCost(param, x[i]) - y[i]
		leastSq += rowScore * rowScore
	}

	return leastSq / 2 * float64(len(y))
}

// rowCost applies the linear function for a single row.
func rowCost(param, x []float64) (sum float64) {
	for i := range param {
		sum += param[i] * x[i]
	}
	return sum
}

// readX reads the X values from the file.
func readX(f string) (x [][]float64, err error) {
	var tc strconv.Strconv

	var r io.ReadCloser
	if r, err = os.Open(f); err != nil {
		return
	}
	defer r.Close()

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		cols := strings.Fields(sc.Text())
		xV := make([]float64, len(cols)+1)
		xV[0] = 1.0
		for i, v := range cols {
			xV[i+1] = tc.Atof(v)
		}

		if err = tc.Clear(); err == nil {
			x = append(x, xV)
		}
	}

	return
}

// readY reads the Y values from the file.
func readY(f string) (y []float64, err error) {
	var tc strconv.Strconv

	var r io.ReadCloser
	if r, err = os.Open(f); err != nil {
		return
	}
	defer r.Close()

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		cols := strings.Fields(sc.Text())
		v := tc.Atof(cols[0])

		if err = tc.Clear(); err == nil {
			y = append(y, v)
		}
	}

	return
}

// normalizeStandardScore normalises each x columns.
func normalizeStandardScore(x [][]float64) [][]float64 {
	avg := mean(x)
	std := standardDeviation(x, avg)

	for i := range x {
		for j, v := range x[i][1:] {
			x[i][j+1] = (v - avg[j+1]) / std[j+1]
		}
	}

	return x
}

// mean calculates the mean of each columns of x
func mean(x [][]float64) []float64 {
	avg := make([]float64, len(x[0]))
	for i := range x {
		for j, v := range x[i] {
			avg[j] += v
		}
	}

	n := float64(len(x))
	for i := range avg {
		avg[i] /= n
	}

	return avg
}

// standardDeviation calculates the standard deviation of each columns of x
func standardDeviation(x [][]float64, avg []float64) []float64 {
	std := make([]float64, len(x[0]))

	var d float64
	for i := range x {
		for j, v := range x[i] {
			d = (v - avg[j])
			std[j] += d * d
		}
	}

	n := float64(len(x))
	for i := range std {
		std[i] = math.Sqrt(std[i] / (n - 1))
	}

	return std
}
