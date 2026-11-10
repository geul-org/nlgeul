//ff:type feature=token type=model
//ff:what tail — a run of markers that ends at one place: where the stem before it ends, its pieces, the stem's last syllable as it was before a piece fused into it, and whether 요 may follow
package token

// base is 0 when no piece is fused into the stem.
type tail struct {
	start  int
	pieces []piece
	base   rune
	slot   int
	noYo   bool
}
