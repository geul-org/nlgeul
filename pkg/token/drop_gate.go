//ff:func feature=token type=reader control=sequence
//ff:what dropGate — the graph the word before a dropped copula stem must pass before a closing of droppableKnown: a nominal of the word table or a new word before the closings of droppableNew (nominalNew), a nominal of the word table before the others (nominal)
package token

import "slices"

func (ix *index) dropGate(closing string) *entry {
	if slices.Contains(droppableNew, closing) {
		return ix.nominalNew
	}
	return ix.nominal
}
