//ff:func feature=token type=reader control=sequence
//ff:what TestStemOf — the letters before the tail; the syllable before fusing; a closed-set word; an atom; the word of a fused form, known under the form's name without a table and under the table's name with one; a stem that a tail reads as a number is named so with the number's verdict
package token

import "testing"

func TestStemOf(t *testing.T) {
	sc := scanOf("마신 그는 ab는 이게")
	if stem, _, v := sc.stemOf(0, tail{start: 2, base: '시'}, false); stem != "마시" || v > 0 {
		t.Fatal(stem, v)
	}
	if stem, name, v := sc.stemOf(3, tail{start: 4}, false); stem != "그" || name != "관형사:그" || v != 1 {
		t.Fatal(stem, name, v)
	}
	if stem, name, v := sc.stemOf(6, tail{start: 8}, false); stem != "ab" || v <= 0 || name == "" {
		t.Fatal(stem, name, v)
	}
	if stem, name, v := sc.stemOf(10, tail{start: 12, shown: "이것", label: "줄임:이게"}, false); stem != "이것" || name != "줄임:이게" || v != 1 {
		t.Fatal(stem, name, v)
	}
	ws := wordsScan("물게", layerCost)
	if stem, name, v := ws.stemOf(0, tail{start: 2, shown: "물", label: "줄임:x", pieces: []piece{{tail: 2}}}, false); stem != "물" || name == "줄임:x" || v <= 0 {
		t.Fatal(stem, name, v)
	}
	if stem, name, v := scanOf("세명").stemOf(0, tail{start: 1, as: numName}, false); stem != "세" || name != numName || v != 1 {
		t.Fatal(stem, name, v)
	}
}
