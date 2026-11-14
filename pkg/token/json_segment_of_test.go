//ff:func feature=token type=helper control=sequence
//ff:what TestJSONSegmentOf — the place, letters, rank, lambda and ways of a segment and a word for each node; a stretch with no node has no words
package token

import (
	"math/big"
	"testing"
)

func TestJSONSegmentOf(t *testing.T) {
	seg := Segment{Region: Span{Start: 0, End: 2}, Rank: 1, Ways: big.NewInt(3), Lambda: -1, Nodes: []Node{{Span: Span{Start: 0, End: 2}, Contest: Contest{Text: "물을"}}}}
	s := jsonSegmentOf(seg, []rune("물을"))
	if s.Region != [2]int{0, 2} || s.Text != "물을" || s.Rank != 1 || s.Ways != int64(3) || len(s.Words) != 1 || s.Lambda == nil {
		t.Fatal(s)
	}
	if s = jsonSegmentOf(Segment{Region: Span{Start: 0, End: 1}}, []rune("물")); s.Words == nil || len(s.Words) != 0 || s.Ways != int64(0) {
		t.Fatal(s)
	}
}
