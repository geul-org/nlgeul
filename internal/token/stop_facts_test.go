//ff:func feature=token type=reader control=sequence
//ff:what TestStopFacts — the sides of the dots and their number; nothing stands at the unit's edges
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestStopFacts(t *testing.T) {
	sc := scanOf("했다.그리고")
	if m := sc.stopFacts(2, 3); m.Left != rule.SideHangul || m.Right != rule.SideHangul || m.Run != 1 {
		t.Fatalf("%+v", m)
	}
	if m := scanOf("...").stopFacts(0, 3); m.Left != 0 || m.Right != 0 || m.Run != 3 {
		t.Fatalf("%+v", m)
	}
}
