//ff:type feature=token type=model
//ff:what restore — one way of taking a fused piece out of a stem's last syllable: the syllable as the stem had it, and the piece as it was before fusing
package token

type restore struct {
	stem  rune
	piece string
}
