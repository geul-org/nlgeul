//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what branchNamed — the candidate of a node printed as given, and its place among the node's candidates; -1 when the node has none so printed
package parse

func branchNamed(n Node, word string) (Branch, int) {
	for i, br := range n.Contest.Branches {
		if br.String() == word {
			return br, i
		}
	}
	return Branch{}, -1
}
