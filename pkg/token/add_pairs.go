//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addPairs — a graph for each candidate of each sign of a pair (pairRows): one for a sign that opens, one for a sign that closes, both for a quotation mark written the same way at both ends, and after them one more closing for a quotation mark that may close with none of its kind open (loneShuts)
package token

import "strings"

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
		if row.shuts && strings.ContainsRune(loneShuts, sign) {
			ix.addLoneShut(sign, row)
		}
	}
}
