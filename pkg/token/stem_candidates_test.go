//ff:func feature=token type=reader control=sequence
//ff:what TestStemCandidates — the letters before the tail; the syllable before fusing; a closed-set word; an atom; the word of a fused form, known under the form's name without a table and under the table's name with one; a misspelt stem is the right form, judged by the kinds of the misspelt line as a misspelt form and by nothing else; a stem that a tail reads as a number is named so with the number's verdict; a stem two nominals of the table make, bound by 사이시옷 (전셋값), is that nominal (사이시옷:ㅅ) and the new word of the same letters, and the first alone when the stem is too long for a new word and not the whole unit (대한민국임시정붓수립기념일); a stem too long for a new word with no such reading has no candidate; an auxiliary verb's stem with a predicate tail right after a word that ends in the linking 어/아 stands higher, fully when seen enough (the rebuttal that it is one syllable long does not hold), under the same names
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStemCandidates(t *testing.T) {
	driveAll(t)
	far := streamReader(t, "words-far.tsv")
	sc := &scan{ix: far.ix, s: []rune("가나다라마바사아자차카타했다")}
	sc.prepare()
	wrong := tail{start: 12, wrote: "wrote"}
	if got := sc.stemCandidates(0, wrong, false, false); got != nil {
		t.Fatal("a stem written wrong that is too long for a new word", got)
	}
	if got := sc.stemCandidates(0, wrong, true, false); len(got) != 1 || got[0].wrote != "wrote" {
		t.Fatal("the same stem as the whole unit", got)
	}
	if got := sc.stemCandidates(2, wrong, false, false); len(got) != 1 {
		t.Fatal("the same stem within ten tokens", got)
	}
	sc = scanOf("마신 그는 ab는 이게")
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
	if stem, name, v, _ := stemIn(ts, 0, tail{start: 2, base: '기', shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{tail: 8}}}, false); stem != "맡기" || name != "열쇠:동사" || math.Abs(v-(2/(1+forceTypo)-1)) > 1e-9 {
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
	aux := wordsScan("해줘", layerCost)
	aux.ix.lex["주"] = word{kinds: kindVerb, count: 10000, by: [6]int{0, 10000}}
	jwo := tail{start: 2, base: '주', pieces: []piece{{text: "(어)", names: "맺음:어(합침)", verdict: 1, tail: rule.TailPred}}}
	stem, apart, low, _ := stemIn(aux, 1, jwo, false)
	aux.connEnds[1] = true
	if _, onto, v, _ := stemIn(aux, 1, jwo, false); stem != "주" || onto != apart || low >= 1 || v != 1 {
		t.Fatal("an auxiliary's stem right after 어/아", stem, apart, low, onto, v)
	}
	sc = &scan{ix: streamReader(t, "words-saisiot.tsv").ix, s: []rune("전셋값이")}
	sc.prepare()
	i := []piece{{"이", "주격:이", 1, -1, rule.TailCase, false, ""}}
	got := sc.stemCandidates(0, tail{start: 3, pieces: i}, false, false)
	if len(got) != 2 || got[0].stem != "전셋값" || got[0].name != saisiotName || got[0].kind != kindNominal || got[0].v <= 0 || got[1].stem != "전셋값" || got[1].name != "" || got[1].v != 0 {
		t.Fatal("a stem of 사이시옷 and the new word of the same letters", got)
	}
	sc = &scan{ix: far.ix, s: []rune("대한민국임시정붓수립기념일을")}
	sc.prepare()
	if got = sc.stemCandidates(0, tail{start: 13, pieces: eul}, false, false); len(got) != 1 || got[0].name != saisiotName || got[0].stem != "대한민국임시정붓수립기념일" {
		t.Fatal("a stem of 사이시옷 too long for a new word stands alone", got)
	}
	if got = sc.stemCandidates(0, tail{start: 13, pieces: eul}, true, false); len(got) != 2 || got[1].name != "" {
		t.Fatal("with its new word as the whole unit", got)
	}
	sc = &scan{ix: far.ix, s: []rune("가나다라마바사아자차카타했다")}
	sc.prepare()
	if got = sc.stemCandidates(0, tail{start: 12}, false, false); len(got) != 0 {
		t.Fatal("a stem too long for a new word, of no 사이시옷", got)
	}
}
