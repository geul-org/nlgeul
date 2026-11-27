//ff:func feature=data type=helper control=sequence
//ff:what TestContentKind — each kind is named once and the kinds are sorted and joined by +
package main

import "testing"

func TestContentKind(t *testing.T) {
	tbl := loadSample(t)
	assertEq(t, contentKind(tbl, "AA1"), "로마자+수")
	assertEq(t, contentKind(tbl, "(가)"), "기호+한글")
	assertEq(t, contentKind(tbl, ""), "")
}
