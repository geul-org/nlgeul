//ff:func feature=token type=helper control=iteration dimension=1
//ff:what anyWrote — whether any of the pieces is a marker written wrong
package token

func anyWrote(pieces []piece) bool {
	for _, p := range pieces {
		if p.wrote != "" {
			return true
		}
	}
	return false
}
