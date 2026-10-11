package matrix

import "fmt"

// Print the matrix to STDOUT. (The only side effect in this file)
func Print(m Matrix) {
	rows, cols := m.Dims()
	for i := range rows {
		for j := range cols {
			fmt.Printf("%f ", m.At(i, j))
		}
		fmt.Printf("\n")
	}
}

