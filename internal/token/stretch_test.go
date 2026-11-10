//ff:func feature=token type=reader control=sequence
//ff:what TestStretch — the words of a reading within a stretch, with their least verdict and summed preference
package token

import "testing"

func TestStretch(t *testing.T) {
	p := grow(grow(grow(&path{verdict: 1}, eojeol{end: 2, marked: "a", names: "x", verdict: 1, pref: -1}), eojeol{end: 3, marked: "b", names: "y", verdict: 0.5, pref: -2}), eojeol{end: 5, marked: "c", names: "z", verdict: 0.2, pref: -4})
	if got := stretch(p, 0, 3); got != (branch{"a b", "x + y", 0.5, -3}) {
		t.Fatal(got)
	}
}
