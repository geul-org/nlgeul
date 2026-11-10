//ff:func feature=token type=reader control=sequence
//ff:what TestReaderCuts — with the word table: 물을마신다 is cut after 물을 in its best reading; an empty sentence has no lines
package token

import (
	"strings"
	"testing"
)

func TestReaderCuts(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	got := r.Cuts("물을마신다")
	if len(got) != 1 || !strings.HasPrefix(got[0], "물을마신다\t2,5") {
		t.Fatal(got)
	}
	if r.Cuts("") != nil {
		t.Fatal("empty")
	}
}
