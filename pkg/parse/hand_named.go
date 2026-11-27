//ff:func feature=parse type=helper control=sequence
//ff:what handNamed — does a hand-written table read the sign: the sign table (a comma, a hyphen, the full stop, = read apart), the table of pairs (a bracket, a quotation mark) or the unit table (%, °); such a sign is not given the kind it would fall into by default
package parse

func (ix *index) handNamed(r rune) bool {
	_, pair := pairRows[r]
	return pair || len(ix.marks[r]) > 0 || len(ix.unitEntries(string(r))) > 0
}
