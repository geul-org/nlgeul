//ff:func feature=token type=helper control=iteration dimension=1
//ff:what standsEmpty — whether no place has a word that may stand as a word of no markers
package token

func standsEmpty(stands []map[int]bool) bool {
	for _, m := range stands {
		if len(m) != 0 {
			return false
		}
	}
	return true
}
