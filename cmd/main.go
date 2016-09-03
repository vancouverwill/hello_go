package main

import (
	"flag"
	"hello_go/concurrency"

	"github.com/pkg/profile"
)

var profileEnabled = flag.Bool("profile", false, "should we profile")

func main() {
	flag.Parse()

	if *profileEnabled {
		// p := profile.Start(profile.CPUProfile, profile.MemProfile, profile.BlockProfile, profile.ProfilePath("."))
		defer profile.Start(profile.ProfilePath(".")).Stop()
		// defer profile.Start().Stop()
		// defer p.Stop()
	}

	// concurrency.UseLoopToClearChannel()
	// findLargestNumberOfZeros(205)
	// exampleBinaryOperations(195)
	// fmt.Println(findLargestNumberOfZeros(195))
	// fmt.Println(findLargestNumberOfZerosSurroundedByOnes(6))
	concurrency.RunAsynchronousTasks()
}
