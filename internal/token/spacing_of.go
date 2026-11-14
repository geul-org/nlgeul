//ff:func feature=token type=reader control=iteration dimension=1
//ff:what spacingOf — the names of the spacing rules that are on and that a boundary is a place of, in the order of SpacingArticles (none when no rule is on)
package token

func (ix *index) spacingOf(g gap) []string {
	var out []string
	for _, r := range spacingRules {
		if ix.spacing[r.name] && r.fits(g) {
			out = append(out, r.name)
		}
	}
	return out
}
