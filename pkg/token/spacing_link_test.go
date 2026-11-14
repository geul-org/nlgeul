//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingLink — article 45: a word of the list after a new word makes it a word of its own: 뚜루뚜 겸 총무, 뚜루뚜 및 운영, 뚜루뚜 등+을 are candidates with the rule on and none with it off
package token

import "testing"

func TestSpacingLink(t *testing.T) {
	on, off := spacingReader(t, "words-spacing.tsv", "45"), spacingReader(t, "words-spacing.tsv")
	for in, line := range map[string]string{
		"뚜루뚜겸총무다": "뚜루뚜 겸[열쇠:체언] 총무[열쇠:체언]+(이)[서술격:이(빠짐)]+다[맺음:다]",
		"뚜루뚜및운영":  "뚜루뚜 및[열쇠:부사] 운영[열쇠:체언]",
		"뚜루뚜등을":   "뚜루뚜 등[열쇠:체언]+을[목적격:을]",
	} {
		if !branchLines(on, in)[line] || branchLines(off, in)[line] {
			t.Fatal(in, line, branchLines(on, in))
		}
	}
}
