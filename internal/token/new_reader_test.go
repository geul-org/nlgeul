//ff:func feature=token type=reader control=sequence
//ff:what TestNew — a reader without a word table, one with the table of testdata, and an error for a table that cannot be read
package token

import "testing"

func TestNew(t *testing.T) {
	if r, err := New(Options{}); err != nil || r.ix.layers != layerNone {
		t.Fatal(err)
	}
	if r, err := New(Options{Words: "testdata/words.tsv"}); err != nil || r.ix.layers != layerCost {
		t.Fatal(err)
	}
	if _, err := New(Options{Words: "testdata/none.tsv"}); err == nil {
		t.Fatal("a missing word table accepted")
	}
}
