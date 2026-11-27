//ff:func feature=parse type=helper control=sequence
//ff:what orderJoint — is the first piece of a tail the letter 제 with a number right after it (요추제1, 갑제13): there the letter is the prefix of the number (제1), not a suffix of the word before it
package parse

import "unicode"

func (sc *scan) orderJoint(t tail) bool {
	if len(t.pieces) == 0 || t.pieces[0].text != string(orderHead) {
		return false
	}
	next := sc.at(t.start + 1)
	return unicode.IsDigit(next) || numberLetter(next)
}
