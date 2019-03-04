// main.go
// Paul Schuster
// 030119

package main

import (
	"flag"
	"fmt"
	"os"
)

// Ya config globals... IDK
var numberOfTries = flag.Int("perm", 10000, "number of trys")
var numberOfWorkers = flag.Int("t", 1, "Number of Threads")
var deckSize = flag.Int("deck", 60, "Deck size")
var handSize = flag.Int("hand", 7, "Hand size")
var maxMul = flag.Int("mul", 5, "Min after muligan hand size")
var debug = flag.Bool("v", false, "Debug/verbose mode")
var quiet = flag.Bool("q", false, "Quiet mode")

func main() {
	// pretty up the usage and other flag stuff
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s <deckstring> <logic>\n\n",
			os.Args[0],
		)
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(),
			"\nExample: %s '4a4b' 'a&b'\n",
			os.Args[0],
		)
		fmt.Fprintf(flag.CommandLine.Output(),
			"\t%s -perm 100000 -mul 5 '4a4b4c4d' '(a & b & (c ^ d)) | ((a ^ b) & c & d)'\n",
			os.Args[0],
		)
	}
	flag.Parse()

	a := flag.Args()
	if len(a) != 2 {
		flag.CommandLine.Usage()
		os.Exit(2)
	}
	pool := a[0]
	eval := a[1]

	// print the header
	if !*quiet {
		fmt.Printf("London muligan sim\n")
		fmt.Printf("Deck Size = %d\n", *deckSize)
		fmt.Printf("Hand Size = %d\n", *handSize)
		fmt.Printf("Min after muligan hand size = %d\n\n", *maxMul)
	}

	// do the thing that you do
	process(pool, eval)
}
