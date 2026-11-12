//ff:func feature=token type=reader control=sequence
//ff:what numberReads — the ways a word is read whose stem from b to stemEnd is a number: the word as it is and, when the stem is one run of the digits 0 and 1 only (named 수, with no sign before it and no sign inside), the same word again with its stem named 수:2진법; the order is the number first
package token

// Nothing else of the second reading differs: its verdict and
// preference are the first reading's, so that the two stand together
// and the stable order of equal readings puts the first one first. A
// number joined by a sign (1,000, 1.01, 10:10) or with a sign or money
// before it is named by its signs' graphs and is read once.
func (sc *scan) numberReads(e step, b, stemEnd int) []step {
	a, ok := sc.atoms[b]
	if !ok || a.end != stemEnd || !a.num || len(a.names) != 1 || a.names[0] != numName || !binaryRun(string(sc.s[b:stemEnd])) || len(e.word.Parts) == 0 || e.verdict <= 0 {
		return []step{e}
	}
	two := e
	two.word.Parts = append([]Part{}, e.word.Parts...)
	two.word.Parts[0].Name = Name{numName, binForm}
	two.sig = two.word.String()
	return []step{e, two}
}
