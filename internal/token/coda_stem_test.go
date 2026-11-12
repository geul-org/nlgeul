//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCodaStem — a new word's stem that ends in ㅆ with markers after it is no stem (곰졌+다, 곰졌+도), judged as it was before a piece was fused into it (곰지+었 is let be); let be are a stem of another final consonant (ㅅ, not ㅆ), a stem some table knows (있 of a closed set, 맛있 once the table has it), a whole new word with no marker, a tail that names its own stem, and every stem without a word table; end to end 떨어졌다 loses 떨어졌+다 and 졌+다 that the reading with no table keeps, and 곰졌다, 곰졌도, 곰졌 each stand as a whole new word only
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCodaStem(t *testing.T) {
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	if scanOf("곰졌다").codaStem(0, tail{start: 2, pieces: da}) {
		t.Fatal("without a word table a new word's stem is not asked")
	}
	ws := wordsScan("곰졌다", layerCost)
	if !ws.codaStem(0, tail{start: 2, pieces: da}) {
		t.Fatal("곰졌, a new word that ends in ㅆ, before 다")
	}
	if ws.codaStem(0, tail{start: 3}) {
		t.Fatal("a whole new word with no marker")
	}
	if ws.codaStem(0, tail{start: 2, shown: "곰졌", pieces: da}) || ws.codaStem(0, tail{start: 2, as: numName, pieces: da}) {
		t.Fatal("a tail that names its own stem")
	}
	past := []piece{{text: "(었)", names: "선어말:ㅆ", verdict: 1, tail: rule.TailPred}, da[0]}
	if ws.codaStem(0, tail{start: 2, base: '지', pieces: past}) {
		t.Fatal("곰지 with 었 fused into its last syllable: the stem as it was has no ㅆ")
	}
	if !wordsScan("곰졌도", layerCost).codaStem(0, tail{start: 2, pieces: []piece{{text: "도", names: "보조사:도", verdict: 1, tail: rule.TailAux}}}) {
		t.Fatal("곰졌, a new word that ends in ㅆ, before a particle only")
	}
	if wordsScan("곰갓다", layerCost).codaStem(0, tail{start: 2, pieces: da}) {
		t.Fatal("곰갓 ends in ㅅ, not in ㅆ")
	}
	have := wordsScan("있다", layerCost)
	if v, _ := verdictOf(have.ix.stem, rule.Piece{Form: "있", Tail: rule.TailPred}, ""); v > 0 || have.codaStem(0, tail{start: 1, pieces: da}) {
		t.Fatal("있, of a closed set: the new word's graph turns it down, the closed set knows it", v)
	}
	tasty := wordsScan("맛있다", layerCost)
	if !tasty.codaStem(0, tail{start: 2, pieces: da}) {
		t.Fatal("맛있, which the table does not have, before 다")
	}
	tasty.ix.lex["맛있"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	if tasty.codaStem(0, tail{start: 2, pieces: da}) {
		t.Fatal("맛있, once the table has it as an adjective")
	}

	r, _ := New(Options{Words: "testdata/words.tsv"})
	got := strings.Join(r.Lines("떨어졌다", false), "\n")
	if strings.Contains(got, "떨어졌+") || strings.Contains(got, " 졌+") || !strings.Contains(got, "\n  떨어졌다\t") {
		t.Fatal("떨어졌다 with a word table\n" + got)
	}
	if bare := strings.Join(Lines("떨어졌다", false), "\n"); !strings.Contains(bare, "  떨어졌+다[맺음:다]\t") || !strings.Contains(bare, " 졌+다[맺음:다]\t") {
		t.Fatal("떨어졌다 without a word table\n" + bare)
	}
	bundles := r.Bundles("곰졌다 곰졌도 곰졌")
	if len(bundles) != 3 {
		t.Fatalf("%+v", bundles)
	}
	for _, b := range bundles {
		if len(b.Branches) != 1 || len(b.Branches[0].Words) != 1 || b.Branches[0].String() != b.Text || len(b.Branches[0].Words[0].Parts[0].Names) != 0 || !b.Branches[0].Words[0].Parts[0].Stem {
			t.Fatalf("%s: a whole new word only, %+v", b.Text, b.Branches)
		}
	}
}
