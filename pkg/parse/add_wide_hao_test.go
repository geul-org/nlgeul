//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddWideHao — addWide puts the 하오체 closings 소, 오, 으오 and 올시다 in the index, each asking for a predicate's stem of the table, after the letters its form asks for (소 after a final consonant, 오 after a vowel, 으오 after a hard final consonant, 올시다 after any), as a final closing that takes no 요
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddWideHao(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	ix.addWide()
	for _, c := range []struct {
		name  string
		after int
	}{{"맺음:소", rule.AfterBatchim}, {"맺음:오", rule.AfterVowel}, {"맺음:으오", rule.AfterHard}, {"맺음:올시다", 0}} {
		found := entryNamed(ix, c.name)
		if found == nil {
			t.Fatal("not in the index", c.name)
		}
		if found.gate != ix.pred || found.after != c.after || !found.noYo || found.class != classClosing {
			t.Fatal(c.name, found)
		}
	}
}
