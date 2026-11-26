//ff:func feature=rule type=rule control=sequence
//ff:what TestStemRun — applies to a stem of Latin letters, digits and underscores only; not to one with a Hangul syllable in it, not to an empty one
package rule

import "testing"

func TestStemRun(t *testing.T) {
	run, _ := StemRun(pieceCtx(Piece{Form: "off_2"}), nil)
	mixed, _ := StemRun(pieceCtx(Piece{Form: "off이"}), nil)
	hangul, _ := StemRun(pieceCtx(Piece{Form: "아직"}), nil)
	none, _ := StemRun(pieceCtx(Piece{}), nil)
	if run != 1 || mixed != -1 || hangul != -1 || none != -1 {
		t.Fatal(run, mixed, hangul, none)
	}
}
