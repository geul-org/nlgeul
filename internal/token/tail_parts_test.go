//ff:func feature=token type=helper control=sequence
//ff:what TestTailParts — the parts after the stem; all of the parts when the word has no stem; none when it has no part
package token

import "testing"

func TestTailParts(t *testing.T) {
	if got := tailParts(spacingWord("씨/열쇠:체언", "가/주격:가", "도/보조사:도")); len(got) != 2 || got[0].Text != "가" {
		t.Fatal(got)
	}
	if got := tailParts(spacingWord("-가/주격:가", "-도/보조사:도")); len(got) != 2 {
		t.Fatal("no stem", got)
	}
	if got := tailParts(Eojeol{}); len(got) != 0 {
		t.Fatal("no part", got)
	}
}
