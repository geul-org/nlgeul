//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what copulas — the copula candidates of some predicate tails: each tail that may stand after the stem 이 (said or dropped: copulaOf) read with the stem as a piece before it (copulaReads); with a word table a whole tail right after a syllable 이 is read once more with the stem dropped after that syllable, for the syllable may be the last of a nominal (아이+(이)+면, 차이+(이)+야: droppedToo)
package parse

// A dropped stem is written (이) and leaves no trace in the letters.
func (sc *scan) copulas(tails []tail, chained bool) []tail {
	var out []tail
	for _, t := range tails {
		start, e, ok := sc.copulaOf(t)
		if !ok {
			continue
		}
		out = append(out, sc.copulaReads(t, start, e, chained)...)
		if sc.droppedToo(t, e) {
			out = append(out, sc.copulaReads(t, t.start, sc.ix.dropped, chained)...)
		}
	}
	return out
}
