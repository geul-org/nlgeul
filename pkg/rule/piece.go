//ff:type feature=rule type=model
//ff:what Piece — the structural facts of one token sequence read as a marker at one place: whether a stem stands before it, the syllable before it, its place in a chain and the place of the piece after it, and the syllable a fused piece sits under
package rule

// Prev is 0 when nothing stands before the piece. Slot and Next are places in a particle chain or the slots before a
// closing; Next is -1 when nothing follows. Base is "" unless the
// piece is fused into a stem's last syllable. Tail is set when the
// piece is a stem: what follows it in its word (Tail…); Count is
// then its observed count (0 when unknown). After is the name of the
// piece that follows, when a rule asks for it. Right is what stands
// right after the piece (Right…), and Stem the kinds of the word before
// it, when a rule asks for them. Num tells that a number stands right
// before the piece.
type Piece struct {
	Form    string
	HasStem bool
	Prev    rune
	Slot    int
	Next    int
	Base    string
	Tail    int
	Count   int
	After   string
	Right   int
	Stem    int
	Num     bool
}
