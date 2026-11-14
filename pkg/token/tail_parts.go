//ff:func feature=token type=helper control=sequence
//ff:what tailParts — the parts of a word after its stem (all of them when it has no stem)
package token

func tailParts(e Eojeol) []Part {
	if len(e.Parts) > 0 && e.Parts[0].Stem {
		return e.Parts[1:]
	}
	return e.Parts
}
