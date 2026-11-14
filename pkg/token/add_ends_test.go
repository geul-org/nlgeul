//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddEnds — what addEnds works out is, for every Hangul syllable and every key of the index, what entriesOf gives; a token no entry ends with is not kept; adding an entry drops it, and a scan then reads the entries from the index as before
package token

import (
	"slices"
	"testing"
)

func TestAddEnds(t *testing.T) {
	ix, _ := newIndex(Options{})
	if ix.ends == nil {
		t.Fatal("a new index has them worked out")
	}
	ix.ends = nil
	ix.addEnds()
	for r := rune(0xAC00); r <= 0xD7A3; r++ {
		if !slices.Equal(ix.ends[r], ix.entriesOf(r)) {
			t.Fatal(string(r))
		}
	}
	for k := range ix.byKey {
		if !slices.Equal(ix.ends[k], ix.entriesOf(k)) {
			t.Fatal(string(k))
		}
	}
	if _, kept := ix.ends['a']; kept {
		t.Fatal("no entry ends with it")
	}
	sc := scanOf("간")
	sc.ix = ix
	before := slices.Clone(sc.entriesAt(1))
	ix.add(&entry{name: "보조사:간", form: "간", class: classParticle})
	if ix.ends != nil {
		t.Fatal("add drops them")
	}
	if after := sc.entriesAt(1); len(after) != len(before)+1 || after[len(ix.byKey['간'])-1].name != "보조사:간" {
		t.Fatal(len(before), len(after))
	}
}
