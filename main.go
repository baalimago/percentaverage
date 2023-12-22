package main

import (
	"fmt"
	"os"

	"github.com/baalimago/percentaverage/internal/averager"
)

func main() {
	regexpAverager := averager.NewRegexpAverager(os.Stdin)
	average, err := regexpAverager.Average()
	if err != nil {
		panic(fmt.Errorf("failed to average: %v", err))
	}
	fmt.Printf("%v%%", average)
}
