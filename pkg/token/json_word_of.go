//ff:func feature=token type=helper control=iteration dimension=1
//ff:what jsonWordOf — a node as it is written in JSON: its place, its letters, whether it is glued, and every candidate of its contest in the order they have (nothing is cut)
package token

func jsonWordOf(n Node) JSONWord {
	out := JSONWord{Span: [2]int{n.Span.Start, n.Span.End}, Text: n.Contest.Text, Glued: n.Contest.Glued, Candidates: make([]JSONCandidate, 0, len(n.Contest.Branches))}
	for _, br := range n.Contest.Branches {
		out.Candidates = append(out.Candidates, jsonCandidateOf(br))
	}
	return out
}
