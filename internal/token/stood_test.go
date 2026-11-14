//ff:func feature=token type=reader control=sequence
//ff:what TestStood — whether letters may stand as a word of no markers because a cue made them one; false for letters no cue made one, and for a scan whose cues were never found (stands shorter than the place)
package token

import "testing"

func TestStood(t *testing.T) {
	sc := spacingScan("뚜루뚜겸총무", "45")
	if !sc.stood(3, 4) || sc.stood(0, 4) || sc.stood(3, 5) {
		t.Fatal(sc.stands)
	}
	bare := &scan{ix: sc.ix, s: []rune("뚜루뚜겸")}
	if bare.stood(3, 4) {
		t.Fatal("no cues found")
	}
}
