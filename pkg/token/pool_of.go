//ff:func feature=token type=helper control=iteration dimension=2
//ff:what poolOf — gather the candidates of each span from the ways of cutting: the words that stand at a span in any way, a word printed the same once (the first one kept); the places where a word printed the same came with another verdict or preference are counted once each
package token

func poolOf(cuts []cut) pool {
	p := pool{words: map[Span][]Branch{}}
	seen := map[slotKey]Branch{}
	clashed := map[slotKey]bool{}
	for _, c := range cuts {
		for i, sp := range c.spans {
			w := c.words[i]
			k := slotKey{span: sp, text: w.String()}
			first, dup := seen[k]
			if !dup {
				seen[k] = w
				p.words[sp] = append(p.words[sp], w)
			}
			if dup && (first.Verdict != w.Verdict || first.Pref != w.Pref) {
				clashed[k] = true
			}
		}
	}
	p.clashes = len(clashed)
	return p
}
