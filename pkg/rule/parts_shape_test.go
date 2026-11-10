//ff:func feature=rule type=rule control=sequence
//ff:what TestPartsShape — applies when the runs a sign joins are not of the shape, or when there are none
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPartsShape(t *testing.T) {
	specs := toulmin.Specs{&Shape{Of: "시각::", Min: 2, Max: 3, Lens: [][2]int{{1, 2}, {2, 2}}, Vals: [][2]int{{0, 24}, {0, 59}}}}
	fits, _ := PartsShape(markCtx(Mark{Parts: []string{"10", "30"}}), specs)
	not, _ := PartsShape(markCtx(Mark{Parts: []string{"25", "70"}}), specs)
	none, _ := PartsShape(markCtx(Mark{}), specs)
	if fits != -1 || not != 1 || none != 1 {
		t.Fatal(fits, not, none)
	}
}
