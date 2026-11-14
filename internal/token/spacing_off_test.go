//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestSpacingOff — a reader with no spacing rule on has none of the candidates the rules make (갑돌 씨+가), a reader with the rules that only mark a boundary reads each input exactly as one with none on (the same contests and word ends), and a reader with every rule on has every candidate the off reader has
package token

import (
	"slices"
	"testing"
)

func TestSpacingOff(t *testing.T) {
	marking := []string{"2", "41", "43다만", "46", "47", "47다만", "48다만", "49", "50"}
	off := spacingReader(t, "words-spacing.tsv")
	same := spacingReader(t, "words-spacing.tsv", marking...)
	all := spacingReader(t, "words-spacing.tsv", SpacingArticles...)
	for _, in := range []string{"갑돌씨가", "갑돌씨앗이", "뚜루뚜겸총무다", "뚜루뚜및운영", "뚜루뚜등을", "뚜루뚜중에", "삼억오천만원을", "오천만원을", "이백삼십개", "솔·겸총무", "마당에서놀았다", "읽어보았다", "고양이세마리", "아버지가방에들어가신다"} {
		if !slices.Equal(off.Cuts(in), same.Cuts(in)) || !slices.Equal(off.Lines(in, false), same.Lines(in, false)) {
			t.Fatal("rules that only mark change the reading", in)
		}
		on, was := branchLines(all, in), branchLines(off, in)
		for line := range was {
			if !on[line] {
				t.Fatal("a candidate left out", in, line)
			}
		}
	}
	if branchLines(off, "갑돌씨가")["갑돌 씨[열쇠:체언]+가[주격:가]"] || !branchLines(all, "갑돌씨가")["갑돌 씨[열쇠:체언]+가[주격:가]"] {
		t.Fatal("갑돌 씨+가")
	}
}
