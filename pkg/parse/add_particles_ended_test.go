//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAddParticlesEnded — the particles read only right after a closing are in the index and marked so: the 야 after a consonant, the 가 of emphasis after 지, and 그려; the 야 after a vowel, the subject particle 가 and the emphatic 가 after a particle are not
package parse

import (
	"slices"
	"testing"
)

func TestAddParticlesEnded(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	ended := map[string]int{}
	for _, e := range slices.Concat(ix.byKey['야'], ix.byKey['가'], ix.byKey['려']) {
		if e.ended {
			ended[e.name]++
		}
	}
	if len(ended) != 3 || ended["보조사:야"] != 1 || ended["보조사:가"] != 1 || ended["보조사:그려"] != 1 {
		t.Fatal(ended)
	}
}
