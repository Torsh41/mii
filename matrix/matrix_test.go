package matrix

import "testing"

// Borrowed from github.com/gonum/gonum/blob/master/mat/matrix_test.go

// TODO:
// func (m *Dense) Dims() (int, int)
// func (m *Dense) Rows() int
// func (m *Dense) Cols() int
// func NewDense(data []DType, h int, w int) Dense
// func NewDenseZeros(h int, w int) Dense
// func (m1 *Dense) NewDenseCopy() (m2 Dense)
// func NewDenseIdentity(n int) Dense
// func T(m *Dense) Dense
// func Add(m1 *Dense, m2 *Dense) Dense
// func Mul(m1 *Dense, m2 *Dense) Dense
// func Inv(m *Dense) Dense
// func (m *Dense) Print()


func flatten(f [][]DType) (r, c int, d []DType) {
    r = len(f)
    if r == 0 {
        panic("bad test: no row")
    }
    c = len(f[0])
    d = make([]DType, 0, r*c)
    for _, row := range f {
        if len(row) != c {
            panic("bad test: ragged input")
        }
        d = append(d, row...)
    }
    return r, c, d
}


func unflatten(r, c int, d []DType) [][]DType {
    m := make([][]DType, r)
    for i := 0; i < r; i++ {
        m[i] = d[i*c : (i+1)*c]
    }
    return m
}

// matrix.go:func Max(m Matrix) DType {
// matrix.go:func Min(m Matrix) DType {
// index_no_bounds_checks.go:func (m *Dense) At(i, j int) DType {
// math.go:func Add(m1, m2 *Dense) *Dense {
// math.go:func Mul(m1, m2 *Dense) *Dense {
// math.go:func Inv(m *Dense) *Dense {
// io.go:func (m *Dense) Print() {

func TestStuff(t *testing.T) {
}
