//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addHeads — the particles among some entries, kept by the first letter of their forms
package token

func (ix *index) addHeads(list []*entry) {
	for _, e := range list {
		if f := []rune(e.form); e.class == classParticle {
			ix.heads[f[0]] = append(ix.heads[f[0]], e.form)
		}
	}
}
