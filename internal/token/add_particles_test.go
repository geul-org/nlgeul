//ff:func feature=token type=reader control=sequence
//ff:what TestAddParticles — a graph for each case reading: 이 has two (주격, 보격), 은 one
package token

import "testing"

func TestAddParticles(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}}
	ix.addParticles()
	if len(ix.byKey['이']) < 2 || ix.byKey['이'][0].name != "주격:이" || ix.byKey['이'][1].name != "보격:이" || ix.byKey['은'][0].name != "보조사:은" {
		t.Fatal(ix.byKey['이'][0].name)
	}
}
