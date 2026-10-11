package matrix

// TODO: Can these methods return anything other than *Dense?
// TODO: Implement other ops with other Matrix use-cases


// Add computes addition of a and b, and returns a new matrix c.
// If a.Dims() != b.Dims(), it will panic.
func Add(a, b Matrix) (c Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		panic(ErrShape)
	}
	if arm, ok := a.(*Dense); ok {
		if brm, ok := b.(*Dense); ok {
			c = NewDenseZeros(ar, ac)
			crm, _ := c.(*Dense)
			for i := range ar*ac {
				crm.data[i] = arm.data[i] + brm.data[i]
			}
			return c;
		}
	}
	return nil
}


// Mul computes multiplication of a and b, and returns a new matrix c.
// if !(a.rows >= b.rows) or !(a.cols <= b.cols), it will panic
func Mul(a, b Matrix) (c Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar < br || ac > bc {
		panic(ErrShape)
	}
	if arm, ok := a.(*Dense); ok {
		if brm, ok := b.(*Dense); ok {
			c = NewDenseZeros(ar, bc)
			crm, _ := c.(*Dense)
			for i := range ar {
				for j := range bc {
					var val DType = 0
					for k := range min(ac, br) {
						val = val + arm.At(i, k) * brm.At(k, j)
					}
					crm.Set(i, j, val)
				}
			}
			return c
		}
	}
	return nil
}


// TODO: Inv computes inverse of m
// func Inv(m *Matrix) *Matrix {
// 	inv := NewDenseZeros(m.rows, m.cols)
// 	return inv
// }
//
//
