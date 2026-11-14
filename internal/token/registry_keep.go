//ff:func feature=token type=helper control=iteration dimension=1
//ff:what keep — note a state that stays after a cut in the registry again: in the order made, by its kind of word, with each state before it, and its word with the number of birth it has
package token

func (reg *registry) keep(p *path) {
	reg.all = append(reg.all, p)
	reg.states[stateKeyOf(p)] = p
	for _, b := range p.preds {
		reg.preds[predKey{p, b}] = true
	}
	k := wordKeyOf(p)
	reg.born[k], reg.lone[k] = p.last.born, true
}
