//ff:func feature=token type=helper control=sequence
//ff:what headed — does a candidate begin with the sign of a pair at its unit's head, named as such (pairName) while the unit was read
package token

func headed(br Branch) bool {
	if len(br.Words) == 0 || len(br.Words[0].Parts) == 0 {
		return false
	}
	return br.Words[0].Parts[0].Name.Kind == pairName
}
