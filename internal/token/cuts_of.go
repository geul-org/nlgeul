//ff:func feature=token type=helper control=iteration dimension=1
//ff:what cutsOf — the ways of cutting a contest's stretch: one for each of its branches, in the branches' order (at is where the stretch begins in the sentence, counted without spaces)
package token

func cutsOf(c Contest, at int) []cut {
	n := len([]rune(c.Text))
	out := make([]cut, 0, len(c.Branches))
	for _, br := range c.Branches {
		out = append(out, cutOf(br, at, n))
	}
	return out
}
