//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stemIn — what stemReads gives for a tail as a test reads it: the stem of the first candidate, the names of all the candidates joined with ",", their best verdict, and the letters the stem is written with
package token

func stemIn(sc *scan, b int, t tail, sole bool) (string, string, float64, string) {
	reads := sc.stemReads(b, t, sole)
	kinds := make([]kindRead, len(reads))
	for i, r := range reads {
		kinds[i] = r.kindRead
	}
	names, best := readsIn(kinds)
	return reads[0].stem, names, best, reads[0].wrote
}
