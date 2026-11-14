//ff:func feature=token type=helper control=sequence
//ff:what withPart — a copy of a candidate whose word k has the one part given in place of its own; the candidate itself is not changed
package token

func withPart(br Branch, k int, p Part) Branch {
	out := br
	out.Words = append([]Eojeol{}, br.Words...)
	out.Words[k].Parts = []Part{p}
	return out
}
