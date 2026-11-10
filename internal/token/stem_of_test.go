//ff:func feature=token type=reader control=sequence
//ff:what TestStemOf — the letters before the tail; the syllable before fusing; a closed-set word; an atom; the word of a fused form, known under the form's name without a table and under the table's name with one
package token

import "testing"

func TestStemOf(t *testing.T) {
	sc := scanOf("마신 그는 ab는 이게")
	if stem, _, v := sc.stemOf(0, tail{start: 2, base: '시'}); stem != "마시" || v > 0 {
		t.Fatal(stem, v)
	}
	if stem, name, v := sc.stemOf(3, tail{start: 4}); stem != "그" || name != "관형사:그" || v != 1 {
		t.Fatal(stem, name, v)
	}
	if stem, name, v := sc.stemOf(6, tail{start: 8}); stem != "ab" || v <= 0 || name == "" {
		t.Fatal(stem, name, v)
	}
	if stem, name, v := sc.stemOf(10, tail{start: 12, shown: "이것", label: "줄임:이게"}); stem != "이것" || name != "줄임:이게" || v != 1 {
		t.Fatal(stem, name, v)
	}
	ws := wordsScan("물게", layerCost)
	if stem, name, v := ws.stemOf(0, tail{start: 2, shown: "물", label: "줄임:x", pieces: []piece{{tail: 2}}}); stem != "물" || name == "줄임:x" || v <= 0 {
		t.Fatal(stem, name, v)
	}
}
