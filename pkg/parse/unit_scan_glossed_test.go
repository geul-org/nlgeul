//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestUnitScanGlossed — a unit is told glossed when its letters that are read (after its head) are Hangul syllables as many as the Han characters the bracket after it holds and no sign was trimmed off its end (대곤 before (大棍), "관치 with a sign of a pair at its head); a count that differs, no bracket, a sign trimmed off its end and a letter that is no Hangul syllable leave it untold, and closed is told as before
package parse

import "testing"

func TestUnitScanGlossed(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		u       unit
		glossed bool
		closed  bool
	}{
		{unit{text: "대곤", gloss: 2}, true, false},
		{unit{text: "\"관치", head: 1, name: pairName, gloss: 2}, true, false},
		{unit{text: "대곤", gloss: 1}, false, false},
		{unit{text: "대곤"}, false, false},
		{unit{text: "대곤,", gloss: 2}, false, false},
		{unit{text: "대곤.", gloss: 2}, false, true},
		{unit{text: "3곤", gloss: 2}, false, false},
		{unit{text: "\"관치", head: 1, name: pairName, gloss: 3}, false, false},
	} {
		sc := unitScan(ix, c.u)
		if sc.glossed != c.glossed || sc.closed != c.closed {
			t.Errorf("unitScan(%q, gloss %d): glossed %v closed %v, want %v %v", c.u.text, c.u.gloss, sc.glossed, sc.closed, c.glossed, c.closed)
		}
	}
}
