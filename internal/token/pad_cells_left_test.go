//ff:func feature=token type=helper control=sequence
//ff:what TestPadCellsLeft — spaces before the text fill it to the width in cells, and a text as wide or wider is left as it is
package token

import "testing"

func TestPadCellsLeft(t *testing.T) {
	if got := padCellsLeft("-4.38", 7); got != "  -4.38" {
		t.Fatalf("%q", got)
	}
	if got := padCellsLeft("가", 5); got != "   가" || cellWidth(got) != 5 {
		t.Fatalf("%q", got)
	}
	if got := padCellsLeft("-11.44", 3); got != "-11.44" {
		t.Fatalf("%q", got)
	}
}
