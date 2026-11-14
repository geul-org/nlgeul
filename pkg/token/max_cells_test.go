//ff:func feature=token type=helper control=sequence
//ff:what TestMaxCells — the widest of the texts in cells, a Hangul letter counting two; none give 0
package token

import "testing"

func TestMaxCells(t *testing.T) {
	if got := maxCells([]string{"abc", "가나", "d"}); got != 4 {
		t.Fatal(got)
	}
	if got := maxCells([]string{"abcd", "가"}); got != 4 {
		t.Fatal(got)
	}
	if got := maxCells(nil); got != 0 {
		t.Fatal(got)
	}
}
