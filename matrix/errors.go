package matrix

// Borrowed from github.com/gonum/gonum/blob/master/mat/errors.go

// Error represents matrix handling errors.
// TODO: These errors can be recovered by Maybe wrappers.
type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	// ErrSliceLengthMismatch = Error{"matrix: input slice length mismatch"}
	// ErrShape               = Error{"matrix: dimension mismatch"}
	// ErrRowLength           = Error{"matrix: row length mismatch"}
	// ErrColLength           = Error{"matrix: col length mismatch"}


	// ErrIndexOutOfRange     = Error{"matrix: index out of range"}

	ErrNegativeDimension   = Error{"matrix: negative dimension"}
	ErrIndexOutOfRange     = Error{"matrix: index out of range"}
	ErrReuseNonEmpty       = Error{"matrix: reuse of non-empty matrix"}
	ErrRowAccess           = Error{"matrix: row index out of range"}
	ErrColAccess           = Error{"matrix: column index out of range"}
	ErrVectorAccess        = Error{"matrix: vector index out of range"}
	ErrZeroLength          = Error{"matrix: zero length in matrix dimension"}
	ErrRowLength           = Error{"matrix: row length mismatch"}
	ErrColLength           = Error{"matrix: col length mismatch"}
	ErrSquare              = Error{"matrix: expect square matrix"}
	ErrNormOrder           = Error{"matrix: invalid norm order for matrix"}
	ErrSingular            = Error{"matrix: matrix is singular"}
	ErrShape               = Error{"matrix: dimension mismatch"}
	ErrIllegalStride       = Error{"matrix: illegal stride"}
	ErrPivot               = Error{"matrix: malformed pivot list"}
	ErrTriangle            = Error{"matrix: triangular storage mismatch"}
	ErrTriangleSet         = Error{"matrix: triangular set out of bounds"}
	ErrBandwidth           = Error{"matrix: bandwidth out of range"}
	ErrBandSet             = Error{"matrix: band set out of bounds"}
	ErrDiagSet             = Error{"matrix: diagonal set out of bounds"}
	ErrSliceLengthMismatch = Error{"matrix: input slice length mismatch"}
	ErrNotPSD              = Error{"matrix: input not positive symmetric definite"}
	ErrFailedEigen         = Error{"matrix: eigendecomposition not successful"}
)

