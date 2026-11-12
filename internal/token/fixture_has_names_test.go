//ff:func feature=token type=helper control=iteration dimension=2
//ff:what hasNames — whether one branch of some contest holds a word whose parts have exactly the given names in order (see wordNames)
package token

func hasNames(bs []Contest, want string) bool {
	for _, b := range bs {
		for _, br := range b.Branches {
			if branchNames(br, want) {
				return true
			}
		}
	}
	return false
}
