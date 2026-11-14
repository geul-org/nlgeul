//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what joinAt — try every joined dependent noun at s[k:] and add what is read to the word: a longer chain replaces the word's, tails are added
package parse

func joinAt(w *word, s []rune, k int) {
	for _, noun := range joinedNouns {
		c, tails := joinNoun(s, k, noun)
		if c.size > w.chain.size {
			w.chain = c
		}
		w.tails = append(w.tails, tails...)
	}
}
