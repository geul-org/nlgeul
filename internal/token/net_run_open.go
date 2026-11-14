//ff:func feature=token type=reader control=iteration dimension=1
//ff:what open — for each place from the last cut to the place worked out whether a state that ends there is part of the window's end: the place lies less than windowSize tokens from the place worked out, or a key of the tables longer than a new word's stem may be still fits the letters from it (keyPrefixes; no key is longer than keyMax letters), or a known word that long may end in the window's last windowSize tokens from it (openFar)
package token

func (r *netRun) open() []bool {
	flags := make([]bool, r.j+1)
	for p := r.c; p <= r.j; p++ {
		flags[p] = r.sc.tokens(p, r.j) < windowSize || (p+r.sc.ix.keyMax > r.j && r.sc.ix.keyPrefixes[r.sc.letters(p, r.j+1)])
	}
	r.openFar(flags)
	return flags
}
