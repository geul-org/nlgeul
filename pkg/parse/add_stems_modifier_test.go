//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddStemsModifier — on the new word's graph, with a word table and without, a new word whose letters are a modifier and a known word it modifies written together (Modifier) stands with a lower verdict, 2/(1+forceNew+forceModifier)-1, with markers after it or not: the fact is a rebuttal and no exception, so no reason of an exception is given; without the fact it stands as a new word does, 2/(1+forceNew)-1
package parse

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
		if math.Abs(v-(2/(1+forceNew+forceModifier)-1)) > 1e-9 || v <= 0 || slices.Contains(why, whyModifier) {
			t.Errorf("lexed %v tail %d: %v %v", c.ix.lexed, c.tail, v, why)
		}
		if free, _ := verdictOf(c.ix.stem, rule.Piece{Form: "못와", Tail: c.tail}, "", false); math.Abs(free-(2/(1+forceNew)-1)) > 1e-9 {
			t.Errorf("lexed %v tail %d without the fact: %v", c.ix.lexed, c.tail, free)
		}
	}
}
