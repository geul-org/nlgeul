//ff:func feature=token type=rule control=sequence
//ff:what TestAfterAdn — right after a word that ends in a modifying closing, a verb's stem (마시, and 마시 as it was before 었 was fused into it, 듣 that 들 stands for) or an adjective's (작) of the word table does not stand before a predicate tail; let be are 듯하 (an auxiliary written onto a modifying form), a noun before a particle, a word no table has, a word that is itself a modifying or a naming form, a tail that names its own stem, a place where no such word ends, and the layer below the kinds' graphs; end to end in one unit 먹+을 is not followed by 마시+다 or 작+다, while 먹+을 사과+를, 먹+을 듯하+다, 먹+을 followed by 마시+다 in another unit, and 먹+을 as an object before 마시+다 stand
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAfterAdn(t *testing.T) {
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	past := []piece{{text: "었", names: "선어말:었", verdict: 1, tail: rule.TailPred}, da[0]}
	sc := wordsScan("먹을마시다", layerCost)
	if sc.afterAdn(2, tail{start: 4, pieces: da}, false, 0) {
		t.Fatal("no reading has a word that ends in a modifying closing end at 2")
	}
	sc.adnEnds[2] = true
	if !sc.afterAdn(2, tail{start: 4, pieces: da}, false, 0) {
		t.Fatal("마시, a verb of the table, before 다 right after 먹+을")
	}
	if sc.afterAdn(2, tail{start: 4, as: numName, pieces: da}, false, 0) {
		t.Fatal("a tail that names its own stem")
	}
	if sc.afterAdn(2, tail{start: 4, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1, tail: rule.TailPred}}}, false, 0) {
		t.Fatal("마시+는 is itself a modifying form")
	}
	if sc.afterAdn(2, tail{start: 4, pieces: []piece{{text: "기", names: "이름:기", verdict: 1, tail: rule.TailPred}, {text: "는", names: "보조사:는", verdict: 1, tail: rule.TailAux}}}, false, 0) {
		t.Fatal("마시+기+는 is itself a naming form")
	}
	low := wordsScan("먹을마시다", layerKnown)
	low.adnEnds[2] = true
	if low.afterAdn(2, tail{start: 4, pieces: da}, false, 0) {
		t.Fatal("below the layer in which the kinds' graphs judge a stem")
	}
	fused := wordsScan("먹을마셨다", layerCost)
	fused.adnEnds[2] = true
	if !fused.afterAdn(2, tail{start: 4, base: '시', pieces: past}, false, 0) {
		t.Fatal("마시 as it was before 었 was fused into 셨")
	}
	if fused.afterAdn(2, tail{start: 4, pieces: da}, false, 0) {
		t.Fatal("마셨 as written is no word of the table")
	}
	irr := wordsScan("먹을들었다", layerCost)
	irr.adnEnds[2] = true
	if !irr.afterAdn(2, tail{start: 3, shown: "듣", pieces: past}, false, 0) {
		t.Fatal("듣, the verb that 들 stands for before 었")
	}
	if irr.afterAdn(2, tail{start: 3, pieces: past}, false, 0) {
		t.Fatal("들 as written is no word of the table")
	}
	adj := wordsScan("먹을작다", layerCost)
	adj.adnEnds[2] = true
	if !adj.afterAdn(2, tail{start: 3, pieces: da}, false, 0) {
		t.Fatal("작, an adjective (and a noun) of the table, before 다")
	}
	aux := wordsScan("먹을듯하다", layerCost)
	aux.adnEnds[2] = true
	if aux.afterAdn(2, tail{start: 4, pieces: da}, false, 0) {
		t.Fatal("듯하, an auxiliary written onto a modifying form")
	}
	noun := wordsScan("먹을사과를", layerCost)
	noun.adnEnds[2] = true
	if noun.afterAdn(2, tail{start: 4, pieces: []piece{{text: "를", names: "목적격:를", verdict: 1, tail: rule.TailCase}}}, false, 0) {
		t.Fatal("사과, a noun, before a particle")
	}
	none := wordsScan("먹을곰다", layerCost)
	none.adnEnds[2] = true
	if none.afterAdn(2, tail{start: 3, pieces: da}, false, 0) {
		t.Fatal("곰, a word no table has")
	}

	r, _ := New(Options{Words: "testdata/words.tsv"})
	got := strings.Join(r.Lines("먹을마시다 먹는작다", false), "\n")
	if strings.Contains(got, "[꾸밈:을] 마시[열쇠:동사]") || strings.Contains(got, "[꾸밈:는] 작[열쇠:형용사]") || !strings.Contains(got, "  먹+을[목적격:을] 마시[열쇠:동사]+다[맺음:다]\t") {
		t.Fatal("a verb or an adjective right after a modifying form in one unit\n" + got)
	}
	if got := strings.Join(r.Lines("먹을사과를 먹을듯하다", false), "\n"); !strings.Contains(got, "  먹[열쇠:동사]+을[꾸밈:을] 사과[열쇠:체언]+를[목적격:를]\t") || !strings.Contains(got, "  먹[열쇠:동사]+을[꾸밈:을] 듯하[열쇠:형용사]+다[맺음:다]\t") {
		t.Fatal("a noun and 듯하 right after a modifying form\n" + got)
	}
	if got := strings.Join(r.Lines("먹을 마시다", false), "\n"); !strings.Contains(got, "  먹[열쇠:동사]+을[꾸밈:을]\t") || !strings.Contains(got, "  마시[열쇠:동사]+다[맺음:다]\t") {
		t.Fatal("a verb in the unit after a modifying form\n" + got)
	}
}
