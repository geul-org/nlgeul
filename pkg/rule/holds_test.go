//ff:func feature=rule type=helper control=sequence
//ff:what TestHolds — without a Not the value follows the fact (+1 yes, -1 no); a Not anywhere among the criteria turns it round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestHolds(t *testing.T) {
	plain := toulmin.Specs{&Bits{Of: "piece", Mask: PieceStem}, &Claim{Name: "a", Concludes: "b"}}
	turned := toulmin.Specs{&Bits{Of: "piece", Mask: PieceStem}, &Not{}, &Claim{Name: "a", Concludes: "b"}}
	bare := toulmin.Specs{&Not{}}
	if Holds(true, nil) != 1 || Holds(false, nil) != -1 {
		t.Fatal("no criteria")
	}
	if Holds(true, plain) != 1 || Holds(false, plain) != -1 {
		t.Fatal("without Not")
	}
	if Holds(true, turned) != -1 || Holds(false, turned) != 1 || Holds(true, bare) != -1 || Holds(false, bare) != 1 {
		t.Fatal("with Not")
	}
}
