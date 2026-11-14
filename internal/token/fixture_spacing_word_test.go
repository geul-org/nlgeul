//ff:func feature=token type=helper control=iteration dimension=1
//ff:what spacingWord — a word of a candidate made by hand for the spacing rules: each spec is the letters, a slash and the name (kind:form) of a part; the first is the stem, and a spec that begins with "-" is a marker (a word whose first spec begins with "-" has no stem)
package token

import "strings"

func spacingWord(specs ...string) Eojeol {
	var e Eojeol
	for i, spec := range specs {
		spec, marker := strings.CutPrefix(spec, "-")
		text, name, _ := strings.Cut(spec, "/")
		kind, form, _ := strings.Cut(name, ":")
		e.Parts = append(e.Parts, Part{Text: text, Name: Name{Kind: kind, Form: form}, Stem: i == 0 && !marker})
	}
	return e
}
