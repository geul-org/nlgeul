//ff:func feature=token type=rule control=sequence
//ff:what TestUnlistedPiece — what the new word's graph is told of a stem no table lists, but the place where the script changes: the stem as it was before a piece was fused into it (곰지 under 곰졌), what follows it (a predicate tail, particles, nothing), the kinds the tables know its letters as (사과 a nominal), whether its letters read as markers after a boundary (에 in 마당에서) and whether they hold, right after the atom at its head, only a title the spacing articles write apart (A씨 before 가); the place where the script changes is not told even where the new word crosses it (A씨, 제 before the atom 1-1 in 제1-1)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestUnlistedPiece(t *testing.T) {
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	noun := wordsScan("사과다", layerCost)
	if p := noun.unlistedPiece(0, 3, "사과", tail{start: 2, pieces: da}); p.Form != "사과" || p.Tail != rule.TailPred || p.Stem != kindNominal || p.Marked || p.Apart || p.Change {
		t.Fatalf("사과 before 다: %+v", p)
	}
	past := []piece{{text: "(었)", names: "선어말:ㅆ", verdict: 1, tail: rule.TailPred}, da[0]}
	if p := wordsScan("곰졌다", layerCost).unlistedPiece(0, 3, "곰지", tail{start: 2, base: '지', pieces: past}); p.Form != "곰지" || p.Tail != rule.TailPred || p.Stem != 0 {
		t.Fatalf("곰지 with 었 fused into its last syllable: %+v", p)
	}

	marked := wordsScan("마당에서", layerCost)
	marked.search()
	if p := marked.unlistedPiece(2, 3, "에", tail{start: 3}); !p.Marked || p.Tail != rule.TailBare || p.Apart || p.Change {
		t.Fatalf("에 after a boundary, read as a marker: %+v", p)
	}
	if p := marked.unlistedPiece(0, 3, "마당에", tail{start: 3}); p.Marked {
		t.Fatalf("the unit's first word is not marked: %+v", p)
	}

	ga := []piece{{text: "가", names: "주격:가", verdict: 1, tail: rule.TailCase}}
	title := wordsScan("A씨가", layerCost)
	title.ix.lex["씨"] = word{kinds: kindNominal, count: 500, by: [6]int{500}}
	if p := title.unlistedPiece(0, 3, "A씨", tail{start: 2, pieces: ga}); !p.Apart || p.Tail != rule.TailCase || p.Change || !title.crossesChange(0, 3, tail{start: 2, pieces: ga}) {
		t.Fatalf("A씨 before 가: written apart, and the place it crosses is not told: %+v", p)
	}
	order := wordsScan("제1-1", layerCost)
	if p := order.unlistedPiece(0, 4, "제1-1", tail{start: 4}); p.Form != "제1-1" || p.Tail != rule.TailBare || p.Marked || p.Apart || p.Change || !order.crossesChange(0, 4, tail{start: 4}) {
		t.Fatalf("제1-1: the closed-set 제 before the atom 1-1, a place that is not told: %+v", p)
	}
}
