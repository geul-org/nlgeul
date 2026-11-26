//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what politeTakersOf — the names of the polite slots that take the link vowel 으 before them: those that stand after a vowel (오, 옵)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func politeTakersOf(rows []entry) []string {
	var out []string
	for _, e := range rows {
		if e.after == rule.AfterVowel {
			out = append(out, e.name)
		}
	}
	return out
}
