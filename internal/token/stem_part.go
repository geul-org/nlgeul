//ff:func feature=token type=helper control=sequence
//ff:what stemPart — the stem of a word (its first part when that is a stem); the empty part when the word has no stem (a word of markers alone)
package token

func stemPart(e Eojeol) Part {
	if len(e.Parts) == 0 || !e.Parts[0].Stem {
		return Part{}
	}
	return e.Parts[0]
}
