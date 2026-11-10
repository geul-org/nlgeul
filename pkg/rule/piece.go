//ff:type feature=rule type=model
//ff:what Piece — the structural facts of one token sequence read as a marker at one place: whether a stem stands before it, the syllable before it, its place in a chain and the place of the piece after it, and the syllable a fused piece sits under
package rule

// Prev is 0 when nothing stands before the piece. Slot and Next are places in a particle chain or the slots before a
// closing; Next is -1 when nothing follows. Base is "" unless the
// piece is fused into a stem's last syllable.
type Piece struct {
	Form    string
	HasStem bool
	Prev    rune
	Slot    int
	Next    int
	Base    string
}
