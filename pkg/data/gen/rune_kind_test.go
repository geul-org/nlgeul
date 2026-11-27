//ff:func feature=data type=helper control=sequence
//ff:what TestRuneKind — each kind of character gives its name, a character in a range is a Han character, and a sign is 기호
package main

import "testing"

func TestRuneKind(t *testing.T) {
	tbl := loadSample(t)
	assertEq(t, runeKind(tbl, '1'), "수")
	assertEq(t, runeKind(tbl, 0x1100), "한글")
	assertEq(t, runeKind(tbl, 0x3041), "가나")
	assertEq(t, runeKind(tbl, 0x5000), "한자")
	assertEq(t, runeKind(tbl, 'A'), "로마자")
	assertEq(t, runeKind(tbl, 0x3B1), "그리스 글자")
	assertEq(t, runeKind(tbl, 0xFB50), "그 밖의 글자")
	assertEq(t, runeKind(tbl, '('), "기호")
}
