package main

import (
	"regexp"
	"strconv"
)

var poolRegexp = regexp.MustCompile(`(\d*)(\w)`)

// using the template make a proto-deck
func deckBuilder(pool string) []byte {
	d := make([]byte, *deckSize)

	// parse the deck template
	p := poolRegexp.FindAllStringSubmatch(pool, -1)
	if p == nil {
		return d
	}

	// fill the deck
	c := 0
	for _, v := range p {
		i := 0
		if v[1] == "" {
			i = 1
		} else {
			i, _ = strconv.Atoi(v[1])
		}

		end := c + i
		for j := c; j < end; j++ {
			t := []byte(v[2])[0]
			d[j] = t
			c++
		}
	}
	return d
}
