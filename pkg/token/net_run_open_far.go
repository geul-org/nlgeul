//ff:func feature=token type=reader control=iteration dimension=2
//ff:what openFar — mark the places a known word farther than a new word's stem may be could begin at, that ends less than windowSize tokens from the place worked out (the places of the tables' long keys, as the scan knows them: begins)
package token

import "sort"

func (r *netRun) openFar(flags []bool) {
	first := r.c + 1 + sort.Search(r.j-r.c, func(i int) bool { return r.sc.tokens(r.c+1+i, r.j) < windowSize })
	for e := first; e <= r.j; e++ {
		for _, b := range r.sc.begins[e] {
			if b >= r.c && r.sc.tokens(b, e) > newStemMax {
				flags[b] = true
			}
		}
	}
}
