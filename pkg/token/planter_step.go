//ff:func feature=token type=reader control=iteration dimension=1
//ff:what step — work out the places of the whole unit up to j (netStep) and note the states that end there and are the sign alone as the first word (no part of the word is left when the sign is taken off)
package token

func (pl *planter) step(j int) {
	pl.sc.netStep(pl.best, j, nil)
	for _, q := range pl.best[j] {
		if _, ok := pl.firstOf(q); !ok && pl.first(q) {
			pl.transparent[q] = true
		}
	}
}
