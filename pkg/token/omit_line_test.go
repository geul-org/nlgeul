//ff:func feature=token type=helper control=sequence
//ff:what TestOmitLine — the count starts one cell after the boxes (width + 4 cells), where the brackets start
package token

import (
	"math/big"
	"testing"
)

func TestOmitLine(t *testing.T) {
	got := omitLine(10, big.NewInt(3))
	if got != "               (3개 생략됨)" || cellWidth(got[:15]) != 15 {
		t.Fatalf("%q", got)
	}
}
