//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAgain — 건 read as 거+(는) where a word that ends in a modifying closing ends is asked again on each reading: it stands on one whose last word ends in such a closing and not on another, and the tail's own facts are left as they were; a tail with no graph to ask again stands on any reading; end to end every reading of 한건 with 거+(는) has the modifying 하+(ㄴ) before it
package token

import (
	"strings"
	"testing"
)

func TestAgain(t *testing.T) {
	bound := scanOf("한건")
	bound.adnEnds[1] = true
	got := bound.fixed(2, -1)
	if len(got) != 1 || got[0].again == nil || got[0].facts == nil {
		t.Fatal(got)
	}
	geon := got[0]
	if !bound.again(geon, true) {
		t.Fatal("거+(는) on a reading whose last word ends in a modifying closing")
	}
	if bound.again(geon, false) || !geon.facts.Adn {
		t.Fatal("거+(는) on a reading whose last word does not; the tail's facts stay as they were")
	}
	if !bound.again(tail{start: 2, pieces: geon.pieces}, false) || !bound.again(tail{start: 2}, true) {
		t.Fatal("a tail with no graph to ask again stands on any reading")
	}

	r, _ := New(Options{})
	seen := 0
	for _, line := range r.Lines("한건", false) {
		if !strings.Contains(line, " 거[의존명사:거]+(는)[보조사:는]") {
			continue
		}
		seen++
		if !strings.Contains(line, "+(ㄴ)[꾸밈:ㄴ] 거[의존명사:거]") {
			t.Error("거+(는) after a word that does not end in a modifying closing: " + line)
		}
	}
	if seen == 0 {
		t.Fatal("한건 has no reading 하+(ㄴ) 거+(는)")
	}
}
