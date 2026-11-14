//ff:func feature=token type=reader control=iteration dimension=1
//ff:what boxLinesOf — what is sent up by a stream drawn: for each stretch, the trace lines of the units read for it, then its boxes (the ways of a stretch cut more than one way up to head of them, 0 for all); last the line "1등 길:" with the letters of the best way of each stretch and the sum of their lambdas. Only the stream is asked: the letters of the sentence are used for a stretch nothing is known of
package token

import (
	"fmt"
	"strings"
)

func boxLinesOf(st wayStream, letters []rune, head int) []string {
	var out, path []string
	total := 0.0
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		out = append(out, st.Notes()...)
		out = append(out, boxRegionLines(takeWays(st, seg, head), letters, head)...)
		path = append(path, pathOf(seg, letters)...)
		total += seg.Lambda
	}
	out = append(out, st.Notes()...)
	if len(path) == 0 {
		return out
	}
	return append(out, fmt.Sprintf("1등 길: %s  (λ 합 %.2f)", strings.Join(path, " │ "), total))
}
