//ff:func feature=token type=reader control=sequence
//ff:what TestGroupAtWide — with a word table, the shortened particle 선 after a nominal of the table is read as the unit's end is looked at (the word after is asked for) and also when the next piece is asked for
package token

import "testing"

func TestGroupAtWide(t *testing.T) {
	sc := wordsScan("마당선", layerCost)
	now, _, _ := sc.groupAt(classParticle, 3, -1, false, false, false)
	more, _, _ := sc.groupAt(classParticle, 3, -1, true, false, false)
	if len(now) != 1 || len(more) != 1 {
		t.Fatal(now, more)
	}
}
