//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what emphKept — the chains before a particle that it may be read after: for 를 of emphasis (emphObjName) only those whose last piece is a particle of emphObjAfter (에, 으로), for any other particle all of them
package parse

import "slices"

func emphKept(e *entry, tails []tail) []tail {
	if e.name != emphObjName {
		return tails
	}
	var out []tail
	for _, t := range tails {
		if slices.Contains(emphObjAfter, t.pieces[len(t.pieces)-1].names) {
			out = append(out, t)
		}
	}
	return out
}
