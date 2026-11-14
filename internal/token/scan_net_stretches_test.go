//ff:func feature=token type=reader control=sequence
//ff:what TestScanNetStretches — a unit read by the net gives its stretches in order, each with cells (아버지가방에들어가신다 is one stretch of 11 letters with 60 ways); a unit with no letters gives none
package token

import "testing"

func TestScanNetStretches(t *testing.T) {
	sc, u, _ := netUnit(t, "words-stream.tsv", "아버지가방에들어가신다")
	got := sc.netStretches(u)
	if len(got) != 1 || got[0].contest.Text != "아버지가방에들어가신다" || got[0].cells == nil || got[0].cells.ways.Int64() != 60 {
		t.Fatal(got)
	}
	empty, u2, _ := netUnit(t, "words-stream.tsv", "가")
	empty.s = nil
	if len(empty.netStretches(u2)) != 0 {
		t.Fatal("no letters")
	}
}
