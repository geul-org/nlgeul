//ff:func feature=token type=helper control=iteration dimension=1
//ff:what namingEnd — is r, or a key of r (the final consonant fused under it, ㅁ also under ㄻ), the last letter of a naming closing a particle chain is read right after (chainHeads: 기, 음, ㅁ, so also 됨, 함, 삶, 름): where such a chain is looked for; whether the closing before it is one of them is the graph's to judge
package token

import (
	"slices"
	"strings"
)

func namingEnd(r rune) bool {
	for _, k := range keysAt(r) {
		if slices.ContainsFunc(chainHeads, func(name string) bool { return strings.HasSuffix(name, string(k)) }) {
			return true
		}
	}
	return false
}
