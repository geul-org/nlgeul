//ff:func feature=token type=reader control=sequence
//ff:what TestAddAlt — a changed form is kept with its stem; for ㄷ and ㅅ the stem's own last syllable is kept under the changed one, once however many stems share it, and apart for each kind
package token

import "testing"

func TestAddAlt(t *testing.T) {
	ix := &index{alts: map[string][]alt{}, irrEnds: map[rune][]alt{}}
	ix.addAlt("도우", "돕", irrB)
	ix.addAlt("들", "듣", irrD)
	ix.addAlt("깨들", "깨듣", irrD)
	ix.addAlt("지", "짓", irrS)
	if len(ix.alts) != 4 || len(ix.irrEnds) != 2 || len(ix.irrEnds['들']) != 1 || ix.irrEnds['들'][0].stem != "듣" || ix.irrEnds['지'][0].kind != irrS {
		t.Fatal(ix.alts, ix.irrEnds)
	}
	ix.addAlt("들", "듣", irrS)
	if len(ix.irrEnds['들']) != 2 {
		t.Fatal("another kind under the same syllable", ix.irrEnds['들'])
	}
}
