package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

func process(pool, eval string) {
	// setup
	n := numberOfTries
	nWorkers := numberOfWorkers
	rand.Seed(time.Now().UTC().UnixNano())

	// create workers
	var wg sync.WaitGroup
	jobs := make(chan []byte, 100)
	results := make(chan bool, 100)

	for w := 1; w <= *nWorkers; w++ {
		wg.Add(1)
		// Handle the work
		go func(j <-chan []byte, r chan<- bool) {
			defer wg.Done()
			for job := range j {
				result := false
				for try := *handSize; try != *maxMul-1; try-- {
					rand.Shuffle(len(job), func(i, j int) {
						job[i], job[j] = job[j], job[i]
					})
					if *debug {
						fmt.Printf("try [%d/%d]\n",
							*handSize-try+1,
							*handSize-*maxMul+1,
						)
					}
					result = check(eval, job[:*handSize])
					if result {
						break
					}
				}
				r <- result
			}
		}(jobs, results)
	}

	// Spawn the decks
	go func() {
		deckTemplate := deckBuilder(pool)
		for i := 0; i < *n; i++ {
			deck := make([]byte, len(deckTemplate))
			copy(deck, deckTemplate)
			jobs <- deck
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	// output
	bar := pb.New(*n)
	if !*quiet {
		bar.ShowPercent = true
		bar.ShowCounters = true
		bar.ShowTimeLeft = true

		bar.Start()
	}

	// count stuff
	s := 0
	for r := range results {
		if r {
			s++
		}
		if !*quiet {
			bar.Increment()
		}
	}

	if !*quiet {
		bar.Finish()
	}

	if !*quiet {
		fmt.Printf("%d of %d successes [%f%%]\n", s, *n, (float64(s)/float64(*n))*100)
	} else {
		fmt.Printf("%d/%d\n", s, *n)
	}
}
