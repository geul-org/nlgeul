//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkBoth — applies when both sides are of the kind; not when one is not
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkBoth(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "sides", Mask: SideDigit}}
	yes, _ := MarkBoth(markCtx(Mark{Left: SideAlnum | SideDigit, Right: SideDigit}), specs)
	no, _ := MarkBoth(markCtx(Mark{Left: SideAlnum, Right: SideDigit}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
