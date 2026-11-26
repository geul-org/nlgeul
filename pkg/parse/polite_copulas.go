//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what politeCopulas — the tails of the copula stem written before an old polite slot (사람+이+옵+니다): those of copulas (the stem stands in the letters: a stem left out is not read before such a slot), but not where the stem and the syllables of the slot are together a word of the word table as they are written (이온: the table's word is read)
package parse

import "unicode/utf8"

// A copula stem left out of the writing is not read before a polite
// slot: the dropped stem's own graph turns it down, since a polite slot
// is none of the pieces it is left out before. The slot is the piece
// right after the stem; a tail with another piece there (a slot before
// the polite one) is kept as it is.
func (sc *scan) politeCopulas(tails []tail, chained bool) []tail {
	var out []tail
	for _, t := range sc.copulas(tails, chained) {
		end := t.start + 1 + utf8.RuneCountInString(t.pieces[1].text)
		if _, word := sc.ix.lex[sc.letters(t.start, min(end, len(sc.s)))]; word && sc.ix.politeNamed(t.pieces[1].names) {
			continue
		}
		out = append(out, t)
	}
	return out
}
