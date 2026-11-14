//ff:func feature=token type=helper control=iteration dimension=1
//ff:what chainWalk — a walk through the nodes of the places given, first to last
package token

func chainWalk(ids ...int) *walk {
	var w *walk
	for _, id := range ids {
		w = &walk{prev: w, node: id}
	}
	return w
}
