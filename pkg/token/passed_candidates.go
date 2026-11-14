//ff:func feature=token type=reader control=iteration dimension=1
//ff:what passedCandidates — the candidates of a stem from the kinds whose graphs let it stand (passed, with each kind's verdict in the order of kindOrder): a closed-set word that stands as the kind of one of its sets is read once for each such set, under the set's name, and also once under the word table's kind (열쇠:체언) for each kind passed that the word table lists (listed) for it; any other stem once for each kind, under the word table's kind
package token

// The names of the closed sets come first, then the word table's kinds in
// the order of kindOrder. A closed-set word's own name never hides the
// word table's: 수 is read as the dependent noun (수[의존명사:수]) and as the
// noun (수[열쇠:체언]) of the table, and 있 as the auxiliary verb and as the verb
// and the adjective (있[열쇠:동사], 있[열쇠:형용사]). Each is judged by its kind's
// graph, so the two names of one kind have the same verdict and preference.
// A kind the table does not list for the word (a set's kind that only the
// grammar tables give it) is named by the set alone.
func passedCandidates(c *entry, passed, listed int, vs []float64) []kindCandidate {
	var out []kindCandidate
	own := ownNames(c, passed)
	for _, name := range own {
		kind := closedKinds[labelOf(name)]
		out = append(out, kindCandidate{kind, name, vs[kindPlace(kind)]})
	}
	for i, k := range kindOrder {
		if passed&k != 0 && (len(own) == 0 || listed&k != 0) {
			out = append(out, kindCandidate{k, keyKind + ":" + kindLabel[k], vs[i]})
		}
	}
	return out
}
