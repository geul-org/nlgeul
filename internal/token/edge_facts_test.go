//ff:func feature=token type=reader control=sequence
//ff:what TestEdgeFacts — the facts of a stretch of one sign at a unit's edge: how many stand in a row, what stands on each side in the unit, whole-unit (MarkAlone), after the read letters (MarkEnds), before them with the letters right after (MarkLeads), at an input line's start (MarkHeads), with a space right after (MarkSpaced)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEdgeFacts(t *testing.T) {
	sc := scanOf("태그")
	m := sc.edgeFacts(unit{text: "#태그", first: true}, 0, 1, false, true)
	if m.Run != 1 || m.Left != 0 || m.Right != rule.SideHangul || m.Facts != rule.MarkLeads|rule.MarkHeads {
		t.Fatal("before the letters", m)
	}
	m = sc.edgeFacts(unit{text: "#@태그"}, 0, 1, false, false)
	if m.Right != 0 || m.Facts != rule.MarkBySign {
		t.Fatal("the stretch is not the last before the letters", m)
	}
	m = scanOf("뭐").edgeFacts(unit{text: "뭐?!", spaced: true}, 1, 3, true, true)
	if m.Run != 2 || m.Left != rule.SideHangul || m.Right != 0 || m.Facts != rule.MarkEnds|rule.MarkSpaced {
		t.Fatal("after the letters", m)
	}
	m = scanOf("뭐").edgeFacts(unit{text: "뭐?!"}, 1, 2, true, false)
	if m.Run != 1 || m.Facts != rule.MarkEnds|rule.MarkBySign {
		t.Fatal("a stretch with another after it", m)
	}
	m = scanOf("").edgeFacts(unit{text: "...", first: true, spaced: true}, 0, 3, true, true)
	if m.Run != 3 || m.Left != 0 || m.Right != 0 || m.Facts != rule.MarkAlone|rule.MarkHeads|rule.MarkSpaced {
		t.Fatal("the whole unit", m)
	}
	m = scanOf("a").edgeFacts(unit{text: "a:1"}, 1, 2, true, false)
	if m.Left != rule.SideAlnum || m.Right != rule.SideAlnum|rule.SideDigit {
		t.Fatal("what stands on each side", m)
	}
}
