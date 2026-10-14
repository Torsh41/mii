package matrix

import "testing"


func TestAdd(t *testing.T) {
	t.Parallel()
    for id, test := range []struct {
		a, b, c Matrix
	}{
		{
			NewDense(2, 3, []DType{
				1, 2, 3,
				4, 5, 6,
			}),
			NewDense(2, 3, []DType{
				1, 2, 3,
				4, 5, 6,
			}),
			NewDense(2, 3, []DType{
				2, 4, 6,
				8, 10, 12,
			}),
		},
    } {
		res := Add(test.a, test.b)
		rows, cols := res.Dims()
		for i := range rows {
			for j := range cols {
				want := test.c.At(i, j)
				got  := res.At(i, j)
				if got != want {
					t.Errorf("unexpected values in function Add on row %d col %d in test %d: got: %v want: %v",
						i, j, id, got, want)
				}
			}
		}

	}
}


func TestMul(t *testing.T) {
	t.Parallel()
    for id, test := range []struct {
		a, b, c Matrix
	}{
		{
			NewDense(3, 2, []DType{
				1, 2,
				3, 4,
				5, 6,
			}),
			NewDense(2, 3, []DType{
				1, 2, 3,
				4, 5, 6,
			}),
			NewDense(3, 3, []DType{
				9, 12, 15,
				19, 26, 33,
				29, 40, 51,
			}),
		},
    } {
		res := Mul(test.a, test.b)
		rows, cols := res.Dims()
		for i := range rows {
			for j := range cols {
				want := test.c.At(i, j)
				got  := res.At(i, j)
				if got != want {
					t.Errorf("unexpected values in function Mul on row %d col %d in test %d: got: %v want: %v",
						i, j, id, got, want)
				}
			}
		}
	}
}


func TestInv(t *testing.T) {
}


func TestAddScalar(t *testing.T) { }

func TestMulScalar(t *testing.T) { }
