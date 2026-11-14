//ff:type feature=token type=model
//ff:what portion — the states of the net that lie between two places every way of reading the unit passes (from, to, counted in the letters of the unit): the states that end a way and are on one, in the order they were made
package token

type portion struct {
	from   int
	to     int
	states []*path
}
