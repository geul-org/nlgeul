//ff:func feature=token type=reader control=sequence
//ff:what TestReaderCuts — with the word table: 물을마신다 is cut after 물을 in its best reading, and its line holds the word ends of every reading the search keeps (as many as the beam holds), after one tab; an empty sentence has no lines
package token

import (
	"strings"
	"testing"
)

func TestReaderCuts(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	got := r.Cuts("물을마신다")
	if len(got) != 1 || !strings.HasPrefix(got[0], "물을마신다\t2,5;") || strings.Count(got[0], "\t") != 1 || strings.Count(got[0], ";") != beam-1 || !strings.HasSuffix(got[0], ";1,4,5") {
		t.Fatal(got)
	}
	if r.Cuts("") != nil {
		t.Fatal("empty")
	}
}
