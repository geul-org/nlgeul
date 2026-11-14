//ff:func feature=token type=helper control=iteration dimension=1
//ff:what keyedWordsDiffer — two words of the candidates of the old reading of a sentence that have the same state key but are not the same word ("" when there are none)
package token

func keyedWordsDiffer(r *Reader, s string) string {
	for _, u := range unitsOf(r.ix, s, false) {
		if d := keyedPathsDiffer(unitScan(r.ix, u).search()); d != "" {
			return d
		}
	}
	return ""
}
