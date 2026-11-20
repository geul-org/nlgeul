//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what streamLines — what the stream of a sentence gives, as lines: every way of every stretch (Next, then Again until there is none), stretch by stretch and way by way, each way as wayChoices prints it (one candidate of each node, joined with " "), so the first line is the first candidates of the best way of the first stretch
package parse

func streamLines(r *Reader, sentence string) []string {
	var out []string
	st := r.Stream(sentence)
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		for _, way := range streamWays(st, seg) {
			out = append(out, wayChoices(way)...)
		}
	}
	return out
}
