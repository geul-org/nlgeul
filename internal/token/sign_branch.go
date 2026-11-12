//ff:func feature=token type=helper control=sequence
//ff:what signBranch — a sign as a reading of its own: one word of one part, which stands (verdict 1) and costs nothing in preference
package token

func signBranch(p Part, end int) Branch {
	return Branch{Words: []Eojeol{{Parts: []Part{p}, End: end}}, Verdict: 1}
}
