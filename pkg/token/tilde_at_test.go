//ff:func feature=token type=reader control=sequence
//ff:what TestTildeAt — the range sign after a number and its unit and before a number closes that tail; with no unit before it, with no number after it, or with no such sign nothing is read
package token

import "testing"

func TestTildeAt(t *testing.T) {
	got := scanOf("1월~3월").tildeAt(3)
	if len(got) != 1 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].pieces[1].text != "~" || got[0].pieces[1].names != "범위" {
		t.Fatal(got)
	}
	if len(scanOf("가~3월").tildeAt(2)) != 0 || len(scanOf("1월~가").tildeAt(3)) != 0 || len(scanOf("1월~3월").tildeAt(2)) != 0 {
		t.Fatal("read where it should not be")
	}
}
