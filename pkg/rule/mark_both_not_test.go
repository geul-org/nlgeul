//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkBothNot — with Not it applies when either side is not of the kind (a dot not between two runs of letters and digits), and not when both are
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkBothNot(t *testing.T) {
	alnum := toulmin.Specs{&Bits{Of: "sides", Mask: SideAlnum}, &Not{}}
	both, bv := MarkBoth(markCtx(Mark{Left: SideAlnum, Right: SideAlnum | SideDigit}), alnum)
	left, _ := MarkBoth(markCtx(Mark{Left: SideHangul, Right: SideAlnum}), alnum)
	right, _ := MarkBoth(markCtx(Mark{Left: SideAlnum}), alnum)
	neither, _ := MarkBoth(markCtx(Mark{}), alnum)
	if both != -1 || left != 1 || right != 1 || neither != 1 || bv != nil {
		t.Fatal(both, left, right, neither, bv)
	}
}
