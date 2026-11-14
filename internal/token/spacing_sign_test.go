//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingSign — a word of the list after a sign between letters is no cue to show the letters before it (a sign shows the boundary already): 솔·겸총무 reads the same with article 45 on as off; with article 48 on as well 총무 is a title and 솔·겸 총무 appears (4 candidates, then 6)
package token

import (
	"slices"
	"testing"
)

func TestSpacingSign(t *testing.T) {
	off, one, two := spacingReader(t, "words-spacing.tsv"), spacingReader(t, "words-spacing.tsv", "45"), spacingReader(t, "words-spacing.tsv", "45", "48")
	if !slices.Equal(off.Lines("솔·겸총무", false), one.Lines("솔·겸총무", false)) {
		t.Fatal("article 45 alone changes it")
	}
	if len(off.Contests("솔·겸총무")[0].Branches) != 4 || len(two.Contests("솔·겸총무")[0].Branches) != 6 || !branchLines(two, "솔·겸총무")["솔·겸 총무[열쇠:체언]"] {
		t.Fatal(branchLines(two, "솔·겸총무"))
	}
}
