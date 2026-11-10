//ff:type feature=token type=model
//ff:what word — one key of a word table: the kinds its readings fold into and its observed count
package token

type word struct {
	kinds int
	count int
}
