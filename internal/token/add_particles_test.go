//ff:func feature=token type=reader control=sequence
//ff:what TestAddParticles — a graph for each case reading: 이 has two (주격, 보격), 은 one; the suffix 이 and the emphatic 야 that only this layer reads are there too; the readings of one particle have their own kinds
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddParticles(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	last := ix.byKey['이'][len(ix.byKey['이'])-1]
	if last.name != "접미:이" || !last.inner || last.slot != rule.LinkSuffix {
		t.Fatal(last)
	}
	ya := ix.byKey['야'][len(ix.byKey['야'])-1]
	if ya.name != "보조사:야" || ya.slot != rule.LinkTopic || ix.byKey['서'][0].slot == ix.byKey['서'][1].slot {
		t.Fatal(ya)
	}
	if len(ix.byKey['이']) < 2 || ix.byKey['이'][0].name != "주격:이" || ix.byKey['이'][1].name != "보격:이" || ix.byKey['은'][0].name != "보조사:은" {
		t.Fatal(ix.byKey['이'][0].name)
	}
}
