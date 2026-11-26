//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what genitiveFormsOf — the forms of the contracts given whose particle is named by one of genitiveNames (내 and 네, the pronouns 나 and 너 with the adnominal particle 의 fused)
package parse

import "slices"

func genitiveFormsOf(cs []contract) []string {
	var out []string
	for _, c := range cs {
		if slices.Contains(genitiveNames, c.name) {
			out = append(out, c.form)
		}
	}
	return out
}
