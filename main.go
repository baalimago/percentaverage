package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/baalimago/percentaverage/internal/averager"
)

var r = flag.Bool("r", false, "Set to true if you wish only the percentage as output, without '%'.")

func main() {
	flag.Parse()
	regexpAverager := averager.NewRegexpAverager(os.Stdin)
	average, err := regexpAverager.Average()
	if err != nil {
		panic(fmt.Errorf("failed to average: %v", err))
	}
	if *r {
		fmt.Printf("%v", average)
	} else {
		fmt.Printf("%v%%", average)
	}
}
