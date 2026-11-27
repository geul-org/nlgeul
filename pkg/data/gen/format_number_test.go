//ff:func feature=data type=helper control=sequence
//ff:what TestFormatNumber — a whole number has no decimal point, a fraction gives its shortest decimal
package main

import "testing"

func TestFormatNumber(t *testing.T) {
	assertEq(t, formatNumber(3), "3")
	assertEq(t, formatNumber(0.5), "0.5")
	assertEq(t, formatNumber(1.0/3), "0.3333333333333333")
}
