//ff:type feature=token type=model
//ff:what lattice — what the net of a unit gives: the states that end the unit (at its last letter, and the ones that are markers written apart from the unit before), every state in the order it was made, how many words a state-making read found with the same print and other values (twins), and the letters of the unit
package token

type lattice struct {
	n     int
	ends  []*path
	all   []*path
	twins int
}
