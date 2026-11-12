//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAfterAdnDet — right after a determiner a verb's stem of the word table (마시) does not stand before a predicate tail, nor does 듯하, which is written onto a modifying form only; a word that is itself a modifying form, a noun before a particle, and a place where no determiner ends are let be, and the same place asked for a modifying closing is not barred; a word that only stands alone (아이고) does not stand right after a modifying closing or a determiner, while an adverb (빨리) does; end to end 어떤 is not followed in its unit by 마시+다 or 아이고 and 먹+을 not by 아이고, while 어떤 마시+는, 어떤 사과+를, 사과 아이고, 먹+을 빨리 and 마시+다 in the unit after 어떤 stand
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAfterAdnDet(t *testing.T) {
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	sc := wordsScan("어떤마시다", layerCost)
	if sc.afterAdn(2, tail{start: 4, pieces: da}, true, 0) {
		t.Fatal("no reading has a determiner end at 2")
	}
	sc.detEnds[2] = true
	if !sc.afterAdn(2, tail{start: 4, pieces: da}, true, 0) || sc.afterAdn(2, tail{start: 4, pieces: da}, false, 0) {
		t.Fatal("마시, a verb of the table, before 다 right after 어떤")
	}
	if sc.afterAdn(2, tail{start: 4, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1, tail: rule.TailPred}}}, true, 0) {
		t.Fatal("마시+는 is itself a modifying form")
	}
	aux := wordsScan("어떤듯하다", layerCost)
	aux.detEnds[2] = true
	if !aux.afterAdn(2, tail{start: 4, pieces: da}, true, 0) {
		t.Fatal("듯하 is written onto a modifying form, not onto a determiner")
	}
	noun := wordsScan("어떤사과를", layerCost)
	noun.detEnds[2] = true
	if noun.afterAdn(2, tail{start: 4, pieces: []piece{{text: "를", names: "목적격:를", verdict: 1, tail: rule.TailCase}}}, true, 0) {
		t.Fatal("사과, a noun, before a particle")
	}
	alone := wordsScan("먹을아이고", layerCost)
	alone.adnEnds[2], alone.detEnds[2] = true, true
	if !alone.afterAdn(2, tail{start: 5}, false, 0) || !alone.afterAdn(2, tail{start: 5}, true, 0) {
		t.Fatal("아이고, a word that only stands alone, right after what modifies a nominal")
	}
	adverb := wordsScan("먹을빨리", layerCost)
	adverb.adnEnds[2], adverb.detEnds[2] = true, true
	if adverb.afterAdn(2, tail{start: 4}, false, 0) || adverb.afterAdn(2, tail{start: 4}, true, 0) {
		t.Fatal("빨리, an adverb, is let be")
	}

	r, _ := New(Options{Words: "testdata/words.tsv"})
	got := strings.Join(r.Lines("어떤마시다 어떤아이고 먹을아이고", false), "\n")
	if strings.Contains(got, "] 마시[열쇠:동사]+다") || strings.Contains(got, "어떤[관형사:어떤] 아이고[열쇠:홀로]") || strings.Contains(got, "[꾸밈:을] 아이고[열쇠:홀로]") || !strings.Contains(got, "  먹+을[목적격:을] 아이고[열쇠:홀로]\t") {
		t.Fatal("a predicate or an interjection right after what modifies a nominal\n" + got)
	}
	got = strings.Join(r.Lines("어떤마시는 어떤사과를 사과아이고 먹을빨리", false), "\n")
	for _, want := range []string{"  어떤[관형사:어떤] 마시[열쇠:동사]+는[꾸밈:는]\t", "  어떤[관형사:어떤] 사과[열쇠:체언]+를[목적격:를]\t", "  사과[열쇠:체언] 아이고[열쇠:홀로]\t", "  먹[열쇠:동사]+을[꾸밈:을] 빨리[열쇠:부사]\t"} {
		if !strings.Contains(got, want) {
			t.Errorf("no %s\n%s", want, got)
		}
	}
	if got := strings.Join(r.Lines("어떤 마시다", false), "\n"); !strings.Contains(got, "어떤[관형사:어떤]") || !strings.Contains(got, "  마시[열쇠:동사]+다[맺음:다]\t") {
		t.Fatal("a verb in the unit after a determiner\n" + got)
	}
}
