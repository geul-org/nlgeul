//ff:func feature=token type=helper control=iteration dimension=2
//ff:what flatNodes — the nodes of a sentence read by a stream to its end, one line each (the span and the best candidate), whatever the stretches they came in; with the window on or off
package token

import "fmt"

func flatNodes(r *Reader, text string, window bool) []string {
	r.ix.noWindow = !window
	defer func() { r.ix.noWindow = false }()
	var out []string
	st := r.Stream(text)
	for sg, ok := st.Next(); ok; sg, ok = st.Next() {
		for _, n := range sg.Nodes {
			out = append(out, fmt.Sprintf("%v %s", n.Span, n.Contest.Branches[0]))
		}
	}
	return out
}
