//ff:func feature=token type=helper control=sequence
//ff:what withHead — a copy of a reading whose first part (the sign of a pair at its unit's head) has the names and the depth given; the reading itself is not changed
package token

func withHead(br Branch, names []Name, depth int) Branch {
	out := br
	out.Words = append([]Eojeol{}, br.Words...)
	out.Words[0].Parts = append([]Part{}, br.Words[0].Parts...)
	out.Words[0].Parts[0].Names, out.Words[0].Parts[0].Depth = names, depth
	return out
}
