//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what jsonWordOfNode — a node's candidates set down as data (jsonCandidateOf), in their order, as the word of a sentence set down as data would hold them
package parse

func jsonWordOfNode(n Node) JSONWord {
	w := JSONWord{Text: n.Contest.Text}
	for _, br := range n.Contest.Branches {
		w.Candidates = append(w.Candidates, jsonCandidateOf(br))
	}
	return w
}
