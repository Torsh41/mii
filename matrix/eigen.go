package matrix

import "math"
import "fmt"


var eps = 1e-6

// Eigen calculates eigenvalues of the matrix a, and stores it in L.
// Also find eigenvectors, and stores them in V.
// Might panic, sometimes.
func Eigen(a Matrix) (L, V Matrix) {
	return eigenJacobi(a)
}


func eigenJacobi(A Matrix) (L, V Matrix) {
	nLim := 100
	rows, cols := A.Dims()
	if rows != cols {
		panic(ErrShape)
	}
	data := make([]DType, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != j {
				data[j + i*cols] = 1
			}
		}
	}
	// Matrix with all elements=1 except the main diag
	funny := NewDense(rows, cols, data)
	// init eigenvectors
	var A1 Matrix = A
	V = NewDenseIdentity(rows)
	for n := 0; n < nLim; n++ {
		// choose p, q: find max val in upper triangular matrix
		var p, q int = 0, 0
		maxi, maxj := 0, 0
		maxv := 0.0
		for i := 0; i < rows; i++ {
			for j := i + 1; j < cols; j++ {
				val := float64(A.At(i,j) * A.At(i,j))
				if val > maxv {
					maxv = val
					maxi, maxj = i, j
				}
			}
		}
		p, q = maxi, maxj
		// calc theta
		theta := 0.5 * math.Atan(2 * A.At(p,q) / (A.At(q,q) - A.At(p,p)))
		costheta := math.Cos(theta)
		sintheta := math.Sin(theta)
		// // build P, P^T matrix
		P := NewDenseIdentity(rows)
		P.Set(p, p, DType(costheta))
		P.Set(q, q, DType(costheta))
		P.Set(p, q, DType(sintheta))
		P.Set(q, p, -DType(sintheta))
		// calc A_{i+1}
		A1 = Mul(Mul(P.T(), A), P)
		// calc V_{i+1}
		V = Mul(V, P)
		// exit condition: Max(A^2), excluding main diagonal
		if Max(MulElem(MulElem(A1, A1), funny)) < eps {
			break
		}
		A = A1
	}

	return A1, V
}


func eigenJacobi2(a Matrix) (L, V Matrix) {
	nLim := 100
	rows, cols := a.Dims()
	if rows != cols {
		panic(ErrShape)
	}
	var a1 Matrix = a
	data := make([]DType, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != j {
				data[j + i*cols] = 1
			}
		}
	}
	// Matrix with all elements=1 except the main diag
	funny := NewDense(rows, cols, data)
	a1 = a
	V = NewDenseIdentity(rows)
	// Init stuff
	for n := 0; n < nLim; n++ {
		// choose p, q
		// TODO: store cache of max val in each row
		var p, q int = 0, 0
		maxi, maxj := 0, 0
		maxv := 0.0
		for i := 0; i < rows; i++ {
			for j := i + 1; j < cols; j++ {
				val := float64(a.At(i,j) * a.At(i,j))
				if val > maxv {
					maxv = val
					maxi, maxj = i, j
				}
			}
		}
		p, q = maxi, maxj
		fmt.Printf("")
		// fmt.Printf("\n")
		// fmt.Printf("Iter: %d; ", n)
		// fmt.Printf("p, q = %d, %d; ", p, q)
		// fmt.Printf("\n")
		// Print(a1)
		// Print(MulElem(MulElem(a1, a1), funny))
		// calc sin, cos, tan
		theta := 0.5 * math.Atan(2 * a.At(p,q) / (a.At(q,q) - a.At(p,p)))
		tantheta := math.Tan(theta)
		costheta := math.Cos(theta)
		sintheta := math.Sin(theta)
		// build P, P^T matrix
		P := NewDenseIdentity(rows)
		P.Set(p, p, DType(costheta))
		P.Set(q, q, DType(costheta))
		P.Set(p, q, DType(sintheta))
		P.Set(q, p, -DType(sintheta))
		// calc A_{i+1}
		if dense, ok := a1.(*Dense); ok {
			for k := range rows {
				if k != p && k != q {
					val1 := costheta * a.At(p, k) - sintheta * a.At(q, k)
					val2 := sintheta * a.At(p, k) + costheta * a.At(q, k)
					dense.Set(k, p, val1)
					dense.Set(p, k, val1)
					dense.Set(k, q, val2)
					dense.Set(q, k, val2)
				}
			}
			dense.Set(p, p, a.At(p, p) - tantheta * a.At(p, q))
			dense.Set(q, q, a.At(q, q) + tantheta * a.At(p, q))
			dense.Set(p, q, 0)
			dense.Set(q, p, 0)
		} else {
			panic("Not the right matrix, what?")
		}
		// calc V_{i+1}
		if dense, ok := V.(*Dense); ok {
			for k := range rows {
				dense.Set(k, p, costheta * V.At(k, p) - sintheta * V.At(k, q))
				dense.Set(k, q, sintheta * V.At(k, p) + costheta * V.At(k, q))
			}
		} else {
			panic("Not the right matrix, what?")
		}
		// check exit condition
		if Max(MulElem(MulElem(a1, a1), funny)) < eps {
			break
		}
		a = a1
	}
	return a1, V
}
