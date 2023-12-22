package averager

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
)

// regexpAverager uses regexp pattern to find and average all the percentages in a io.Reader
type regexpAverager struct {
	io.Reader
	isInited *atomic.Bool
}

var r = regexp.MustCompile(`-?\d+[,.]*\d*%`)

// NewRegexpAverager returns an regexpAverager and panics if the postConstructInit fails
func NewRegexpAverager(r io.Reader) averager {
	tmp := &regexpAverager{}
	ret, err := tmp.postConstructInit(r)
	if err != nil {
		panic(err)
	}
	return ret
}

// postConstructInit by checking that the reader isn't nil, then returning a instanced
// version of *regexpAverager
func (a *regexpAverager) postConstructInit(r io.Reader) (averager, error) {
	if r == nil {
		return nil, ErrNilReader
	}
	b := atomic.Bool{}
	b.Store(true)
	a = &regexpAverager{
		Reader:   r,
		isInited: &b,
	}
	return a, nil
}

// Average all the percentages by scanning through the Reader using regexp pattern and then iterating through the matches
//
// Returns the precentage as a float on success, returns error when failing according to the averager interface
func (a *regexpAverager) Average() (float64, error) {
	if a == nil || a.isInited == nil || !a.isInited.Load() {
		return 0, ErrUninitiated
	}
	b, err := io.ReadAll(a)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	if len(b) == 0 {
		return 0, ErrEmptyReader
	}

	percByteMatches := r.FindAll(b, -1)
	if len(percByteMatches) == 0 {
		return 0, ErrNoPercentages
	}
	var sum float64
	for _, percByteSlice := range percByteMatches {
		percStr := string(percByteSlice)
		// We know that last character is '%', trim it
		percStr = strings.TrimRight(percStr, "%")
		percStr = strings.ReplaceAll(percStr, ",", ".")
		percFloat, err := strconv.ParseFloat(percStr, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse percString: %w", err)
		}

		sum += percFloat
	}
	return sum / float64(len(percByteMatches)), nil
}
