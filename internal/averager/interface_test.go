package averager

import (
	"bytes"
	"errors"
	"testing"
)

type retVal struct {
	perc float64
	err  error
}

// interfaceTest ensures that the implementation of averager
// upholds the functional requirements required
func interfaceTest(t *testing.T, target averager) {
	t.Helper()
	t.Run("it should NilWriterError on nil Writer", func(t *testing.T) {
		_, got := target.postConstructInit(nil)
		if !errors.Is(got, ErrNilReader) {
			t.Fatalf("expected NilWriterError, got: %v", got)
		}
	})

	t.Run("it should error on non-initiated averager", func(t *testing.T) {
		_, got := target.Average()
		if !errors.Is(got, ErrUninitiated) {
			t.Fatalf("expected ErrUninitiated, got: %v", got)
		}
	})

	t.Run("happy tests", func(t *testing.T) {
		testCases := []struct {
			desc  string
			given string
			want  retVal
		}{
			{
				desc:  "no content",
				given: "",
				want: retVal{
					perc: 0,
					err:  ErrEmptyReader,
				},
			},
			{
				desc:  "no percentages",
				given: "well, theres no percentages here at least",
				want: retVal{
					perc: 0,
					err:  ErrNoPercentages,
				},
			},
			{
				desc:  "correct average on same input",
				given: `10.00% 10.00% 10.00%`,
				want: retVal{
					perc: float64(10),
					err:  nil,
				},
			},
			{
				desc: "weird whitespaces",
				given: `10.00% 10.00%10.00%   \t \n \t

10%
        10%`,
				want: retVal{
					perc: float64(10),
					err:  nil,
				},
			},
			{
				desc:  "correct average on negative numbers",
				given: `100.00% -100.00% 0.50% -0.10%`,
				want: retVal{
					perc: float64(0.4) / float64(4),
					err:  nil,
				},
			},
			{
				desc:  "handle periods and commas",
				given: `10.0% 10% 10,0%`,
				want: retVal{
					perc: float64(30) / float64(3),
					err:  nil,
				},
			},
			{
				desc:  "ignore non-percentages",
				given: `100,00% 0.0% 0,0% 100.00 100,00 %%% abc.def% abc,def% 🎅%`,
				want: retVal{
					perc: float64(100) / float64(3),
					err:  nil,
				},
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				target, err := target.postConstructInit(bytes.NewBufferString(tC.given))
				if err != nil {
					t.Fatalf("failed to init averager: %v", err)
				}
				got, err := target.Average()
				if !errors.Is(err, tC.want.err) {
					t.Fatalf("expected: %v, got: %v", tC.want.err, err)
				}
				if got != tC.want.perc {
					t.Fatalf("expected: %v, got: %v", tC.want.perc, got)
				}
			})
		}
	})
}
