//ff:func feature=token type=helper control=iteration dimension=1
//ff:what pooled — the candidates of one span with a candidate printed the same as one before it left out (the first kept, in the order given: as poolOf keeps the first of the ways of cutting), and how many words printed the same came with another verdict or preference
package token

func pooled(bs []Branch) ([]Branch, int) {
	seen := map[string]Branch{}
	clashed := map[string]bool{}
	var out []Branch
	for _, b := range bs {
		text := b.String()
		first, dup := seen[text]
		if !dup {
			seen[text] = b
			out = append(out, b)
		} else if first.Verdict != b.Verdict || first.Pref != b.Pref {
			clashed[text] = true
		}
	}
	return out, len(clashed)
}
