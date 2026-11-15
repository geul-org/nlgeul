//ff:func feature=token type=helper control=iteration dimension=2
//ff:what wayChoices — a way printed every way one candidate of each of its nodes can be chosen: the chosen candidates (Branch.String) joined with " ", the first node's candidates outermost, each in its order (all the first candidates first); a way with no nodes gives none
package token

func wayChoices(seg Segment) []string {
	if len(seg.Nodes) == 0 {
		return nil
	}
	out := []string{""}
	for i, n := range seg.Nodes {
		var next []string
		for _, prefix := range out {
			for _, b := range n.Contest.Branches {
				next = append(next, prefix+map[bool]string{true: " ", false: ""}[i > 0]+b.String())
			}
		}
		out = next
	}
	return out
}
