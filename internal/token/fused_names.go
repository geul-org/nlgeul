//ff:func feature=token type=reader control=iteration dimension=2
//ff:what fusedNames — the names of the index's markers that fuse into the vowel of the syllable before them (어/아 put into the vowel, ㅆ put under it), each once, sorted
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (ix *index) fusedNames() []string {
	var out []string
	for _, entries := range ix.byKey {
		for _, e := range entries {
			if (e.lead == rule.LeadVowel || e.lead == rule.JongSS) && !slices.Contains(out, e.name) {
				out = append(out, e.name)
			}
		}
	}
	slices.Sort(out)
	return out
}
