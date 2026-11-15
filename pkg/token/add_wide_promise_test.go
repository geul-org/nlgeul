//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddWidePromise — addWide puts the closing 마 of a promise and its 으 form 으마 in the index, each asking for a predicate's stem of the table, 마 after a vowel or ㄹ and 으마 after another final consonant, as a final closing that the closing set of verbs only holds and that bars the copula
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddWidePromise(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	ix.addWide()
	for _, c := range []struct {
		name  string
		after int
	}{{"맺음:마", rule.AfterOpen}, {"맺음:으마", rule.AfterHard}} {
		found := entryNamed(ix, c.name)
		if found == nil {
			t.Fatal("not in the index", c.name)
		}
		if found.gate != ix.pred || found.after != c.after || !found.bars || found.class != classClosing || !slices.Contains(verbClosings, c.name) {
			t.Fatal(c.name, found)
		}
	}
}
