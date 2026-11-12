//ff:func feature=token type=reader control=sequence
//ff:what TestAtomStemWide — a whole run of digits and a unit that was cut in two has the run's name and verdict; its number alone is still a number; a part of the run has nothing
package token

import "testing"

func TestAtomStemWide(t *testing.T) {
	sc := scanOf("64GB와 7m")
	if name, v := sc.atomStem(0, 4); name != runName || v != sc.ix.run {
		t.Fatal(name, v)
	}
	if name, v := sc.atomStem(0, 2); name != numName || v != sc.ix.number {
		t.Fatal(name, v)
	}
	if _, v := sc.atomStem(0, 3); v != 0 {
		t.Fatal("part of a run")
	}
	if name, v := sc.atomStem(6, 8); name != runName || v != sc.ix.run {
		t.Fatal(name, v)
	}
}
