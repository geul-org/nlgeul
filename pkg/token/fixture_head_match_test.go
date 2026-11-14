//ff:func feature=token type=helper control=iteration dimension=1
//ff:what headMatch — what is wrong with a way given with a head of candidates (got) against the same way given in full (want): the place, rank, ways, lambda and nodes must be the same, and the candidates of each node exactly the first head of the full ones (all of them when head is 0)
package token

import (
	"fmt"
	"reflect"
)

func headMatch(want, got Segment, head int) []string {
	if got.Region != want.Region || got.Rank != want.Rank || got.Lambda != want.Lambda || fmt.Sprint(got.Ways) != fmt.Sprint(want.Ways) || len(got.Nodes) != len(want.Nodes) {
		return []string{fmt.Sprintf("%v rank %d: not the same way (%v rank %d)", want.Region, want.Rank, got.Region, got.Rank)}
	}
	var bad []string
	for i, n := range want.Nodes {
		keep := n.Contest.Branches
		if head > 0 && len(keep) > head {
			keep = keep[:head]
		}
		g := got.Nodes[i]
		if g.Span != n.Span || g.Contest.Unit != n.Contest.Unit || g.Contest.Glued != n.Contest.Glued || g.Contest.Text != n.Contest.Text || !reflect.DeepEqual(g.Contest.Branches, keep) {
			bad = append(bad, fmt.Sprintf("%v rank %d node %v: %d candidates, not the first %d of %d", want.Region, want.Rank, n.Span, len(g.Contest.Branches), head, len(n.Contest.Branches)))
		}
	}
	return bad
}
