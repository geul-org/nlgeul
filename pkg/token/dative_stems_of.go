//ff:func feature=token type=helper control=iteration dimension=1
//ff:what dativeStemsOf — the words of the contracts given whose particle is named by one of dativeNames (the pronoun forms 내, 네 and 제 of 내게, 네게 and 제게, and the pronouns 자네 and 자기 of 자네게 and 자기게)
package token

import "slices"

func dativeStemsOf(cs []contract) []string {
	var out []string
	for _, c := range cs {
		if slices.Contains(dativeNames, c.name) {
			out = append(out, c.stem)
		}
	}
	return out
}
