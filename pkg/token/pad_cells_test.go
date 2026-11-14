//ff:func feature=token type=helper control=sequence
//ff:what TestPadCells — spaces after the text fill it to the width in cells (a Hangul letter is two), and a text as wide or wider is left as it is
package token

import "testing"

func TestPadCells(t *testing.T) {
	if got := padCells("가a", 6); got != "가a   " || cellWidth(got) != 6 {
		t.Fatalf("%q", got)
	}
	if got := padCells("가a", 3); got != "가a" {
		t.Fatalf("%q", got)
	}
	if got := padCells("가a", 1); got != "가a" {
		t.Fatalf("%q", got)
	}
}
