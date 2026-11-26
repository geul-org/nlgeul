//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what endsNamed — may one of some closings (by their names: namedEnds of a particle) end in the letter r: r is the last letter of one, or one is fused into the syllable before it and shows no letter of its own (어(합침)); where a particle of closingParticles is looked for, and whether the closing before it is one of them is the graph's to judge
package parse

import "strings"

func endsNamed(names []string, r rune) bool {
	for _, name := range names {
		if strings.HasSuffix(name, string(r)) || strings.HasSuffix(name, "(합침)") {
			return true
		}
	}
	return false
}
