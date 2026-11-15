//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingSign — a word of the list after a sign between letters is no cue to show the letters before it (a sign shows the boundary already): 솔·겸총무 reads the same with article 45 on as off; with article 48 on as well 총무 is a title and 솔·겸 총무 appears (4 candidates, then 6)
package token

import (
	"slices"
	"testing"
)

func TestSpacingSign(t *testing.T) {
	off, one, two := spacingReader(t, "words-spacing.tsv"), spacingReader(t, "words-spacing.tsv", "45"), spacingReader(t, "words-spacing.tsv", "45", "48")
	if !slices.Equal(off.BoxLines("솔·겸총무", false, 0), one.BoxLines("솔·겸총무", false, 0)) {
		t.Fatal("article 45 alone changes it")
	}
	if len(branchLines(off, "솔·겸총무")) != 4 || len(branchLines(two, "솔·겸총무")) != 6 || !branchLines(two, "솔·겸총무")["솔·겸 총무[열쇠:체언]"] {
		t.Fatal(branchLines(two, "솔·겸총무"))
	}
}
