//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestClosingParticles — every name of closingParticles is a name of the index made with a word table: the first column a particle's, the second the names of closings
package parse

import (
	"strings"
	"testing"
)

func TestClosingParticles(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	class := map[string]int{}
	for _, es := range r.ix.byKey {
		for _, e := range es {
			class[e.name] = e.class + 1
		}
	}
	var closings []string
	for name, befores := range closingParticles {
		if class[name] != classParticle+1 {
			t.Errorf("%s is no particle of the index", name)
		}
		closings = append(closings, strings.Split(befores, "|")...)
	}
	for _, before := range closings {
		if class[before] != classClosing+1 {
			t.Errorf("%s is no closing of the index", before)
		}
	}
}
