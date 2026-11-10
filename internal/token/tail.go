//ff:type feature=token type=model
//ff:what tail — a run of markers that ends at one place: where the stem before it ends, its pieces, the place of its first piece, and whether 요 may follow
package token

type tail struct {
	start  int
	pieces []piece
	slot   int
	noYo   bool
}
