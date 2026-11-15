//ff:func feature=token type=helper control=iteration dimension=1
//ff:what redoWaysRuled — redoWays with rules of a sentence: the first n ways of the third stretch a window cuts off a stream of 아버지가방에들어가신다 six times over words-stream.tsv, as the stretch gives them after the rules have judged it (the first of them is what Next sends)
package token

import "testing"

func redoWaysRuled(t *testing.T, r *Reader, rules []sentenceRule, n int) []Segment {
	t.Helper()
	st := r.Stream(redoText)
	st.rules = rules
	st.Next()
	st.Next()
	st.fill()
	rg := st.queue[0]
	if rg.redo == nil || rg.span != (Span{6, 11}) || !rg.judged {
		t.Fatal("not the stretch cut off by a window", rg.span)
	}
	var out []Segment
	for i := 0; i < n; i++ {
		seg, ok := rg.deal()
		if !ok {
			break
		}
		out = append(out, seg)
	}
	return out
}
