//ff:func feature=token type=reader control=sequence
//ff:what TestChange — between a run and a Hangul syllable the script changes, either way; not inside either, not next to a sign, not at the edges
package token

import "testing"

func TestChange(t *testing.T) {
	sc := scanOf("API호출a.가")
	if !sc.change(3) || !sc.change(5) || sc.change(1) || sc.change(4) || sc.change(6) || sc.change(7) || sc.change(0) || sc.change(8) {
		t.Fatal("change")
	}
}
