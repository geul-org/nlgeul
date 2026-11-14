//ff:func feature=token type=helper control=iteration dimension=1
//ff:what branchPref — the preference of the branch that prints as the given line, and whether there is one
package token

func branchPref(bs []Branch, line string) (float64, bool) {
	for _, b := range bs {
		if b.String() == line {
			return b.Pref, true
		}
	}
	return 0, false
}
