//ff:func feature=token type=helper control=sequence
//ff:what byEnd — the states kept in the order of where they end, those that end at the same place as they were made (every state before a state ends before it begins to end)
package token

import "sort"

func (r *netRun) byEnd() []*path {
	order := append([]*path{}, r.sc.registry().all...)
	sort.SliceStable(order, func(i, j int) bool { return order[i].last.end < order[j].last.end })
	return order
}
