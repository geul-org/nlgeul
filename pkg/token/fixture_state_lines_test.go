//ff:func feature=token type=helper control=iteration dimension=2
//ff:what stateLines — every way through the net to a state, from the root (the state with nothing before it) by the states before each (preds), each printed as its words (Eojeol.String) joined with " "
package token

func stateLines(p *path) []string {
	if p.prev == nil {
		return []string{""}
	}
	var out []string
	for _, before := range p.preds {
		for _, line := range stateLines(before) {
			out = append(out, line+map[bool]string{true: " ", false: ""}[line != ""]+p.last.word.String())
		}
	}
	return out
}
