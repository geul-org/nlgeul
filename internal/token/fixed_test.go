//ff:func feature=token type=reader control=sequence
//ff:what TestFixed — 이게 is 이것 and the subject particle under it; before 요 it stands as its particle does, before another case particle it does not; other letters give nothing; 건 is 거 and 는 only where a word that ends in a modifying closing ends right before it; 게 alone as a unit read as attaching is 거 and 이 after the modifying form of the unit before, but not at another place
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestFixed(t *testing.T) {
	sc := scanOf("이게")
	got := sc.fixed(2, -1)
	if len(got) != 1 || got[0].start != 2 || got[0].only != "이게" || got[0].shown != "이것" || got[0].label != "줄임:이게" || got[0].pieces[0].text != "(이)" || got[0].pieces[0].names != "주격:이" || got[0].pieces[0].tail != rule.TailCase {
		t.Fatal(got)
	}
	if got = sc.fixed(2, rule.LinkCase); len(got) != 0 {
		t.Fatal("a case particle after a case particle", got)
	}
	if got = scanOf("땐").fixed(1, -1); len(got) != 1 || got[0].shown != "때" {
		t.Fatal(got)
	}
	if got = scanOf("곰이").fixed(2, -1); len(got) != 0 {
		t.Fatal(got)
	}
	bound := scanOf("한건")
	if got = bound.fixed(2, -1); len(got) != 0 {
		t.Fatal("거 with a particle where no modifying closing ends", got)
	}
	bound.adnEnds[1] = true
	if got = bound.fixed(2, -1); len(got) != 1 || got[0].shown != "거" || !got[0].adn || got[0].pieces[0].text != "(는)" {
		t.Fatal(got)
	}
	ge := scanOf("게")
	if got = ge.fixed(1, -1); len(got) != 0 {
		t.Fatal("게 alone with nothing before it", got)
	}
	ge.attach = true
	if got = ge.fixed(1, -1); len(got) != 1 || got[0].shown != "거" || !got[0].adn || got[0].only != "게" || got[0].pieces[0].text != "(이)" {
		t.Fatal("게 as a unit attaching to the modifying form before", got)
	}
	inner := scanOf("곰게")
	inner.attach = true
	if got = inner.fixed(2, -1); len(got) != 0 {
		t.Fatal("게 not at the unit's start", got)
	}
}
