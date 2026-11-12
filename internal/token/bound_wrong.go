//ff:func feature=token type=rule control=iteration dimension=1
//ff:what boundWrong — whether the stem is words bound by a slash or a hyphen and a word of it is a wrong way to write another word of the word table
package token

func (ix *index) boundWrong(stem string) bool {
	parts := partsBound(stem)
	if len(parts) < 2 {
		return false
	}
	for _, part := range parts {
		if len(ix.typos[part]) > 0 {
			return true
		}
	}
	return false
}
