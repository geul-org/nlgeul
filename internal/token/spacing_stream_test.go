//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingStream — the stream of 갑돌씨가 with article 48 on has the way [0-2] [2-4] among the ways Next and Again give; with it off it has not, and has fewer ways; no word is made twice with different values (clashes is 0)
package token

import "testing"

func TestSpacingStream(t *testing.T) {
	ways := func(spacing ...string) (map[string]int, *Stream, Segment) {
		st := spacingReader(t, "words-spacing.tsv", spacing...).Stream("갑돌씨가")
		first, _ := st.Next()
		return segmentKeys(streamWays(st, first)), st, first
	}
	on, st, first := ways("48")
	off, _, offFirst := ways()
	if on["0-2,2-4"] == 0 || off["0-2,2-4"] != 0 || first.Ways.Cmp(offFirst.Ways) <= 0 || len(on) <= len(off) {
		t.Fatal(on, off, first.Ways, offFirst.Ways)
	}
	if len(st.sent) != 1 || st.sent[0].clashes != 0 {
		t.Fatal("clashes", st.sent)
	}
}
