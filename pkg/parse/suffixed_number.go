//ff:func feature=parse type=helper control=sequence
//ff:what suffixedNumber — is the word before a tail a number a suffix of numberSuffixes (사) is read right after: its stem is named as a number is (numberKinds: 수, 차례), a number ends where the tail starts, and the tail's first piece is such a suffix (3사, 제2사)
package parse

import "slices"

func (sc *scan) suffixedNumber(t tail, e step) bool {
	if len(t.pieces) == 0 || !sc.numEnds[t.start] || !slices.Contains(numberSuffixes, t.pieces[0].text) {
		return false
	}
	return slices.Contains(numberKinds, e.word.Parts[0].Name.Kind)
}
