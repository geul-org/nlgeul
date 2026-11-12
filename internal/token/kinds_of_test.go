//ff:func feature=token type=helper control=iteration dimension=1
//ff:what kindsOf — the first name of the first part of each reading of a bundle, as kind:form (a reading with no name gives "")
package token

func kindsOf(b Bundle) []string {
	var out []string
	for _, br := range b.Branches {
		if len(br.Words) == 0 || len(br.Words[0].Parts) == 0 || len(br.Words[0].Parts[0].Names) == 0 {
			out = append(out, "")
			continue
		}
		n := br.Words[0].Parts[0].Names[0]
		out = append(out, n.Kind+":"+n.Form)
	}
	return out
}
