//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingUnit — article 43: a unit of the unit table right after a Hangul numeral that no table lists is parted from it: 이백삼십개 has 이백삼십 개 with the rule on, and has not with it off
package token

import "testing"

func TestSpacingUnit(t *testing.T) {
	const line = "이백삼십 개[열쇠:체언]"
	on, off := spacingReader(t, "words-spacing.tsv", "43"), spacingReader(t, "words-spacing.tsv")
	if !branchLines(on, "이백삼십개")[line] || branchLines(off, "이백삼십개")[line] {
		t.Fatal(branchLines(on, "이백삼십개"))
	}
}
