package averager

import (
	"errors"
	"io"
)

// averager interface takes an average on something. Uses the 'initForTests' method
// to allow validation that the average is done in a correct manner
type averager interface {
	// postConstructInit the averager, taking some io.Reader which should be averaged. Should error on nil io.Reader.
	//
	// This method is required to be implemented in order to use the interface_test.go function interfaceTest
	postConstructInit(io.Reader) (averager, error)

	// Average io.Reader set in init.
	//
	// Return the average percent, as float.
	//
	// Return ErrNilReader if the reader in the implementation is nil.
	// Return ErrUninitiated if the averager hasn't called init before Average.
	// Return ErrEmptyReader if the reader contains no bytes.
	// Return ErrNoPercentages if the reader contains no percentages.
	Average() (float64, error)
}

var (
	ErrNilReader     = errors.New("reader is nil")
	ErrUninitiated   = errors.New("averager is uninitiated")
	ErrEmptyReader   = errors.New("reader contains no bytes")
	ErrNoPercentages = errors.New("there are no percentages to average")
)
