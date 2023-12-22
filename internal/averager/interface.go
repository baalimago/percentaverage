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
	// Return NilReaderError if the reader in the implementation is nil.
	// Return UninitiatedError if the averager hasn't called init before Average.
	// Return EmptyReaderError if the reader contains no bytes.
	// Return NoPercentagesError if the reader contains no percentages.
	Average() (float64, error)
}

var (
	NilReaderError     = errors.New("reader is nil")
	UninitiatedError   = errors.New("averager is uninitiated")
	EmptyReaderError   = errors.New("reader contains no bytes")
	NoPercentagesError = errors.New("there are no percentages to average")
)
