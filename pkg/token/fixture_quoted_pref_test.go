//ff:func feature=token type=helper control=iteration dimension=1
//ff:what quotedPref — the preference of the candidate that prints as the given text among branches, and whether there is one
package token

func quotedPref(branches []Branch, text string) (float64, bool) {
	for _, b := range branches {
		if b.String() == text {
			return b.Pref, true
		}
	}
	return 0, false
}
