//ff:func feature=token type=reader control=sequence
//ff:what TestAdd — an entry goes under its key with a graph, and the totals grow
package token

import "testing"

func TestAdd(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.add(&entry{name: "부사격:에서", form: "에서", class: classParticle, count: 5})
	if len(ix.byKey['서']) != 1 || ix.byKey['서'][0].g == nil || ix.graphs != 1 || ix.edges != 2 || ix.total != 5 {
		t.Fatalf("%+v", ix)
	}
}
