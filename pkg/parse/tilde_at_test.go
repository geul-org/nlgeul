//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestTildeAt — the range sign (the tilde or a variant) after a number and its unit and before a number closes that tail; with no unit before it, with no number after it, or with no such sign nothing is read
package parse

import "testing"

func TestTildeAt(t *testing.T) {
	for _, sign := range []string{"~", "∼", "～", "〜"} {
		got := scanOf("1월" + sign + "3월").tildeAt(3)
		if len(got) != 1 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].pieces[1].text != sign || got[0].pieces[1].names != "범위" {
			t.Fatal(sign, got)
		}
	}
	if len(scanOf("가~3월").tildeAt(2)) != 0 || len(scanOf("1월~가").tildeAt(3)) != 0 || len(scanOf("1월~3월").tildeAt(2)) != 0 || len(scanOf("1월–3월").tildeAt(3)) != 0 {
		t.Fatal("read where it should not be")
	}
}
