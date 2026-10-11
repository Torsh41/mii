// This file should be kept in sync with index_bounds_checks.go

//go:build bounds

package matrix

// At returns the value of a matrix element at row i, column j.
// If indices are out of range and when built with bounds tag, it will panic.
func (m *Dense) At(i, j int) DType {
	if i < 0 || i >= m.rows {
		panic(ErrRowAccess)
	}
	if j < 0 || j >= m.cols {
		panic(ErrColAccess)
	}
	return m.at(i, j)
}


// Set is used to assign a value to a matrix element at row i, column j
// If indices are out of range and when built with bounds tag, it will panic.
func (m *Dense) Set(i, j int, val DType) {
	if i < 0 || i >= m.rows {
		panic(ErrRowAccess)
	}
	if j < 0 || j >= m.cols {
		panic(ErrColAccess)
	}
	m.set(i, j, val)
}


