//ff:func feature=token type=reader control=sequence
//ff:what farNew — whether a stem from b to the tail's start is too long to be a new word: more than newStemMax tokens, and not the whole of its unit (sole: the unit as one new word stays), and not the word at the unit's head (a sign of a pair or a span of code, which is no word of the text)
package token

func (sc *scan) farNew(b int, t tail, sole bool) bool {
	return !sole && (b != 0 || t.start != sc.head) && sc.tokens(b, t.start) > newStemMax
}
