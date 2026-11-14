//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sameStrings — whether two lists of strings are the same, in order
package token

func sameStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
