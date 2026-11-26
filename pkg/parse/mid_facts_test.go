//ff:func feature=parse type=reader control=sequence
//ff:what TestMidFacts — the facts of a sign between letters: the run, what stands on each side, and that a sign right next to it is told (MarkBySign), and that a sign that is no punctuation mark of a sentence stands in its run (MarkDrawn)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMidFacts(t *testing.T) {
	m := scanOf("밥·a1").midFacts(1, 2)
	if m.Run != 1 || m.Left != rule.SideHangul || m.Right != rule.SideAlnum || m.Facts != 0 {
		t.Fatal("between letters", m)
	}
	m = scanOf("밥··국").midFacts(1, 3)
	if m.Run != 2 {
		t.Fatal("a run", m)
	}
	if m = scanOf("밥·(국").midFacts(1, 2); m.Facts != rule.MarkBySign {
		t.Fatal("a sign right after", m)
	}
	if m = scanOf("밥)·국").midFacts(2, 3); m.Facts != rule.MarkBySign {
		t.Fatal("a sign right before", m)
	}
	if m = scanOf("밥-->국").midFacts(1, 3); m.Facts != rule.MarkBySign|rule.MarkDrawn {
		t.Fatal("a sign that draws right after", m)
	}
}
