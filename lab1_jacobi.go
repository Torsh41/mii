package main

import (
    "fmt"
	// "rand"
	// "gonum.org/v1/gonum/mat"
	"github.com/Torsh/mii/matrix"
)

// TODO:
// 1) Jacobi method
// 2) Make a public interface (a library?)
// 3) use contracts
// 4) test suite
// 5) docs
// 6) build a package???

func main() {
    fmt.Println("HelloWorld!")
	a := makeDenseMatrix(5, 5, []float32{
		1,2,3,4,5,
		6,7,8,9,10,
		11,12,13,14,15,
		16,17,18,19,20,
		21,22,23,24,25,
	})
	a.print()
}
