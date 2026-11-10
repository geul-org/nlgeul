//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkEnclosed — applies when a run stands on both sides
package rule

import "testing"

func TestMarkEnclosed(t *testing.T) {
	yes, _ := MarkEnclosed(markCtx(Mark{Left: SideAlnum, Right: SideAlnum}), nil)
	no, _ := MarkEnclosed(markCtx(Mark{Left: SideHangul, Right: SideAlnum}), nil)
	end, _ := MarkEnclosed(markCtx(Mark{Left: SideAlnum}), nil)
	if yes != 1 || no != -1 || end != -1 {
		t.Fatal(yes, no, end)
	}
}
