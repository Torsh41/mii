package matrix

// TODO:
// - get/at(), set, slice???
// - jacobi
// - inverse, norm, det
// - Add/Mul by number
// - Module description (a doc.go file?)

type DType = float64;

// NewDenseCopy of gonum.mat.Matrix
type Matrix interface {
	// Dims returns the dimensions of a Matrix.
	Dims() (r, c int)

	// At returns the value of a matrix element at row i, column j.
	// It will panic if i or j are out of bounds for the matrix.
	At(i, j int) DType

	// T returns the transpose of the Matrix. Whether T returns a copy of the
	// underlying data is implementation dependent.
	// This method may be implemented using the Transpose type, which
	// provides an implicit matrix transpose.
	T() Matrix
}


// Min finds max value of the matrix
func Max(m Matrix) DType {
	// TODO: extend like gonum: implement RawMatrixer interface
	if dense, ok := m.(*Dense); ok {
		max := dense.data[0]
		for _, num := range dense.data {
			if num > max {
				max = num
			}
		}
		return max
	} else {
		max := m.At(0,0)
		rows, cols := m.Dims()
		for i := range rows {
			for j := range cols {
				e := m.At(i,j)
				if e > max {
					max = e
				}
			}
		}
		return max
	}
}


// Min finds min value of the matrix
func Min(m Matrix) DType {
	// TODO: extend like gonum: implement RawMatrixer interface
	if dense, ok := m.(*Dense); ok {
		min := dense.data[0]
		for _, num := range dense.data {
			if num < min {
				min = num
			}
		}
		return min
	} else {
		min := m.At(0,0)
		rows, cols := m.Dims()
		for i := range rows {
			for j := range cols {
				e := m.At(i,j)
				if e < min {
					min = e
				}
			}
		}
		return min
	}
}


