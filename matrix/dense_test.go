package matrix

import "testing"

// dense.go:func (m *Dense) at(i, j int) DType { return m.data[j + i*m.cols] }
// dense.go:func (m *Dense) Dims() (r, c int) { return m.rows, m.cols }
// dense.go:func (m *Dense) Rows() int { return m.rows }
// dense.go:func (m *Dense) Cols() int { return m.cols }
// dense.go:func NewDense(r, c int, data []DType) *Dense {
// dense.go:func NewDenseCopy(m *Dense) *Dense {
// dense.go:func NewDenseZeros(r, c int) *Dense {
// dense.go:func NewDenseIdentity(n int) *Dense {
// dense.go:func (m *Dense) T() Matrix {

func TestNewDense(t *testing.T) {
	t.Parallel()
    for id, test := range []struct {
		a []DType
		rows, cols int
		min, max DType
		mat Dense
	}{
		{
			[]DType{
				1, 2, 3,
				4, 5, 6,
			},
			2, 3,
			1, 6,
			Dense{
				[]DType{
					1, 2, 3,
					4, 5, 6,
				},
				2, 3,
			},
		},
    } {
		m1 := NewDense(test.rows, test.cols, test.a)
		if rows := m1.Rows(); rows != test.rows {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", id, rows, test.rows)
		}
		if cols := m1.Cols(); cols != test.cols {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", id, cols, test.cols)
		}
		if min := Min(m1); min != test.min {
			t.Errorf("unexpected min for test %d: got: %v want: %v", id, min, test.min)
		}
		if max := Max(m1); max != test.max {
			t.Errorf("unexpected max for test %d: got: %v want: %v", id, max, test.max)
		}
		// if fro := Norm(m, 2); math.Abs(Norm(m, 2)-test.fro) > 1e-14 {
		// 	t.Errorf("unexpected Frobenius norm for test %d: got: %v want: %v", id, fro, test.fro)
		// }
		// if !reflect.DeepEqual(m, test.mat) {
		// 	t.Errorf("unexpected matrix for test %d", id)
		// }
		// if !Equal(m, test.mat) {
		// 	t.Errorf("matrix does not equal expected matrix for test %d", id)
		// }
	}
}


func TestNewDenseZeros(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic happened: got: %v want: no panic", r)
        }
	}()
	rows, cols := 2, 3
	m := NewDenseZeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var got  DType = m.At(i, j)
			var want DType = 0
			if got != 0 {
				t.Errorf("unexpected matrix value on row %d col %d: got: %v want: %v", i, j, got, want)
			}
		}
	}
}


func TestNewDenseIdentity(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic happened: got: %v want: no panic", r)
        }
	}()
	n := 5
	m := NewDenseIdentity(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var got  DType = m.At(i, j)
			var want DType = 0
			if i == j {
				want = 1
			}
			if got != want {
				t.Errorf("unexpected matrix value on row %d col %d: got: %v want: %v", i, j, got, want)
			}
		}
	}
}


func TestNewDenseCopy(t *testing.T) {
	rows, cols := 2, 3
	m1 := NewDense(rows, cols, []DType{
		1, 2, 3,
		4, 5, 6,
	})
	m2 := NewDenseCopy(m1)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var got  DType = m2.At(i, j)
			var want DType = m1.At(i, j)
			if got != want {
				t.Errorf("unexpected matrix value on row %d col %d: got: %v want: %v", i, j, got, want)
			}
		}
	}
	// TODO: test modify m1, and check that m2 remained the same
}


func TestDenseAt(t *testing.T) {
	rows, cols := 2, 3
	data := []DType{
		1, 2, 3,
		4, 5, 6,
	}
	m := NewDense(rows, cols, data)
	for i := range rows {
		for j := range cols {
			want := data[j + i*cols]
			got  := m.At(i,j)
			if got != want {
				t.Errorf("unexpected values returned from matrix At on row %d col %d: got: %v want: %v",
					i, j, got, want)
			}
		}
	}
}


func TestDenseDims(t *testing.T) {
	t.Parallel()
    a := struct {
		m []DType
		rows, cols int
	}{
        []DType{
            1, 2, 3,
            4, 5, 6,
        },
		2, 3,
    }
	m := NewDense(a.rows, a.cols, a.m)
	if rows, cols := m.Dims(); rows != a.rows || cols != a.cols {
		t.Errorf("unexpected values returned from matrix Dims: got: %v want: %v",
			[]int{rows, cols}, []int{a.rows, a.cols} )
	}
	if rows := m.Rows(); rows != a.rows {
		t.Errorf("unexpected values returned from matrix Rows: got: %v want: %v",
			rows, a.rows )
	}
	if cols := m.Cols(); cols != a.cols {
		t.Errorf("unexpected values returned from matrix Cols: got: %v want: %v",
			cols, a.cols )
	}
}


func TestDenseTranspose(t *testing.T) {
	rows, cols := 2, 3
	m1 := NewDense(rows, cols, []DType{
		1, 2, 3,
		4, 5, 6,
	})
	m2 := m1.T()
	for i := range rows {
		for j := range cols {
			want := m1.At(i, j)
			got  := m2.At(j, i)
			if got != want {
				t.Errorf("unexpected values returned from matrix Transpose on row %d col %d: got: %v want: %v",
					i, j, got, want)
			}
		}
	}
}


