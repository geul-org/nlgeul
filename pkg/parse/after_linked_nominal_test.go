//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAfterLinkedNominal — a copula tail whose dropped stem asks for a nominal before it (the gate of droppableKnown, of a nominal or of a nominal or a new word) is read after the linking 어서 (먹+어서+(이)+라는) without that gate: the word before the stem is the closing there; another gate the tail has is kept
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAfterLinkedNominal(t *testing.T) {
	sc := scanOf("먹어서라는")
	nominal, other := &entry{name: "앞말:체언"}, &entry{name: "앞말:용언"}
	fresh := &entry{name: "앞말:체언·새 낱말"}
	sc.ix.nominal, sc.ix.nominalNew = nominal, fresh
	pieces := []piece{{text: "(이)", names: droppedName, tail: rule.TailCopula, unseen: true}, {text: "라는", names: "꾸밈:라는", tail: rule.TailPred}}
	for _, c := range []struct {
		gate, want *entry
	}{
		{nominal, nil}, {fresh, nil}, {other, other},
	} {
		got := sc.afterLinked(tail{start: 3, pieces: pieces, gate: c.gate})
		i := slices.IndexFunc(got, func(l tail) bool { return l.pieces[0].names == "맺음:어서" })
		if i < 0 || got[i].gate != c.want {
			t.Errorf("%s: %+v", c.gate.name, got)
		}
	}
}
