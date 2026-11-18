//ff:func feature=token type=rule control=iteration dimension=1
//ff:what pronounParticled — 조건식: are the letters of a word a pronoun of the closed sets (pronouns: 이, 그, 저; personalPronouns: 나, 너, 저, 우리 …) followed by a particle of the particle tables written out in full (나는, 이를, 우리가); a pronoun alone or with letters that are not a particle is not
package token

import "strings"

func pronounParticled(text string) bool {
	for _, p := range argumentPronouns {
		rest, ok := strings.CutPrefix(text, p)
		if ok && particleForms[rest] {
			return true
		}
	}
	return false
}
