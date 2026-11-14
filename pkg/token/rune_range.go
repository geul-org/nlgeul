//ff:type feature=token type=model
//ff:what runeRange — the letters from Lo up to Hi, both included (a run of the code points a cell width is told for)
package token

type runeRange struct {
	Lo rune
	Hi rune
}
