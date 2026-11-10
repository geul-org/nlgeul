//ff:type feature=token type=model
//ff:what word — one key of a word table: the kinds its readings fold into, its greatest observed count, and the count of each kind (by the kind's place in kindOrder)
package token

type word struct {
	kinds int
	count int
	by    [6]int
}
