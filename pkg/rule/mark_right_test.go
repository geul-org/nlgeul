//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkRight — applies when what stands right of the sign is not of the kind asked for
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkRight(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "sides", Mask: SideDigit}}
	yes, _ := MarkRight(markCtx(Mark{Right: SideHangul}), specs)
	no, _ := MarkRight(markCtx(Mark{Right: SideAlnum | SideDigit}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
