//ff:func feature=token type=reader control=iteration dimension=1
//ff:what passedCandidates — the candidates of a stem from the kinds whose graphs let it stand (passed, with each kind's verdict in the order of kindOrder): a closed-set word that stands as the kind of one of its sets is read once for each such set, under the set's name; any other stem once for each kind, under the word table's kind (열쇠:체언)
package token

func passedCandidates(c *entry, passed int, vs []float64) []kindCandidate {
	var out []kindCandidate
	if own := ownNames(c, passed); len(own) > 0 {
		for _, name := range own {
			kind := closedKinds[labelOf(name)]
			out = append(out, kindCandidate{kind, name, vs[kindPlace(kind)]})
		}
		return out
	}
	for i, k := range kindOrder {
		if passed&k != 0 {
			out = append(out, kindCandidate{k, keyKind + ":" + kindLabel[k], vs[i]})
		}
	}
	return out
}
