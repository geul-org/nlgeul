//ff:func feature=token type=reader control=sequence
//ff:what TestFarBegins — a place where a word that is not a new word may begin farther than a new word's stem may be: the head of the unit (a span of code or a sign of a pair, longer than a stem) is one, a key of the tables with its last letter as it stands or taken off is one, and a place before the head that is no key is not; the units of the tests are read by the net too
package token

import (
	"slices"
	"testing"
)

func TestFarBegins(t *testing.T) {
	driveAll(t)
	sc := scanOf("가나다라마바사아자차카타")
	sc.head = 12
	sc.beginnings()
	if got := sc.begins[12]; len(got) == 0 || got[0] != 0 || slices.Contains(got, 1) {
		t.Fatal("the head of the unit", got)
	}
	if got := sc.farBegins(11, 1); slices.Contains(got, 0) {
		t.Fatal("not at the head", got)
	}
	r := streamReader(t, "words-far.tsv")
	far := &scan{ix: r.ix, s: []rune("고양이가대한민국임시정부수립기념일을")}
	far.prepare()
	if got := far.farBegins(17, 5); !slices.Contains(got, 4) {
		t.Fatal("a long known word", got)
	}
}
