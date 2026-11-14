//ff:func feature=token type=reader control=sequence
//ff:what TestWordOfDet — a word of no markers named only as a determiner is noted as one (어떤, 이 alone); a noun, a determiner's letters before a particle (이+가, named as the table's nominal) and a word ending in a modifying closing are not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWordOfDet(t *testing.T) {
	sc := wordsScan("어떤이가사과먹는", layerCost)
	if e := sc.wordOf(0, 2, tail{start: 2}); !e.det || e.adn || e.sig != "어떤[관형사:어떤]" {
		t.Fatal("어떤", e)
	}
	if e := sc.wordOf(2, 3, tail{start: 3}); !e.det || e.sig != "이[관형사:이]" {
		t.Fatal("이 alone", e)
	}
	ga := []piece{{text: "가", names: "주격:가", verdict: 1, tail: rule.TailCase}}
	if e := sc.wordOf(2, 4, tail{start: 3, pieces: ga}); e.det || e.sig != "이[열쇠:체언]+가[주격:가]" {
		t.Fatal("이 before a particle", e)
	}
	if e := sc.wordOf(4, 6, tail{start: 6}); e.det {
		t.Fatal("사과", e)
	}
	neun := []piece{{text: "는", names: "꾸밈:는", verdict: 1, tail: rule.TailPred}}
	if e := sc.wordOf(6, 8, tail{start: 7, pieces: neun}); e.det || !e.adn {
		t.Fatal("먹+는", e)
	}
}
