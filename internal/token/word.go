//ff:type feature=token type=model
//ff:what word — one key of a word table: the kinds its readings fold into, its greatest observed count, the count of each kind (by the kind's place in kindOrder), and the irregular types its predicate rows have (irr…)
package token

type word struct {
	kinds int
	count int
	by    [6]int
	irr   int
}
