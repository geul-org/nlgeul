//ff:func feature=token type=helper control=iteration dimension=1
//ff:what branchNames — whether a branch holds a word whose parts have exactly the given names in order (see wordNames)
package token

func branchNames(br Branch, want string) bool {
	for _, w := range br.Words {
		if wordNames(w) == want {
			return true
		}
	}
	return false
}
