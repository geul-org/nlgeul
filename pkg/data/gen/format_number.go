//ff:func feature=data type=helper control=sequence
//ff:what formatNumber — a number as the table writes it in the value column: an integer when whole, else its shortest decimal
package main

import (
	"math"
	"strconv"
)

func formatNumber(n float64) string {
	if n == math.Trunc(n) {
		return strconv.FormatInt(int64(n), 10)
	}
	return strconv.FormatFloat(n, 'f', -1, 64)
}
