//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what complementNodes — the nodes a stream sends, one per word in order, and the trace lines (Notes) that came with each of them, each joined with newlines
package parse

import "strings"

func complementNodes(st *Stream) ([]Node, []string) {
	var nodes []Node
	var notes []string
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		for _, n := range seg.Nodes {
			nodes = append(nodes, n)
			notes = append(notes, strings.Join(st.Notes(), "\n"))
		}
	}
	return nodes, notes
}
