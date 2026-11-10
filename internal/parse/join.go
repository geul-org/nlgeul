//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what join — add to the word the readings of a dependent noun or a shortened 것 written onto what stands before it
package parse

func join(w *word, s []rune) {
	if c, ok := joinShort(s); ok && c.size > w.chain.size {
		w.chain = c
	}
	for k := 1; k < len(s); k++ {
		joinAt(w, s, k)
	}
}
