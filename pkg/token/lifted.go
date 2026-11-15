//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lifted — a node with the lifts of its span added to its candidates' preferences and the candidates in order by lambda again; the candidates are a copy, so that what the unit's reading keeps is not touched; the node as it is when no lift is of its span
package token

func lifted(n Node, lifts []lift) Node {
	var own []lift
	for _, l := range lifts {
		if l.span == n.Span {
			own = append(own, l)
		}
	}
	if len(own) == 0 {
		return n
	}
	bs := make([]Branch, len(n.Contest.Branches))
	for i, br := range n.Contest.Branches {
		bs[i] = raised(br, own)
	}
	sortByLambda(bs)
	n.Contest.Branches = bs
	return n
}
