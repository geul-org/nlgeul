//ff:func feature=token type=helper control=iteration dimension=2
//ff:what widestWay — the widest width inside of the boxes of all the ways, so that the boxes of one contest of ways are drawn as wide as each other
package token

func widestWay(ways []boxWay) int {
	w := 0
	for _, way := range ways {
		for _, b := range way.boxes {
			w = max(w, b.width)
		}
	}
	return w
}
