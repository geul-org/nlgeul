//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkSides — applies when either side is not of the kind asked for
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkSides(t *testing.T) {
	alnum := toulmin.Specs{&Bits{Of: "sides", Mask: SideAlnum}}
	both, _ := MarkSides(markCtx(Mark{Left: SideAlnum, Right: SideAlnum | SideDigit}), alnum)
	left, _ := MarkSides(markCtx(Mark{Left: SideHangul, Right: SideAlnum}), alnum)
	right, _ := MarkSides(markCtx(Mark{Left: SideAlnum}), alnum)
	if both != -1 || left != 1 || right != 1 {
		t.Fatal(both, left, right)
	}
}
