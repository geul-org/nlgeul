//ff:func feature=data type=helper control=sequence
//ff:what TestRangeCategory — a code point in a range gives its category, the ends count, and one outside is Cn
package main

import "testing"

func TestRangeCategory(t *testing.T) {
	tbl := &ucdTable{Ranges: []charRange{{Lo: 10, Hi: 20, Category: "Lo"}}}
	assertEq(t, rangeCategory(tbl, 10), "Lo")
	assertEq(t, rangeCategory(tbl, 20), "Lo")
	assertEq(t, rangeCategory(tbl, 21), "Cn")
}
