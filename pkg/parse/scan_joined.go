//ff:func feature=parse type=helper control=sequence
//ff:what joined — a state before a state of the net added to its preds, once (the registry knows each edge)
package parse

func (sc *scan) joined(state, before *path) {
	k := predKey{state, before}
	reg := sc.registry()
	if reg.preds[k] {
		return
	}
	reg.preds[k] = true
	state.preds = append(state.preds, before)
}
