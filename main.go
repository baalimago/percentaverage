package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/baalimago/percentaverage/internal/averager"
)

var r = flag.Bool("r", false, "Set to true if you wish only the percentage as output, without '%'.")
var round = flag.Bool("round", false, "Set to true if you wish to round the output to the nearest integer.")

func main() {
	flag.Parse()
	regexpAverager := averager.NewRegexpAverager(os.Stdin)
	average, err := regexpAverager.Average()
	if err != nil {
		panic(fmt.Errorf("failed to average: %v", err))
	}

	averageStr := fmt.Sprintf("%v", average)
	if *round {
		averageStr = fmt.Sprintf("%.3f", average)
	}
	if *r {
		fmt.Printf("%v", averageStr)
	} else {
		fmt.Printf("%v%%", averageStr)
	}
}
