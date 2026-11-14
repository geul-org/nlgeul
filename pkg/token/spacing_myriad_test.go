//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingMyriad — article 44 (the only rule on): 삼억오천만원을 is parted after 삼억 (삼억 오천만+원+을, the first word a new word); 오천만원을 has one group and reads the same on or off
package token

import (
	"slices"
	"testing"
)

func TestSpacingMyriad(t *testing.T) {
	on, off := spacingReader(t, "words-spacing.tsv", "44"), spacingReader(t, "words-spacing.tsv")
	if !branchLines(on, "삼억오천만원을")["삼억 오천만[수]+원[단위:원]+을[목적격:을]"] || branchLines(off, "삼억오천만원을")["삼억 오천만[수]+원[단위:원]+을[목적격:을]"] {
		t.Fatal(branchLines(on, "삼억오천만원을"))
	}
	if !slices.Equal(on.Lines("오천만원을", false), off.Lines("오천만원을", false)) {
		t.Fatal("one group")
	}
}
