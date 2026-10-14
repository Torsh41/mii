package matrix

import "testing"
import "math"


func matrixEqual(a, b Matrix) bool {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		panic(ErrShape)
	}
	const eps = 1e-4
	if arm, ok := a.(*Dense); ok {
		if brm, ok := b.(*Dense); ok {
			for i := range ar*ac {
				if math.Abs(arm.data[i] - brm.data[i]) >= eps {
					return false
				}
			}
			return true
		}
	}
	return false
}


func TestEigenJacobi(t *testing.T) {
	t.Parallel()
    for id, test := range []struct {
		A, L, V Matrix
	}{
		{
			NewDense(4, 4, []DType{
				4, -30, 60, -35,
				-30, 300, -675, 420,
				60, -675, 1620, -1050,
				-35, 420, -1050, 700,
			}),
			NewDense(4, 4, []DType{
				0.166643, -0.000000, -0.000007, -0.000194,
				0.000000, 37.101491, 0.000000, -0.000000,
				-0.000007, 0.000000, 2585.253811, -0.000009,
				-0.000194, -0.000000, -0.000009, 1.478055,
			}),
			NewDense(4, 4, []DType{
				0.792695, -0.179186, 0.029193, -0.581958,
				0.451868, 0.741918, -0.328712, 0.370569,
				0.322341, -0.100228, 0.791411, 0.509626,
				0.252085, -0.638283, -0.514553, 0.514086,
			}),
		},
    } {
		eigenval, eigenvec := Eigen(test.A)
		if !matrixEqual(eigenval, test.L) {
			t.Errorf("eigenvalues don't match in test %d: got: %v want: %v",
					id, eigenval, test.L)
		}
		if !matrixEqual(eigenvec, test.V) {
			t.Errorf("eigenvectors don't match in test %d: got: %v want: %v",
					id, eigenvec, test.V)
		}
	}
}
