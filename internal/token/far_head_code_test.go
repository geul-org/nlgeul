//ff:func feature=token type=reader control=sequence
//ff:what TestFarHeadCode — a span of code at the head of a unit that is longer than a new word's stem, with letters written onto it, is read by the stream and has its candidate; a sign of a pair with a long run of letters after it is read by the stream too (the old search, which spreads every path, is only given short units)
package token

import (
	"strings"
	"testing"
)

func TestFarHeadCode(t *testing.T) {
	r := streamReader(t, "words-far.tsv")
	in := "`" + strings.Repeat("가", 30) + "`를마신다"
	if !netHas(r, in, "[코드]+를[목적격:를]") {
		t.Fatal("a long span of code with a particle onto it")
	}
	st := r.Stream(in)
	if seg, ok := st.Next(); !ok || seg.Ways == nil {
		t.Fatal("the stream", seg)
	}
	onto := ")" + strings.Repeat("고양이가마당에서물을마신다", 4)
	if len(netRegions(r, onto)) == 0 {
		t.Fatal("a sign of a pair with a long run after it")
	}
	if !unitHas(r, "`가가`를마신다", "[코드]+를[목적격:를]") {
		t.Fatal("a short span of code with a particle onto it, the old search")
	}
}
