//ff:func feature=data type=helper control=sequence
//ff:what TestFormatValue — a whole number is an integer, a fraction has six significant digits, no number is a dash
package main

import "testing"

func TestFormatValue(t *testing.T) {
	tbl := loadSample(t)
	assertEq(t, formatValue(tbl, 0x2460), "1")
	assertEq(t, formatValue(tbl, 0xBD), "0.5")
	assertEq(t, formatValue(tbl, 0x2153), "0.333333")
	assertEq(t, formatValue(tbl, 0x41), "-")
}
