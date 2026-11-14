//ff:func feature=token type=helper control=iteration dimension=1
//ff:what wordNames — the names of a word's parts joined by +: each part's first name as kind:form, a part with none as its text
package token

import "strings"

func wordNames(w Eojeol) string {
	var got []string
	for _, p := range w.Parts {
		if p.Name == (Name{}) {
			got = append(got, p.Text)
			continue
		}
		got = append(got, p.Name.Kind+":"+p.Name.Form)
	}
	return strings.Join(got, "+")
}
