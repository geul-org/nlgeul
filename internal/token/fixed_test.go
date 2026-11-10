//ff:func feature=token type=reader control=sequence
//ff:what TestFixed — 이게 is 이것 and the subject particle under it; before 요 it stands as its particle does, before another case particle it does not; other letters give nothing
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
}
