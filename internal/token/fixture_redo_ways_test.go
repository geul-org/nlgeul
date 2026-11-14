//ff:func feature=token type=helper control=iteration dimension=1
//ff:what redoWays — the first n ways of the third stretch (there are several ways of it) a window cuts off a stream of 아버지가방에들어가신다 six times over words-stream.tsv, as the stretch itself gives them before the stream keeps only what is needed to read it again (the first of them is what Next sends)
package token

import (
	"strings"
	"testing"
)

var redoText = strings.Repeat("아버지가방에들어가신다", 6)

func redoWays(t *testing.T, r *Reader, n int) []string {
	t.Helper()
	st := r.Stream(redoText)
	st.Next()
	st.Next()
	st.fill()
	rg := st.queue[0]
	if rg.redo == nil || rg.span != (Span{6, 11}) {
		t.Fatal("not the stretch cut off by a window", rg.span, strings.Count(redoText, "아"))
	}
	var out []string
	for i := 0; i < n; i++ {
		seg, ok := rg.deal()
		if !ok {
			break
		}
		out = append(out, segmentLine(seg))
	}
	return out
}
