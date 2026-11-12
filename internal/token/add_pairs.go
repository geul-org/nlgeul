//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addPairs — a graph for each candidate of each sign of a pair (pairRows): one for a sign that opens, one for a sign that closes, both for a quotation mark written the same way at both ends
package token

// An entry's form tells which candidate it is (열기, 닫기); its name is
// the graph's, with the sign.
func (ix *index) addPairs() {
	ix.pairs = map[rune][]*entry{}
	for _, sign := range pairOrder {
		row := pairRows[sign]
		if row.opens {
			ix.addPair(sign, row, openName)
		}
		if row.shuts {
			ix.addPair(sign, row, shutName)
		}
	}
}
