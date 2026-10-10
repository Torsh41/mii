package matrix

import (
	"fmt"
)


type DenseMatrix struct {
	w,h int
	_data []float32
}

func makeDenseMatrix(w int, h int, data []float32) DenseMatrix {
	return DenseMatrix{w, h, data}
}

func (m DenseMatrix)print() {
	for i := 0; i < m.h; i++ {
		for j := 0; j < m.w; j++ {
			fmt.Printf("%f ", m._data[j + m.w*i])
		}
	}
}

