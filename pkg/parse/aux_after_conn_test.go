//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAuxAfterConn — an auxiliary verb's stem (주 for 줘) with a predicate tail where some candidate has a word that ends in the linking 어/아 end right before it is one written onto its main verb; not at the unit's head, not where no such word ends, not before particles (주+는 the noun), not for a stem that is no auxiliary's (조), and 가, 내 and 드리 only with a word table
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxAfterConn(t *testing.T) {
	pred := tail{start: 2, base: '주', pieces: []piece{{text: "(어)", names: "맺음:어(합침)", verdict: 1, tail: rule.TailPred}}}
	particle := tail{start: 3, pieces: []piece{{text: "는", names: "보조사:는", verdict: 1, tail: rule.TailAux}}}
	sc := scanOf("해줘")
	if sc.auxAfterConn(1, "주", pred) {
		t.Fatal("no word that ends in 어/아 ends before it")
	}
	sc.connEnds[0], sc.connEnds[1] = true, true
	if !sc.auxAfterConn(1, "주", pred) {
		t.Fatal("주 with a predicate tail right after a word that ends in 어/아")
	}
	if sc.auxAfterConn(0, "주", pred) {
		t.Fatal("at the unit's head")
	}
	if sc.auxAfterConn(1, "주", particle) {
		t.Fatal("before a particle")
	}
	if sc.auxAfterConn(1, "조", pred) {
		t.Fatal("a stem that is no auxiliary's")
	}
	for _, stem := range []string{"보", "버리", "놓", "두", "있"} {
		if !sc.auxAfterConn(1, stem, pred) {
			t.Error("an auxiliary of the closed set", stem)
		}
	}
	lexed := wordsScan("해줘", layerCost)
	lexed.connEnds[1] = true
	for _, stem := range []string{"가", "내", "드리"} {
		if sc.auxAfterConn(1, stem, pred) || !lexed.auxAfterConn(1, stem, pred) {
			t.Error("an auxiliary read so only with a word table", stem)
		}
	}
}
