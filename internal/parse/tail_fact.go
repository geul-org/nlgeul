//ff:type feature=parse type=model
//ff:what tailFact — one reading of the word's end as a predicate tail: the closing and its kind, the stem before it, the length read, whether the copula stem 이 was read, the observed count, and the word divided into its pieces
package parse

// closing is the closing's name with its lead written out (ㄴ다, 어서,
// ㅂ니다). size is in half syllables: 2 for every syllable of the
// closing, 요, a slot before the closing and the copula stem, 1 for a
// piece fused into the stem's last syllable (0: no tail). stem is the
// text before the tail with a fused piece taken out of its last
// syllable. count is the table's count of the copula when its stem was
// read, else of the closing; 0 when the table gives none. marked is
// the whole word with "+" between the pieces this reading found.
type tailFact struct {
	closing string
	kind    int
	stem    string
	size    int
	copula  bool
	count   int
	marked  string
}
