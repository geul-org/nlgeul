//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingBound — article 42: a dependent noun of the list right after a new word makes it a word of its own: 뚜루뚜중에 has 뚜루뚜 중+에, and has not with the rule off
package token

import "testing"

func TestSpacingBound(t *testing.T) {
	const line = "뚜루뚜 중[열쇠:체언]+에[부사격:에]"
	on, off := spacingReader(t, "words-spacing.tsv", "42"), spacingReader(t, "words-spacing.tsv")
	if !branchLines(on, "뚜루뚜중에")[line] || branchLines(off, "뚜루뚜중에")[line] {
		t.Fatal(branchLines(on, "뚜루뚜중에"))
	}
}
