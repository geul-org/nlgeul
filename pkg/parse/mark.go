//ff:func feature=parse type=rule control=sequence
//ff:what mark — the word with "+" between its pieces: divided at the particle chain when a particle reading stands, else as the standing tail reading divides it; with no reading standing, a dependent noun written onto a modifier is still set off
package parse

func mark(w *word, r *ruling) string {
	if r.chain || (r.tail == "" && w.chain.kase == caseBare) {
		return w.chain.marked
	}
	if r.tail != "" {
		return r.tail
	}
	return w.text
}
