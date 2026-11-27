//ff:func feature=data type=helper control=sequence
//ff:what formatValue — the number column of a character: an integer when whole, else six significant digits, and - when it has no number
package main

import (
	"math"
	"strconv"
)

func formatValue(t *ucdTable, r rune) string {
	n, ok := parseNumeric(t.Chars[r].Numeric)
	if !ok {
		return "-"
	}
	if n == math.Trunc(n) {
		return strconv.FormatInt(int64(n), 10)
	}
	return strconv.FormatFloat(n, 'g', 6, 64)
}
