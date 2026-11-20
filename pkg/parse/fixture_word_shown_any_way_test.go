//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what wordShownAnyWay — the candidates (as jsonWordShown sets them down) of the first node whose letters are the text given, looked for in every way of every stretch of a sentence (Next, then Again), the best way's first; and whether there is one
package parse

func wordShownAnyWay(r *Reader, sentence, text string) ([]string, bool) {
	var ways []Segment
	st := r.Stream(sentence)
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		ways = append(ways, streamWays(st, seg)...)
	}
	for _, way := range ways {
		for _, n := range way.Nodes {
			if n.Contest.Text == text {
				return jsonWordShown(jsonWordOfNode(n)), true
			}
		}
	}
	return nil, false
}
