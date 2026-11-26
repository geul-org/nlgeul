//ff:func feature=parse type=reader control=sequence
//ff:what TestAddParticles — a graph for each case candidate: 이 has two (주격, 보격), 은 one; the suffix 이 (read with a particle after it, and once more for the place of one who is called, with the graph the name must pass) and the emphatic 야 that only this layer reads are there too; the candidates of one particle have their own kinds
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddParticles(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	n := len(ix.byKey['이'])
	inner, last := ix.byKey['이'][n-2], ix.byKey['이'][n-1]
	if inner.name != nameSuffix || !inner.inner || inner.called || inner.gate != nil || inner.slot != rule.LinkSuffix {
		t.Fatal(inner)
	}
	if last.name != nameSuffix || last.inner || !last.called || last.gate == nil || last.gate.name != nameGate || last.slot != rule.LinkSuffix {
		t.Fatal(last)
	}
	ya := ix.byKey['야'][len(ix.byKey['야'])-1]
	if ya.name != "보조사:야" || ya.slot != rule.LinkTopic || ix.byKey['서'][1].slot == ix.byKey['서'][2].slot {
		t.Fatal(ya)
	}
	if len(ix.byKey['이']) < 2 || ix.byKey['이'][0].name != "주격:이" || ix.byKey['이'][1].name != "보격:이" || ix.byKey['은'][0].name != "보조사:은" {
		t.Fatal(ix.byKey['이'][0].name)
	}
}
