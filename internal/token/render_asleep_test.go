//ff:func feature=token type=reader control=sequence
//ff:what TestRenderAsleep — a head line with the unit and the number of readings asleep, then each with its verdict and preference; nothing when none is asleep
package token

import "testing"

func TestRenderAsleep(t *testing.T) {
	got := renderAsleep("곰이", []*path{{sig: "곰 이", verdict: 0.5, pref: -3}})
	if len(got) != 2 || got[0] != "곰이\t잠듦(1)" || got[1] != "  곰 이\t성립 0.50 선호 -3.00" {
		t.Fatal(got)
	}
	if renderAsleep("곰이", nil) != nil {
		t.Fatal("none asleep")
	}
}
