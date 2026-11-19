//ff:func feature=token type=helper control=iteration dimension=1
//ff:what adnominalAll — is each candidate of a word an adnominal (adnominal: a modifying form, a determiner standing alone)
package token

func adnominalAll(n Node) bool {
	for _, br := range n.Contest.Branches {
		if len(br.Words) == 0 || !adnominal(br.Words[0].Parts) {
			return false
		}
	}
	return true
}
