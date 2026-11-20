//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what bestNodes — the best candidate of each node of the first way of every stretch of a sentence, read by a stream to its end, in order
package parse

func bestNodes(r *Reader, text string) []Branch {
	var out []Branch
	st := r.Stream(text)
	for sg, ok := st.Next(); ok; sg, ok = st.Next() {
		for _, n := range sg.Nodes {
			out = append(out, n.Contest.Branches[0])
		}
	}
	return out
}
