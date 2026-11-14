//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestSpacingTitleStrict — a new word before a title is followed by the title itself: 갑돌씨앗이 has 갑돌 씨 앗+이 and not 갑돌 씨앗+이, and every candidate with 갑돌 as its first word has 씨 as its second
package token

import "testing"

func TestSpacingTitleStrict(t *testing.T) {
	r := spacingReader(t, "words-spacing.tsv", "48")
	lines := branchLines(r, "갑돌씨앗이")
	if !lines["갑돌 씨[열쇠:체언] 앗+이[주격:이]"] || lines["갑돌 씨앗+이[주격:이]"] {
		t.Fatal(lines)
	}
	for _, b := range r.Contests("갑돌씨앗이")[0].Branches {
		if len(b.Words) > 1 && b.Words[0].Parts[0].Text == "갑돌" && b.Words[1].Parts[0].Text != "씨" {
			t.Fatal(b)
		}
	}
}
