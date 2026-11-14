//ff:func feature=token type=helper control=sequence
//ff:what registered — a candidate taken as a new state of the net: it stands for its word with only the state before it in preds, begins where that ends, and is printed as its word alone (the print of the way to it is not kept); a word printed the same as one already a state, of other values, is counted as a twin
package token

func (sc *scan) registered(p *path, k stateKey) *path {
	reg := sc.registry()
	p.preds, p.start, p.sig = []*path{sc.stateBefore(p)}, k.word.start, p.last.sig
	p.prev = p.preds[0]
	if reg.lone[k.word] {
		reg.twins++
	}
	reg.lone[k.word] = true
	reg.states[k] = p
	reg.all = append(reg.all, p)
	return p
}
