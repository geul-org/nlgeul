//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sharedGate — the graph the word before a particle must pass, when every reading of the particle's form that stands at a place asks for the same one; nil when one of them asks for none (the word before then stands by that reading)
package token

func sharedGate(entries []*entry) *entry {
	gate := entries[0].gate
	for _, e := range entries[1:] {
		if e.gate != gate {
			return nil
		}
	}
	return gate
}
