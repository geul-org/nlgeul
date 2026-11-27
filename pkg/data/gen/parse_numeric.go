//ff:func feature=data type=helper control=sequence
//ff:what parseNumeric — the numeric value of a UnicodeData.txt numeric column (an integer, a decimal or a fraction like 1/2), and whether there is one
package main

import (
	"strconv"
	"strings"
)

func parseNumeric(s string) (float64, bool) {
	if s == "" {
		return 0, false
	}
	num, den, isFraction := strings.Cut(s, "/")
	n, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, false
	}
	if !isFraction {
		return n, true
	}
	d, err := strconv.ParseFloat(den, 64)
	if err != nil || d == 0 {
		return 0, false
	}
	return n / d, true
}
