package matrix

import (
	"fmt"
)


type DenseMatrix struct {
	_data []float32
	w,h int
}

func NewDenseMatrix(data []float32, w int, h int) DenseMatrix {
	return DenseMatrix{data, w, h}
}

func (m *DenseMatrix) Print() {
	for i := 0; i < m.h; i++ {
		for j := 0; j < m.w; j++ {
			fmt.Printf("%f ", m._data[j + m.w*i])
		}
		fmt.Printf("\n")
	}
}

func DoSomething() {
	fmt.Printf("Hello World!")
}
