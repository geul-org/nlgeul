//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what copulaNamingEnd — is r, or a key of r (the final consonant fused under it, ㅁ also under ㄻ), the last letter of a naming closing the copula is read right after (copulaNamings: 음, ㅁ, so also 됨, 함, 삶): where such a copula tail is looked for; whether the closing before it is one of them is afterNaming's to see
package parse

func copulaNamingEnd(r rune) bool {
	for _, k := range keysAt(r) {
		if lastOfAny(copulaNamings, k) {
			return true
		}
	}
	return false
}
