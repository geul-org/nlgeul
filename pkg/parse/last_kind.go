//ff:func feature=parse type=helper control=sequence
//ff:what lastKind — the kind of the last part of a word (끝 갈래; "" when it has no part)
package parse

func lastKind(e Eojeol) string {
	if len(e.Parts) == 0 {
		return ""
	}
	return e.Parts[len(e.Parts)-1].Name.Kind
}
