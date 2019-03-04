// check.go
// Paul Schuster
// 030119

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"mul/evalbool"
)

var checkRegexp = regexp.MustCompile(`(\d*)(\w|\&|\||\(|\)|\^|!)`)

// Check for the eval condition
func check(eval string, h []byte) bool {
	cards := make(map[byte]int)

	// Count the cards in the hand
	for _, v := range h {
		_, ok := cards[v]
		if ok {
			cards[v]++
		} else {
			cards[v] = 1
		}
	}

	// Parse the eval condition
	e := checkRegexp.FindAllStringSubmatch(eval, -1)
	exp := make([]byte, len(e))

	// rewrite the eval condition with the found truth values
	for k, v := range e {
		// get the first byte of the 2nd catch group
		b := []byte(v[2])[0]

		// if it is a opperation pass it thru
		if isOpp(b) {
			exp[k] = b
			continue
		}

		// check to see how many of the cards we want
		c, err := strconv.Atoi(v[1])
		if err != nil {
			c = 1
		}

		// see if the card was in the hand count
		num, ok := cards[b]
		if ok {
			// card was found is there enough?
			if num >= c {
				exp[k] = 'T'
			} else {
				exp[k] = 'F'
			}
		} else {
			// card not found
			exp[k] = 'F'
		}

		if *debug {
			fmt.Println(num, c, string(exp[k]))
		}
	}

	// evaluate the expression
	x, err := evalbool.Eval(exp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *debug {
		fmt.Println("\ntry", eval, cards, exp, x, e)
	}
	return x
}

// is the character an opperation?
func isOpp(b byte) bool {
	switch b {
	case '&', '|', '(', ')', '^', '!':
		return true
	}
	return false
}
