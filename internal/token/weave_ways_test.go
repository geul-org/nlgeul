//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveWays — the lines of spans that chain from where the portion begins to where it ends, counted: two in the made-up portion (the line of spans that fit without chaining is not one), and how many (span, set) keys were counted
package token

import "testing"

func TestWeaveWays(t *testing.T) {
	w, _ := oddWeave()
	if got := w.ways(); got.Int64() != 2 || w.sets != 6 {
		t.Fatal(got, w.sets)
	}
}
