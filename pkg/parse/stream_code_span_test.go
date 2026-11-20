//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestStreamCodeSpan — the places a stream sends for a sentence with a span of code that holds spaces leave those spaces out: one, two in a row, several spans, a span in brackets and quotes, a span that starts a unit with more words after it; Again of a stretch with several ways after a span; a span with no space is as it was; the words keep the spaces in their letters
package parse

import "testing"

func TestStreamCodeSpan(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, s := range []string{
		"`go test all`을 돌리고 밥을 먹는다.",
		"`a b`이라고했는데요 `c d e`에서",
		"(`a b`)와 \"`c d`\"를",
		"가`e  f`나다",
		"`gotest`를 돌린다",
		"아버지가 방에 들어가신다",
	} {
		codeSpanFit(t, r, s)
	}
	st := r.Stream("`go test all`을 돌리고 밥을 먹는다.")
	seg, ok := st.Next()
	if !ok || seg.Region != (Span{0, 12}) || seg.Nodes[0].Contest.Text != "`go test all`을" {
		t.Fatal(seg, ok)
	}
	if next, ok := st.Next(); !ok || next.Region != (Span{12, 15}) {
		t.Fatal(next, ok)
	}
	many := r.Stream("`a b` 돌아가셨다")
	many.Next()
	first, ok := many.Next()
	again, more := many.Again(first.Region)
	if !ok || !more || first.Region != (Span{4, 9}) || again.Region != first.Region || again.Nodes[0].Span != (Span{4, 6}) {
		t.Fatal(first, again, ok, more)
	}
	plain := r.Stream("`gotest`를 돌린다")
	if plain.places != nil {
		t.Fatal(plain.places)
	}
}
