//ff:func feature=token type=helper control=iteration dimension=1
//ff:what linksOf — for each node, the places of the nodes that follow it (by the edges), the nearest first
package token

import "sort"

func linksOf(nodes []Node, edges []edge) [][]int {
	at := map[Span]int{}
	for i, n := range nodes {
		at[n.Span] = i
	}
	next := make([][]int, len(nodes))
	for _, e := range edges {
		next[at[e.from]] = append(next[at[e.from]], at[e.to])
	}
	for _, ids := range next {
		sort.Ints(ids)
	}
	return next
}
