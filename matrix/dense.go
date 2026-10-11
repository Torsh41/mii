package matrix


type Dense struct {
	data []DType
	rows, cols int
}


// Dims returns matrix's Rows and Cols in the given order.
func (m *Dense) Dims() (r, c int) { return m.rows, m.cols }
// Rows returns the number of rows in the matrix
func (m *Dense) Rows() int { return m.rows }
// Cols returns the number of colsumns  in the matrix
func (m *Dense) Cols() int { return m.cols }


// At returns the value of a matrix element at row i, column j.
// If indices are out of range and when built with bounds tag, it will panic.
// This function has two wrappers in files index_bounds_checks.go and
// index_no_bounds_checks.go
func (m *Dense) at(i, j int) DType { return m.data[j + i*m.cols] }


// Set is used to assign a value to a matrix element at row i, column j
// If indices are out of range and when built with bounds tag, it will panic.
// This function has two wrappers in files index_bounds_checks.go and
// index_no_bounds_checks.go
func (m *Dense) set(i, j int, val DType) { m.data[j + i*m.cols] = val }


// NewDense creates a new matrix object with rows r and cols c and
// assigns it values of data. If length of data doesn't match matrix's cols
// and rows, it will panic.
func NewDense(r, c int, data []DType) *Dense {
	if len(data) != r * c {
		panic(ErrSliceLengthMismatch)
	}
	return &Dense{data, r, c}
}


// NewDenseCopy creates a copy of the matrix object.
func NewDenseCopy(m *Dense) *Dense {
	data := make([]DType, m.rows * m.cols)
	copy(data, m.data)
	return &Dense{data, m.rows, m.cols}
}


// NewDenseZeros creates a new matrix object with rows r and cols c and
// fills matrix with zeros
func NewDenseZeros(r, c int) *Dense {
	data := make([]DType, r * c)
	return &Dense{data, r, c}
}


// NewDenseIdentity creates the square identity matrix of size n.
func NewDenseIdentity(n int) *Dense {
	r := NewDenseZeros(n, n)
	for i := range n {
		r.Set(i, i, 1)
	}
	return r
}


// T computes the transpose of matrix m.
func (m *Dense) T() Matrix {
	rows, cols := m.Dims()
	r := NewDenseZeros(cols, rows)
	for i := range rows {
		for j := range cols {
			r.Set(j, i, m.At(i, j))
		}
	}
	return r
}

