//ff:func feature=token type=helper control=iteration dimension=1
//ff:what withGate — tails with the graph the word before them must pass, when there is one
package token

func withGate(tails []tail, gate *entry) []tail {
	if gate == nil {
		return tails
	}
	for i := range tails {
		tails[i].gate = gate
	}
	return tails
}
