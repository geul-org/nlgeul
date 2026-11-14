//ff:func feature=token type=helper control=iteration dimension=1
//ff:what extend — put in the frontier the walk taken one word further, for each word that follows its last (nearest first)
package token

func (rg *region) extend(w *walk) {
	got := w.got + rg.best[w.node]
	for _, j := range rg.next[w.node] {
		rg.front.push(&walk{prev: w, node: j, got: got, f: got + rg.rest[j]})
	}
}
