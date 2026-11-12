//ff:func feature=token type=rule control=sequence
//ff:what TestAddTypoKeys — with a word table the closed table's wrong way to write a closed-set word (꺼 for 거) is among the misspelt forms, with the kind 거 has as a dependent noun, though the table has no such line; a table that has the same line too gives the one entry, not two; with no word table the index has none
package token

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddTypoKeys(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	if got := ix.typos["꺼"]; len(got) != 1 || got[0] != (typo{"거", kindNominal}) {
		t.Fatal("the closed table's line", ix.typos["꺼"])
	}
	if len(ix.typos["역활"]) != 1 {
		t.Fatal("the table's own lines stay", ix.typos["역활"])
	}
	both := filepath.Join(t.TempDir(), "both.tsv")
	if err := os.WriteFile(both, []byte("거\t명사\t\t6000\n꺼\t명사\t\t0\t거\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	joined, err := newIndex(Options{Words: both})
	if err != nil {
		t.Fatal(err)
	}
	if got := joined.typos["꺼"]; len(got) != 1 || got[0] != (typo{"거", kindNominal}) {
		t.Fatal("the same two forms are one entry", got)
	}
	bare, err := newIndex(Options{})
	if err != nil {
		t.Fatal(err)
	}
	if len(bare.typos) != 0 {
		t.Fatal("no word table, no misspelt forms", bare.typos)
	}
}
