//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddStemsModifier — on the new word's graph, with a word table and without, a new word whose letters are a modifier and a known word it modifies written together (Modifier) does not stand, with markers after it or not, for the reason on its own node; without the fact it stands as a new word does (1/3)
package token

import (
	"math"
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddStemsModifier(t *testing.T) {
	bare, lexed := &index{}, &index{lexed: true}
	bare.addStems()
	lexed.addStems()
	for _, c := range []struct {
		ix   *index
		tail int
	}{{bare, rule.TailBare}, {bare, rule.TailCase}, {lexed, rule.TailBare}, {lexed, rule.TailCase}} {
		v, why := verdictOf(c.ix.stem, rule.Piece{Form: "못와", Tail: c.tail, Modifier: true}, "", true)
		if v > 0 || !slices.Contains(why, whyModifier) {
			t.Errorf("lexed %v tail %d: %v %v", c.ix.lexed, c.tail, v, why)
		}
		if free, _ := verdictOf(c.ix.stem, rule.Piece{Form: "못와", Tail: c.tail}, "", false); math.Abs(free-(2/(1+forceNew) - 1)) > 1e-9 {
			t.Errorf("lexed %v tail %d without the fact: %v", c.ix.lexed, c.tail, free)
		}
	}
}
