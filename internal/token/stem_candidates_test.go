//ff:func feature=token type=reader control=sequence
//ff:what TestStemCandidates — the letters before the tail; the syllable before fusing; a closed-set word; an atom; the word of a fused form, known under the form's name without a table and under the table's name with one; a misspelt stem is the right form, judged by the kinds of the misspelt line as a misspelt form and by nothing else; a stem that a tail reads as a number is named so with the number's verdict
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStemCandidates(t *testing.T) {
	sc := scanOf("마신 그는 ab는 이게")
	if stem, _, v, _ := stemIn(sc, 0, tail{start: 2, base: '시'}, false); stem != "마시" || v > 0 {
		t.Fatal(stem, v)
	}
	if stem, name, v, _ := stemIn(sc, 3, tail{start: 4}, false); stem != "그" || name != "관형사:그" || v != 1 {
		t.Fatal(stem, name, v)
	}
	if stem, name, v, _ := stemIn(sc, 6, tail{start: 8}, false); stem != "ab" || v <= 0 || name == "" {
		t.Fatal(stem, name, v)
	}
	if stem, name, v, _ := stemIn(sc, 10, tail{start: 12, shown: "이것", label: "줄임:이게"}, false); stem != "이것" || name != "줄임:이게" || v != 1 {
		t.Fatal(stem, name, v)
	}
	ws := wordsScan("물게", layerCost)
	if stem, name, v, _ := stemIn(ws, 0, tail{start: 2, shown: "물", label: "줄임:x", pieces: []piece{{tail: 2}}}, false); stem != "물" || name == "줄임:x" || v <= 0 {
		t.Fatal(stem, name, v)
	}
	ts := typoScan("맞겨", layerCost)
	if stem, name, v, _ := stemIn(ts, 0, tail{start: 2, base: '기', shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{tail: 8}}}, false); stem != "맡기" || name != "열쇠:동사" || v <= 0 || v < 0.54 || v > 0.56 {
		t.Fatal("a misspelt stem is judged as a misspelt form", stem, name, v)
	}
	if _, name, v, _ := stemIn(ts, 0, tail{start: 2, base: '기', shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{tail: 2}}}, false); v > 0 || name != "" {
		t.Fatal("and by nothing else: a verb does not take a particle", name, v)
	}
	if stem, name, v, _ := stemIn(scanOf("세명"), 0, tail{start: 1, as: numName}, false); stem != "세" || name != numName || v != 1 {
		t.Fatal(stem, name, v)
	}
	bs := typoScan("입력/역활을", layerCost)
	eul := []piece{{"을", "목적격:을", 1, -1, rule.TailCase, false, ""}}
	stem, name, v, wrote := stemIn(bs, 0, tail{start: 5, pieces: eul, boundTypo: true}, false)
	if stem != "입력/역할" || name != "대비:/" || v <= 0 || wrote != "입력/역활" {
		t.Fatal(stem, name, v, wrote)
	}
	stem, name, v, wrote = stemIn(bs, 0, tail{start: 5, pieces: eul}, false)
	if stem != "입력/역활" || name != "" || v != 0 || wrote != "" {
		t.Fatal(stem, name, v, wrote)
	}
	right := typoScan("입력/출력을", layerCost)
	if stem, name, v, wrote = stemIn(right, 0, tail{start: 5, pieces: eul, boundTypo: true}, false); stem != "입력/출력" || name != "대비:/" || v <= 0 || wrote != "" {
		t.Fatal(stem, name, v, wrote)
	}
	if stem, _, _, wrote = stemIn(bs, 0, tail{start: 2, shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{tail: 8}}}, false); stem != "맡기" || wrote != "맞기" {
		t.Fatal(stem, wrote)
	}
}
