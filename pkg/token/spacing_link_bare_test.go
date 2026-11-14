//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingLinkBare — article 45 when the word table has no 겸: 반장겸총무다 has 반장 겸 총무+(이)+다 with 겸 a new word with no name, and without the rule it has no word between 반장 and 총무
package token

import "testing"

func TestSpacingLinkBare(t *testing.T) {
	const line = "반장[열쇠:체언] 겸 총무[열쇠:체언]+(이)[서술격:이(빠짐)]+다[맺음:다]"
	on, off := spacingReader(t, "words-spacing-bare.tsv", "45"), spacingReader(t, "words-spacing-bare.tsv")
	if !branchLines(on, "반장겸총무다")[line] || branchLines(off, "반장겸총무다")[line] {
		t.Fatal(branchLines(on, "반장겸총무다"))
	}
}
