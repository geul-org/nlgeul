//ff:func feature=data type=helper control=sequence
//ff:what TestCategoryOf — a listed character gives its category, one in a range gives the range's, and an unlisted one is Cn
package main

import "testing"

func TestCategoryOf(t *testing.T) {
	tbl := loadSample(t)
	assertEq(t, categoryOf(tbl, 0x28), "Ps")
	assertEq(t, categoryOf(tbl, 0x5000), "Lo")
	assertEq(t, categoryOf(tbl, 0x0378), "Cn")
}
