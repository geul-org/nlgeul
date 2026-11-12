//ff:func feature=token type=rule control=sequence
//ff:what TestUnlisted — a stem no table lists has a new word's verdict; with a word table it does not stand as a word the table (사과) or a closed set (거) knows only as a nominal before a predicate tail, though it does before particles and as a predicate's stem too (물), nor when it ends in ㅆ before markers (곰졌), judged as it was before a piece was fused into it (곰지); without a word table all of these stand; after the search a new word whose letters read as markers after a boundary does not stand whole (에), and stands with markers after it only without a word table
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestUnlisted(t *testing.T) {
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	ne := []piece{{text: "네", names: "맺음:네", verdict: 1, tail: rule.TailPred}}
	sc := scanOf("곰다")
	if v := sc.unlisted(0, 2, "곰", tail{start: 1, pieces: da}); v != sc.ix.unknown {
		t.Fatal("a new word", v)
	}
	if v := scanOf("거네").unlisted(0, 2, "거", tail{start: 1, pieces: ne}); v != sc.ix.unknown {
		t.Fatal("without a word table a closed-set word before a predicate tail is not asked", v)
	}
	if v := scanOf("사과다").unlisted(0, 3, "사과", tail{start: 2, pieces: da}); v != sc.ix.unknown {
		t.Fatal("without a word table the tables know no 사과", v)
	}
	if v := scanOf("곰졌다").unlisted(0, 3, "곰졌", tail{start: 2, pieces: da}); v != sc.ix.unknown {
		t.Fatal("without a word table a stem in ㅆ is not told from 맛있", v)
	}
	if v := wordsScan("거네", layerCost).unlisted(0, 2, "거", tail{start: 1, pieces: ne}); v > 0 {
		t.Fatal("거, a dependent noun of a closed set, before a predicate tail", v)
	}
	noun := wordsScan("사과다", layerCost)
	if v := noun.unlisted(0, 3, "사과", tail{start: 2, pieces: da}); v > 0 {
		t.Fatal("사과, a noun of the table, before a predicate tail", v)
	}
	if v := noun.unlisted(0, 3, "사과", tail{start: 2, pieces: []piece{{text: "다", verdict: 1, tail: rule.TailCase}}}); v != noun.ix.unknown {
		t.Fatal("the same noun before particles", v)
	}
	if v := wordsScan("물다", layerCost).unlisted(0, 2, "물", tail{start: 1, pieces: da}); v != noun.ix.unknown {
		t.Fatal("물, a noun and a verb of the table", v)
	}
	coda := wordsScan("곰졌다", layerCost)
	if v := coda.unlisted(0, 3, "곰졌", tail{start: 2, pieces: da}); v > 0 {
		t.Fatal("곰졌, a new word that ends in ㅆ, before 다", v)
	}
	past := []piece{{text: "(었)", names: "선어말:ㅆ", verdict: 1, tail: rule.TailPred}, da[0]}
	if v := coda.unlisted(0, 3, "곰지", tail{start: 2, base: '지', pieces: past}); v != coda.ix.unknown {
		t.Fatal("곰지 with 었 fused into its last syllable: the stem as it was has no ㅆ", v)
	}
	if v := coda.unlisted(0, 3, "곰졌다", tail{start: 3}); v != coda.ix.unknown {
		t.Fatal("a whole new word with no marker", v)
	}

	seo := []piece{{text: "서", names: "부사격:서", verdict: 1, tail: rule.TailCase}}
	bare := scanOf("마당에서")
	bare.search()
	if v := bare.unlisted(2, 3, "에", tail{start: 3}); v > 0 {
		t.Fatal("에 after a boundary, read as a marker, as a whole new word", v)
	}
	if v := bare.unlisted(2, 4, "에", tail{start: 3, pieces: seo}); v != bare.ix.unknown {
		t.Fatal("without a word table a stem with markers after it stands though its letters read as markers", v)
	}
	lexed := wordsScan("마당에서", layerCost)
	lexed.search()
	if v := lexed.unlisted(2, 4, "에", tail{start: 3, pieces: seo}); v > 0 {
		t.Fatal("with a word table it does not", v)
	}
	if v := lexed.unlisted(0, 4, "마당에", tail{start: 3, pieces: seo}); v != lexed.ix.unknown {
		t.Fatal("the unit's first word is not marked", v)
	}
}
