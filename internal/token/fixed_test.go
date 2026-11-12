//ff:func feature=token type=reader control=sequence
//ff:what TestFixed — 이게 is 이것 and the subject particle under it; before 요 it stands as its particle does, before another case particle it does not; other letters give nothing; 내게 is 내 and the 게 of 에게, 네게 so before 도 (with a word table too) but not before a case particle, 어디다 is 어디 and the 다 of 에다, 거기다가 거기 and 다가 of 에다가 or 거기다 before the emphatic 가, and 다 after another nominal is no such form; 건 is 거 and 는 only where a word that ends in a modifying closing ends right before it; 게 alone as a unit read as attaching is 거 and 이 after the modifying form of the unit before, but not at another place
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestFixed(t *testing.T) {
	sc := scanOf("이게")
	got := sc.fixed(2, -1)
	if len(got) != 1 || got[0].start != 2 || got[0].only != "이게" || got[0].shown != "이것" || got[0].label != "줄임:이게" || got[0].pieces[0].text != "(이)" || got[0].pieces[0].names != "주격:이" || got[0].pieces[0].tail != rule.TailCase || got[0].again != nil {
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
	if got = scanOf("내게").fixed(2, -1); len(got) != 1 || got[0].start != 2 || got[0].only != "내게" || got[0].shown != "내" || got[0].label != "줄임:내게" || got[0].pieces[0].text != "게" || got[0].pieces[0].names != "부사격:에게" {
		t.Fatal("내게 is 내 and 에게's 게", got)
	}
	lexed := wordsScan("네게도", layerCost)
	if got = lexed.fixed(2, rule.LinkTopic); len(got) != 1 || got[0].shown != "네" || got[0].pieces[0].names != "부사격:에게" {
		t.Fatal("네게 before 도, with a word table", got)
	}
	if got = lexed.fixed(2, rule.LinkCase); len(got) != 0 {
		t.Fatal("네게 before a case particle", got)
	}
	if got = scanOf("어디다").fixed(3, -1); len(got) != 1 || got[0].shown != "어디" || got[0].label != "줄임:어디다" || got[0].pieces[0].text != "다" || got[0].pieces[0].names != "부사격:에다" {
		t.Fatal("어디다 is 어디 and 에다's 다", got)
	}
	there := scanOf("거기다가")
	if got = there.fixed(4, -1); len(got) != 1 || got[0].shown != "거기" || got[0].only != "거기다가" || got[0].pieces[0].text != "다가" || got[0].pieces[0].names != "부사격:에다가" {
		t.Fatal("거기다가 is 거기 and 에다가's 다가", got)
	}
	if got = there.fixed(3, rule.LinkEmph); len(got) != 1 || got[0].shown != "거기" || got[0].pieces[0].names != "부사격:에다" {
		t.Fatal("거기다 before the emphatic 가", got)
	}
	if got = scanOf("학교다").fixed(3, -1); len(got) != 0 {
		t.Fatal("다 after another nominal", got)
	}
	bound := scanOf("한건")
	if got = bound.fixed(2, -1); len(got) != 0 {
		t.Fatal("거 with a particle where no modifying closing ends", got)
	}
	bound.adnEnds[1] = true
	if got = bound.fixed(2, -1); len(got) != 1 || got[0].shown != "거" || got[0].again == nil || got[0].facts == nil || !got[0].facts.Adn || got[0].pieces[0].text != "(는)" {
		t.Fatal(got)
	}
	ge := scanOf("게")
	if got = ge.fixed(1, -1); len(got) != 0 {
		t.Fatal("게 alone with nothing before it", got)
	}
	ge.attach = true
	if got = ge.fixed(1, -1); len(got) != 1 || got[0].shown != "거" || got[0].again == nil || got[0].facts == nil || !got[0].facts.Adn || got[0].only != "게" || got[0].pieces[0].text != "(이)" {
		t.Fatal("게 as a unit attaching to the modifying form before", got)
	}
	inner := scanOf("곰게")
	inner.attach = true
	if got = inner.fixed(2, -1); len(got) != 0 {
		t.Fatal("게 not at the unit's start", got)
	}
}
