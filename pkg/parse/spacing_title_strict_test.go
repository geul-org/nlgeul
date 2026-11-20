//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestSpacingTitleStrict — a new word before a title is followed by the title itself: 갑돌씨앗이 has 갑돌 씨 앗+이 and not 갑돌 씨앗+이, and every candidate with 갑돌 as its first word has 씨 as its second (branchLines)
package parse

import "testing"

func TestSpacingTitleStrict(t *testing.T) {
	r := spacingReader(t, "words-spacing.tsv", "48")
	lines := branchLines(r, "갑돌씨앗이")
	if !lines["갑돌 씨[열쇠:체언] 앗+이[주격:이]"] || lines["갑돌 씨앗+이[주격:이]"] {
		t.Fatal(lines)
	}
	for line := range lines {
		if words := lineWords(line); len(words) > 1 && wordStem(words[0]) == "갑돌" && wordStem(words[1]) != "씨" {
			t.Fatal(line)
		}
	}
}
